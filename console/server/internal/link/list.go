package link

import (
	"context"
	"github.com/bear-san/bealink/console/server/internal/session"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func List(req *gin.Context) {
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
