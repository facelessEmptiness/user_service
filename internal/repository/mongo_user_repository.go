package repository

import (
	"context"
	"github.com/facelessEmptiness/user_service/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepo struct {
	coll *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) UserRepository {
	return &mongoUserRepo{coll: db.Collection("users")}

}
func (r *mongoUserRepo) Create(user *domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}

	result, err := r.coll.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	oid := result.InsertedID.(primitive.ObjectID).Hex()
	return oid, nil
}

func (r *mongoUserRepo) GetByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u domain.User
	if err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *mongoUserRepo) GetByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, _ := primitive.ObjectIDFromHex(id)
	var u domain.User
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
