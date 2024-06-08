package link

import (
	"crypto/rand"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	if l.Path == "" {
		var path string
		path, err = MakeRandomStr(6)
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

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
