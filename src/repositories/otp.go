package repositories

import (
	"context"
	"get-otp-go/src/models"
	"log"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection_name = "otps"

var validate *validator.Validate



type OtpRepository struct {
	Database *mongo.Database
}



type IOtpRepository interface {
	Get(id string) (*models.Otp, error)
}

func NewOtpRepository(db *mongo.Database) IOtpRepository {
	validate = validator.New()
	return &OtpRepository{Database: db}
}

func (br *OtpRepository) Get(id string) (*models.Otp, error) {
	log.Println(id)
	collection := br.Database.Collection(collection_name)
	var otp models.Otp
	
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	err = collection.FindOne(context.Background(), filter).Decode(&otp)
	return &otp, err
}
