package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CountryId string             `bson:"country_id"`
	Name      string             `bson:"name"`
}

type Farmer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FarmerId  string             `bson:"farmer_id"`
	Phone     string             `bson:"phone"`
	Name      string             `bson:"name"`
	Language  string             `bson:"language"`
	CountryID string             `bson:"country_id"`
}

type Farm struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FarmId     string             `bson:"farm_id"`
	Area       float64            `bson:"area"`
	Village    string             `bson:"village"`
	Crop       string             `bson:"crop"`
	SowingDate time.Time          `bson:"sowing_date"`
	FarmerID   string             `bson:"farmer_id"`
}

type Schedule struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ScheduleId      string             `bson:"schedule_id"`
	DaysAfterSowing int                `bson:"days_after_sowing"`
	Fertilizer      string             `bson:"fertilizer"`
	Quantity        float64            `bson:"quantity"`
	QuantityUnit    string             `bson:"quantity_unit"`
	FarmID          string             `bson:"farm_id"`
}
