package link

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(req *gin.Context) {
	l := Values{}
	err := req.BindJSON(&l)
	if err != nil {
		req.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var r *mongo.InsertOneResult
	r, err = collection.InsertOne(req.Request.Context(), l)
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var insertedID primitive.ObjectID
	insertedID = r.InsertedID.(primitive.ObjectID)

	req.JSON(201, gin.H{"id": insertedID.Hex()})
}
