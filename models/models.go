package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CountryId string             `bson:"country_id" json:"country_id"`
	Name      string             `bson:"name" json:"name" validate:"required,min=2,max=100"`
}

type Farmer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FarmerId  string             `bson:"farmer_id" json:"farmer_id"`
	Phone     string             `bson:"phone" json:"phone" validate:"required,min=10"`
	Name      string             `bson:"name" json:"name" validate:"required,min=2"`
	Language  string             `bson:"language" json:"language"`
	CountryId string             `bson:"country_id" json:"country_id" validate:"required"`
}

type Farm struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FarmId     string             `bson:"farm_id" json:"farm_id"`
	Area       float64            `bson:"area" json:"area"`
	Village    string             `bson:"village" json:"village"`
	Crop       string             `bson:"crop" json:"crop" validate:"required,min=2"`
	SowingDate string             `bson:"sowing_date" json:"sowing_date" validate:"required"`
	FarmerId   string             `bson:"farmer_id" json:"farmer_id" validate:"required"`
}

type Schedule struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ScheduleId      string             `bson:"schedule_id" json:"schedule_id"`
	DaysAfterSowing int                `bson:"days_after_sowing" json:"days_after_sowing" validate:"required"`
	Fertilizer      string             `bson:"fertilizer" json:"fertilizer"`
	Quantity        float64            `bson:"quantity" json:"quantity"`
	QuantityUnit    string             `bson:"quantity_unit" json:"quantity_unit"`
	FarmId          string             `bson:"farm_id" json:"farm_id" validate:"required"`
	FarmerId        string             `bson:"farmer_id" json:"farmer_id"`
}
