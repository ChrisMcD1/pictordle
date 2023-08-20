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
)

func main() {
	fmt.Println("hello world")
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	http.HandleFunc("/", RootPath)

	http.ListenAndServe(":"+httpPort, nil)
}

type Book struct {
	Title  string
	Author string
}

func RootPath(w http.ResponseWriter, r *http.Request) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	book := Book{"My first book", "Christopher McDonenll"}

	av, err := dynamodbattribute.MarshalMap(book)
	if err != nil {
		log.Fatalf("Got error marshalleling my new book %s", err)
	}

	tableName := "Books"

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
