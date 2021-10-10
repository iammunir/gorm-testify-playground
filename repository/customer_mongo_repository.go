package repository

import (
	"errors"
	"learn-gorm/config"
	"learn-gorm/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerMongoRepository interface {
	FindByName(string) ([]models.CustomerMongo, error)
}

type customerMongoRepository struct {
	Collection *mongo.Collection
}

func GetCustomerMongoRepository(conn *mongo.Database) CustomerMongoRepository {
	return &customerMongoRepository{
		Collection: conn.Collection("customer_mapping"),
	}
}

func (repo *customerMongoRepository) FindByName(name string) ([]models.CustomerMongo, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{"$text": bson.M{"$search": name}})
	if err != nil {
		return nil, err
	}

	var customers []models.CustomerMongo
	for cursor.Next(ctx) {
		var customer models.CustomerMongo
		errDec := cursor.Decode(&customer)
		if errDec != nil {
			log.Println(errDec.Error())
		}
		customers = append(customers, customer)
	}

	if len(customers) == 0 {
		return nil, errors.New("record not found")
	}

	return customers, nil
}
