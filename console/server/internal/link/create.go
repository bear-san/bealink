package link

import (
	"github.com/bear-san/bealink/console/server/internal/session"
	"github.com/bear-san/bealink/console/server/pkg/random_string"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/url"
)

func Create(req *gin.Context) {
	token := session.ExtractToken(req)
	if token == nil {
		req.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	_, err := session.Validate(req.Request.Context(), *token)
	if err != nil {
		req.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	l := Values{}
	err = req.BindJSON(&l)
	if err != nil {
		req.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = url.ParseRequestURI(l.URL)
	if l.URL == "" || err != nil {
		req.JSON(400, gin.H{"error": "invalid URL"})
		return
	}

	if l.Path == "" {
		var path string
		path, err = random_string.Create(6)
		if err != nil {
			req.JSON(500, gin.H{"error": err.Error()})
			return
		}
		l.Path = path
	}

	cols := collection.FindOne(req.Request.Context(), bson.M{"path": l.Path})
	if cols.Err() == nil {
		req.JSON(400, gin.H{"error": "path already exists"})
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
