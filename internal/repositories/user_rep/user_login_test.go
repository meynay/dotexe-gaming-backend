package user_rep_test

import (
	"context"
	"store/internal/repositories/user_rep"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAll(t *testing.T) {
	connectionString := "mongodb://admin:password@localhost:27017/?authSource=admin"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	database := client.Database("store")
	userCollection := database.Collection("user")
	userRep := user_rep.NewUserRepository(userCollection)
	user, err := userRep.InsertUserByEmail("man_test@email_test.com", "123")
	assert.Nil(t, err)
	user2, err := userRep.GetUserByEmail("man_test@email_test.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, user2.ID)
	// userRep.SaveToken(user.ID, "123")
	// err = userRep.TokenExists(user.ID, "123")
	// assert.Nil(t, err)
	user3, err := userRep.CheckUser("man_test@email_test.com", "123")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, user3.ID)
	user, err = userRep.InsertUserByPhone("09123123123")
	assert.Nil(t, err)
	user2, err = userRep.GetUserByPhone("09123123123")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, user2.ID)
	userCollection.DeleteMany(ctx, bson.M{"email": "man_test@email_test.com"})
	userCollection.DeleteMany(ctx, bson.M{"phone": "09123123123"})
}
