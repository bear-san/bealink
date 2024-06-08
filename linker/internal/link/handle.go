package link

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Handle(req *gin.Context) {
	path := req.Param("path")

	var l Link
	err := collection.FindOne(req.Request.Context(), bson.M{"path": path}).Decode(&l)
	if err != nil {
		req.String(404, "not found")
		return
	}

	req.Redirect(301, l.URL)
}

type Link struct {
	ID     string `json:"id" bson:"_id"`
	Values `bson:",inline"`
}

type Values struct {
	URL         string `json:"url" bson:"url"`
	Path        string `json:"path" bson:"path"`
	Description string `json:"description" bson:"description"`
}
