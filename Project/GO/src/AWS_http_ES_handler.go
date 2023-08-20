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
  //"github.com/aws/aws-sdk-go/events"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context){
  // Basic information for the Amazon Elasticsearch Service domain
  domain := "https://search-ddb-to-es-vxbv3zpxppqxmlolen5ttlkgde.us-east-1.es.amazonaws.com" // e.g. https://my-domain.region.es.amazonaws.com
  index := "lambda-index"
  type1:="lambda-type"
  endpoint := domain + "/" + index + "/" + type1
  region := "us-east-1" // e.g. us-east-1
  service := "es"

  // Sample JSON document to be included as the request body
  json := `{ "title": "Thor: Ragnarok", "director": "Taika Waititi", "year": "2017" }`
  body := strings.NewReader(json)

  // Get credentials from environment variables and create the AWS Signature Version 4 signer
  credentials := credentials.NewEnvCredentials()
  signer := v4.NewSigner(credentials)

  // An HTTP client for sending the request
  client := &http.Client{}

  // Form the HTTP request
  req, err := http.NewRequest(http.MethodPost, endpoint, body)
  if err != nil {
    fmt.Println(err)
  }

  // You can probably infer Content-Type programmatically, but here, we just say that it's JSON
  req.Header.Add("Content-Type", "application/json")

  // Sign the request, send it, and Printf the response
  signer.Sign(req, body, service, region, time.Now())
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
  //return fmt.SPrintf(ddbEvent), nil
}