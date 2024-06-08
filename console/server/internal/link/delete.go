package link

import (
	"github.com/bear-san/bealink/console/server/internal/session"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Delete(req *gin.Context) {
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

	lid, err := primitive.ObjectIDFromHex(req.Param("lid"))
	if err != nil {
		req.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = collection.DeleteOne(req.Request.Context(), bson.M{"_id": lid})
	if err != nil {
		req.JSON(500, gin.H{"error": err.Error()})
		return
	}

	req.JSON(http.StatusNoContent, nil)
}
