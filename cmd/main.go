package main

import (
	"context"
	"example/infrastructure/arpservice"
	"example/routers"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb://localhost:27017/example"))
	if err != nil {
		panic(err)
	}

	node, err := snowflake.NewNode(1)
	routers.AddressBookService = arpservice.NewArpAddressBookService(mongoClient, node)

	r := gin.Default()
	mbrs := r.Group("/addressbook")
	routers.SetAddressBookRoutes(mbrs)
	r.Run() // listen and serve on 0.0.0.0:8080
}
