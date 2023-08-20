package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "fmt"
    "log"
	"strconv"
)

type Item struct{
	Year int
	Year2 int
	Title string
	
}

func main(){
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	
	//Add item to the table
	item :=Item{Year:1970, Year2:2020, Title: "Spider"}
	AddDBItem(item, svc)
	fmt.Println("main")
}

func AddDBItem(item Item, svc *dynamodb.DynamoDB) (){
	
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	tableName := "MoviesDummy"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	year := strconv.Itoa(item.Year)

	fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)

}