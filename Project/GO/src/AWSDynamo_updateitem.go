package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "fmt"
	"time"
	"math/rand"
    "log"
	"strconv"
)
type Item struct{
	PK_ID int
	FK_ID int
	Country string
	Cupcake int
	update_time string
}

func main(){
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	
	
	d :=250
	for i := 0; i < 100; i++{
		sa :=rand.Intn(d)
		var id string = strconv.Itoa(sa)
		P_ID := id
		F_ID := id
		// Update item in table geoMap
		
		Cupcake := strconv.Itoa(rand.Intn(d))
		update_time := time.Now().String()
		UpdateItem(update_time,"update_time",P_ID,F_ID,svc)
		UpdateItem(Cupcake,"Cupcake",P_ID,F_ID,svc)
		fmt.Println("Successfully updated PK_ID",P_ID)
		fmt.Println()
    }
	
}
func UpdateItem(data string, col string,P_ID string,F_ID string, svc *dynamodb.DynamoDB){
	u_s:="set "+col+" = :r"
	tableName := "geoMap"
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(data),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK_ID": {
				N: aws.String(P_ID),
			},		
			"FK_ID": {
				N: aws.String(F_ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(u_s),
	}
	
	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}
		
}

