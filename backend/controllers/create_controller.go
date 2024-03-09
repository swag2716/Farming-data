package controllers

import (
	"Farming_data/database"
	"Farming_data/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var countryCollection = database.OpenCollection(database.Client, "countryCollection")
var farmerCollection = database.OpenCollection(database.Client, "farmerCollection")
var farmCollection = database.OpenCollection(database.Client, "farmCollection")
var scheduleCollection = database.OpenCollection(database.Client, "scheduleCollection")

var validate = validator.New()

func CreateCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var country models.Country

		if err := c.BindJSON(&country); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(country)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		country.ID = primitive.NewObjectID()
		country.CountryId = country.ID.Hex()

		result, insertErr := countryCollection.InsertOne(ctx, country)
		if insertErr != nil {
			msg := fmt.Sprintln("country data was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return

		}

		c.JSON(http.StatusOK, result)
	}

}

func CreateFarmer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var farmer models.Farmer
		var country models.Country

		if err := c.BindJSON(&farmer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(farmer)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := countryCollection.FindOne(ctx, bson.M{"country_id": farmer.CountryId}).Decode(&country)

		if err != nil {
			msg := fmt.Sprintln("country was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		farmer.ID = primitive.NewObjectID()
		farmer.FarmerId = farmer.ID.Hex()

		result, insertErr := farmerCollection.InsertOne(ctx, farmer)
		if insertErr != nil {
			msg := fmt.Sprintln("farmer data was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return

		}

		c.JSON(http.StatusOK, result)
	}

}
func CreateFarm() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var farm models.Farm
		var farmer models.Farmer

		if err := c.BindJSON(&farm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(farm)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := farmerCollection.FindOne(ctx, bson.M{"farmer_id": farm.FarmerId}).Decode(&farmer)

		if err != nil {
			msg := fmt.Sprintln("farmer was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		farm.ID = primitive.NewObjectID()
		farm.FarmId = farm.ID.Hex()

		date, err := time.Parse("2006-01-02", farm.SowingDate)
		if err != nil {
			msg := fmt.Sprintln("Invalid date format. Please use YYYY-MM-DD.")
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		// Format the date as a string without the time component
		formattedDate := date.Format("2006-01-02")

		farm.SowingDate = formattedDate

		result, insertErr := farmCollection.InsertOne(ctx, farm)
		if insertErr != nil {
			msg := fmt.Sprintln("farm data was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return

		}

		c.JSON(http.StatusOK, result)
	}

}
func CreateSchedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var schedule models.Schedule
		var farm models.Farm

		if err := c.BindJSON(&schedule); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(schedule)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := farmCollection.FindOne(ctx, bson.M{"farm_id": schedule.FarmId}).Decode(&farm)

		if err != nil {
			msg := fmt.Sprintln("farm was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		schedule.ID = primitive.NewObjectID()
		schedule.ScheduleId = schedule.ID.Hex()
		schedule.FarmerId = farm.FarmerId

		result, insertErr := scheduleCollection.InsertOne(ctx, schedule)
		if insertErr != nil {
			msg := fmt.Sprintln("schedule data was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return

		}

		c.JSON(http.StatusOK, result)
	}

}
