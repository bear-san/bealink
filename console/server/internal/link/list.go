package link

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func List(req *gin.Context) {
	ctx := req.Request.Context()
	var links []Link
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer cur.Close(ctx)
	err = cur.All(ctx, &links)
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	req.JSON(200, links)
}
