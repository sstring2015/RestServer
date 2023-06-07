package store

import (
	"context"
	"log"

	"github.com/RestServer/pkg/config"
	"github.com/RestServer/pkg/database"
	"github.com/RestServer/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Store struct {
	db *mongo.Database
}

var instance *Store

// GetStore returns the singleton instance of the Store.
func GetStore(cfg config.DatabaseConfig) *Store {
	if instance == nil {
		db, err := database.ConnectToMongoDB(cfg)
		if err != nil {
			log.Fatal(err)
		}

		instance = &Store{
			db: db,
		}
	}
	return instance
}

// GetUserByEmail retrieves a user by email.
func (s *Store) GetUserByEmail(filter bson.M) (user models.User, err error) {
	collection := s.db.Collection("users")
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

// InsertUser inserts a new user into the database.
func (s *Store) InsertUser(data models.User) error {
	collection := s.db.Collection("users")
	id := primitive.NewObjectID()
	user := models.User{
		ID:    id,
		Name:  data.Name,
		Age:   data.Age,
		Email: data.Email,
	}
	_, err := collection.InsertOne(context.Background(), user)

	return err
}

// GetCarByModelBrand retrieves a car by model and brand.
func (s *Store) GetCarByModelBrand(filter bson.M) (car models.Car, err error) {
	collection := s.db.Collection("cars")
	err = collection.FindOne(context.Background(), filter).Decode(&car)
	return car, err
}

// GetCarByModelBrand retrieves a car by model and brand.
func (s *Store) GetCarByID(filter bson.M) (car models.Car, err error) {
	collection := s.db.Collection("cars")
	err = collection.FindOne(context.Background(), filter).Decode(&car)
	return car, err
}

// InsertCar inserts a new car into database
func (s *Store) InsertCar(data models.Car) error {
	collection := s.db.Collection("cars")
	id := primitive.NewObjectID()
	car := models.Car{
		ID:    id,
		Brand: data.Brand,
		Model: data.Model,
		DOM:   data.DOM,
		Price: data.Price,
	}
	_, err := collection.InsertOne(context.Background(), car)

	return err
}

func (s *Store) GetAllCars() ([]models.Car, error) {
	collection := s.db.Collection("cars")

	// Set options for sorting or filtering if required
	findOptions := options.Find()

	// Find all cars
	cur, err := collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var cars []models.Car

	// Iterate over the result cursor
	for cur.Next(context.Background()) {
		var car models.Car
		if err := cur.Decode(&car); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}
