package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	`github.com/rhmdnrhuda/simple-rest-api/models`
)

type HotelRepository struct {
	db *mongo.Collection
}

func NewHotelRepository(db *mongo.Database) *HotelRepository {
	return &HotelRepository{
		db: db.Collection("hotels"),
	}
}

func (r *HotelRepository) CreateHotel(hotel models.Hotel) (*models.Hotel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = result.InsertedID.(primitive.ObjectID)
	return &hotel, nil
}

func (r *HotelRepository) GetHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hotel models.Hotel
		if err = cursor.Decode(&hotel); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func (r *HotelRepository) GetHotelById(id string) (*models.Hotel, error) {
	var hotel models.Hotel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.db.FindOne(ctx, bson.M{"_id": objectId}).Decode(&hotel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &hotel, nil
}

func (r *HotelRepository) UpdateHotel(id string, hotel models.Hotel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": hotel})
	return err
}

func (r *HotelRepository) DeleteHotel(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}
