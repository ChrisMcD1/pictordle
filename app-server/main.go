package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("hello world")
	httpPort := os.Getenv("PORT")

	router := gin.Default()
	router.GET("/", RootPath)
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	router.Run(":" + httpPort)
}

type Post struct {
	User           string
	Day            time.Time
	Interpretation string
}

func RootPath(c *gin.Context) {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	post := Post{
		User:           "ME",
		Day:            time.Now(),
		Interpretation: "This looks really cool",
	}
	tableName := os.Getenv("POST_TABLE_NAME")

	av, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		log.Fatalf("Error calling marshal map %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error called PutItem %s", err)
	}

	c.String(http.StatusOK, "I added an item!")
}
