package repositories

import (
	"context"
	"get-otp-go/src/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type OtpRepository struct {
	Database *mongo.Database
}

type IOtpRepository interface {
	Post(otp *models.Otp) (*models.Otp, error)
}

func NewOtpRepository(db *mongo.Database) IOtpRepository {
	return &OtpRepository{Database: db}
}

func (br *OtpRepository) Post(otp *models.Otp) (*models.Otp, error) {
	collection := br.Database.Collection("otps")
	_, err := collection.InsertOne(context.TODO(), otp)
	if err != nil {
		// Log the error and return
		log.Printf("Error while inserting OTP: %v", err)
		return nil, err
	}
	// Return the inserted document and nil for the error
	return otp, nil
}
