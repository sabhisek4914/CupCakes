package main

import (
  "fmt"
  "net/http"
  "strings"
  "time"
  "log"
  "io/ioutil"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/signer/v4"
  "github.com/aws/aws-sdk-go/lambda"
  "context"
  "github.com/aws/aws-sdk-go/events"
  //"reflect"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, e events.DynamoDBEvent){
  // Basic information for the Amazon Elasticsearch Service domain
  domain := "https://search-ddb-to-es-vxbv3zpxppqxmlolen5ttlkgde.us-east-1.es.amazonaws.com" // e.g. https://my-domain.region.es.amazonaws.com
  index := "lambda-index"
  type1:="lambda-type"
  endpoint := domain + "/" + index + "/" + type1
  region := "us-east-1" // e.g. us-east-1
  service := "es"

  // Get credentials from environment variables and create the AWS Signature Version 4 signer
  credentials := credentials.NewEnvCredentials()
  signer := v4.NewSigner(credentials)

  // An HTTP client for sending the request
  client := &http.Client{}  
  
  for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)
		var co=""
		var cup=""
		var o_co=""
		var o_cup=""

		// Print new values for attributes of type String
		for name, value := range record.Change.NewImage {
			if value.DataType() == events.DataTypeString {
				if name=="Country"{
						co =value.String()
				}
				if name=="Cupcake"{
						cup =value.String()
				}
			}
			
			//fmt.Printf("Attribute name: %s, value: %s", name, value.String())
		}
		//fmt.Printf("Country: %s Cupcake: %s",co,cup)
		for name, value := range record.Change.OldImage {
			if value.DataType() == events.DataTypeString {
				if name=="Country"{
						o_co =value.String()
				}
				if name=="Cupcake"{
						o_cup =value.String()
				}
			}
			//fmt.Printf("OLDAttribute name: %s, value: %s", name, value.String())
		}
		//fmt.Printf("OLDCountry: %s OLDCupcake: %s",o_co,o_cup)
		
		dt := time.Now()
		json := `{"eventTimeStamp": "`+dt.String()+`","eventType": "`+record.EventName+`","NewImage": {"Cupcake": {"S": "`+cup+`"},"Country": {"S": "`+co+`"}},"OldImage": {"Cupcake": {"S": "`+o_cup+`"},"Country": {"S": "`+o_co+`"}}}`
		fmt.Printf(json)
		body := strings.NewReader(json)
		
		// Form the HTTP request
		req, err := http.NewRequest(http.MethodPost, endpoint, body)
		if err != nil {
			fmt.Println(err)
		}
		// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
		 req.Header.Add("Content-Type", "application/json")

		// Sign the request, send it, and Printf the response
		signer.Sign(req, body, service, region, dt)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resp.Status + "\n")
		  
		body1, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body1)
		log.Printf(sb)
	}
}