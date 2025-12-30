package colletions

import (
	"time"
)

type SubscriptionCenter struct {
	Key         string    `bson:"key"`
	Option      string    `bson:"option"`
	Value       string    `bson:"value"`
	Description *string   `bson:"description"`
	Type        string    `bson:"type"`
	CreatedAt   time.Time `bson:"created_at"`
}

type SubscriptionCenterFrequency struct {
	Key       string    `bson:"key"`
	Name      string    `bson:"name"`
	Enabled   bool      `bson:"enabled"`
	CreatedAt time.Time `bson:"created_at"`
}