package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
    "log"
	"os"
)

func main(){
	createMyTable()
}

func createMyTable() () {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	//sess, err := session.NewSessionWithOptions(&aws.config{Region: aws.String("us-east-1")},)
	
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "geoMap"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("PK_ID"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("FK_ID"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("PK_ID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("FK_ID"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error Ceate Table")
		fmt.Println(err.Error())
		log.Fatalf("Got error calling CreateTable: %s", err)
		os.Exit(1)
	}

	fmt.Println("Created the table", tableName)
}