package controllers

import (
	"Farming_data/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getFarmerByFarmerID(farmerId string) (models.Farmer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var farmer models.Farmer
	err := farmerCollection.FindOne(ctx, bson.M{"farmer_id": farmerId}).Decode(&farmer)
	if err != nil {
		return models.Farmer{}, err
	}

	return farmer, nil
}

func GetAllFarms() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := farmCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer result.Close(ctx)

		var allFarms []models.Farm
		if err := result.All(ctx, &allFarms); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, allFarms)
	}
}

func GetDueSchedules() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		farmId := c.Param("farmId")

		var farm models.Farm
		err := farmCollection.FindOne(ctx, bson.M{"farm_id": farmId}).Decode(&farm)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sowingDate, err := time.Parse("2006-01-02", farm.SowingDate)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result, err := scheduleCollection.Find(ctx, bson.M{"farm_id": farmId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer result.Close(ctx)

		var dueSchedules []models.Schedule

		for result.Next(ctx) {
			var schedule models.Schedule
			if err := result.Decode(&schedule); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			dueDate := sowingDate.Add(time.Duration(schedule.DaysAfterSowing) * 24 * time.Hour)
			today := time.Now().Format("2006-01-02")
			tomorrow := time.Now().Add(24 * time.Hour).Format("2006-01-02")

			if dueDate.Format("2006-01-02") == today || dueDate.Format("2006-01-02") == tomorrow {
				dueSchedules = append(dueSchedules, schedule)
			}
		}
		if len(dueSchedules) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No schedules due today or tomorrow"})
			return
		}
		c.JSON(http.StatusOK, dueSchedules)

	}
}

func GetAllFarmersGrowingCrop() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := farmCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer result.Close(ctx)

		uniqueFarmers := make(map[string]models.Farmer)

		for result.Next(ctx) {
			var farm models.Farm
			if err := result.Decode(&farm); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Check if the farmer ID is already in the map
			if _, exists := uniqueFarmers[farm.FarmerId]; !exists {
				farmer, err := getFarmerByFarmerID(farm.FarmerId)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				// Add the farmer to the map
				uniqueFarmers[farm.FarmerId] = farmer
			}
		}

		// Convert the map values to a slice
		var farmers []models.Farmer
		for _, farmer := range uniqueFarmers {
			farmers = append(farmers, farmer)
		}

		c.JSON(http.StatusOK, farmers)
	}
}

func CalculateBillOfMaterials() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get farmerID and fertilizer prices from the request parameters
		farmerId := c.Param("farmerId")
		fertilizerData := c.Query("fertilizerData")

		// Parse the JSON-encoded string into a map
		var fertilizerPricesMap map[string]float64
		if err := json.Unmarshal([]byte(fertilizerData), &fertilizerPricesMap); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		// Assuming scheduleCollection is your MongoDB collection reference
		var schedules []models.Schedule
		result, err := scheduleCollection.Find(ctx, bson.M{"farmer_id": farmerId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = result.All(ctx, &schedules); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		totalBill := 0.0

		for _, schedule := range schedules {
			// Assuming the fertilizer name is stored in the "Fertilizer" field of the Schedule model
			fertilizerName := schedule.Fertilizer

			// Check if the fertilizer name exists in the provided prices
			if price, ok := fertilizerPricesMap[fertilizerName]; ok {
				totalBill += price * schedule.Quantity
			} else {
				// Handle the case where the fertilizer name is not found in the provided prices
				c.JSON(http.StatusBadRequest, gin.H{"error": "Price not provided for fertilizer: " + fertilizerName})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"totalBill": totalBill})
	}
}
