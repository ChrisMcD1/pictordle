package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("hello world")
	httpPort := os.Getenv("PORT")
	http.HandleFunc("/", RootPath)

	http.ListenAndServe(":"+httpPort, nil)
}

type Post struct {
	User           string
	Day            time.Time
	Interpretation string
}

func RootPath(w http.ResponseWriter, r *http.Request) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	book := Post{
		User:           "Test-123",
		Day:            time.Now(),
		Interpretation: "I see some cool clouds",
	}

	av, err := dynamodbattribute.MarshalMap(book)
	if err != nil {
		log.Fatalf("Got error marshalleling my new book %s", err)
	}

	tableName := os.Getenv("POST_TABLE_NAME")

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem %s", err)
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = fmt.Fprint(w, "Successfuly created item")
	if err != nil {
		log.Fatalf("Failed to write to response %s", err)
	}
}
