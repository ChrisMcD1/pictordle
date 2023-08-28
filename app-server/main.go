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

	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	router := gin.Default()

	router.Use(svcProvider(svc))
	router.GET("/", RootPath)
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	router.Run(":" + httpPort)

}

func svcProvider(svc *dynamodb.DynamoDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dynamoDB", svc)
	}
}

func getSvc(c *gin.Context) *dynamodb.DynamoDB {
	return c.MustGet("dynamoDB").(*dynamodb.DynamoDB)
}

type Post struct {
	User           string
	Day            time.Time
	Interpretation string
}

func RootPath(c *gin.Context) {
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

	_, err = getSvc(c).PutItem(input)
	if err != nil {
		log.Fatalf("Got error called PutItem %s", err)
	}

	c.String(http.StatusOK, "I added an item!")
}
