package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	
	"os"
    "fmt"
    "log"
	"strconv"
	"encoding/csv"
	"time"
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
	
	
	lines, err := ReadCsv("geoMap.csv")
    if err != nil {
        panic(err)
    }
	i := 1
    // Loop through lines & turn into object
    for _, line := range lines {
		
		if line[1] == "" {line[1]="0"}
		CupInt, err := strconv.Atoi(line[1])
		if err != nil {
			//log.Fatalf("Got error in conversion: %s", err)
		}
		u_time := time.Now().String()
		
        item := Item{
			PK_ID: i,
			FK_ID: i,
            Country: line[0],
            Cupcake: CupInt,
			update_time: u_time,
        }
		i=i+1
		AddDBItem(item, svc)
    }
	
}

func AddDBItem(item Item, svc *dynamodb.DynamoDB) (){
	
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new Cupcake item: %s", err)
	}
	tableName := "geoMap"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Println("Successfully added '" + item.Country + "to table " + tableName)

}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}