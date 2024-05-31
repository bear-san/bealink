package link

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func List(req *gin.Context) {
	ctx := req.Request.Context()
	var links []Link
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, ctx)
	err = cur.All(ctx, &links)
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	req.JSON(200, links)
}
