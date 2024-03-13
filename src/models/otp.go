package models

import (
	"time"
)

type Otp struct {
	Otp              string    `bson:"otp"`
	CreationDate     time.Time `bson:"creationDate"`
	ModificationDate time.Time `bson:"modificationDate"`
	Status           string    `bson:"status"`
	ValidAttempts    int       `bson:"validAttempts"`
	RetryAttempts    int       `bson:"retryAttempts"`
	UserPhoneNumber  string    `bson:"userPhoneNumber"`
	UserID           string    `bson:"userId"`
}
