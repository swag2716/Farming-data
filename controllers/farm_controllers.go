package controllers

import (
	"Farming_data/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getSowingDateForFarm(farmID string) (time.Time, error) {
	var farm models.Farm
	err := farmCollection.FindOne(context.TODO(), bson.M{"farm_id": farmID}).Decode(&farm)
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse("2006-01-02", farm.SowingDate)
}

func GetDueSchedules() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := scheduleCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		defer result.Close(ctx)

		var dueSchedules []models.Schedule

		for result.Next(ctx) {
			var schedule models.Schedule
			if err := result.Decode(&schedule); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			sowingDate, err := getSowingDateForFarm(schedule.FarmId)
			if err != nil {
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
