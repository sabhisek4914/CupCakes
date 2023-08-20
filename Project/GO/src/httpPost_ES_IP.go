package main
import (
   "bytes"
   "encoding/json"
   "io/ioutil"
   "log"
   "net/http"
)

func main() {
//Encode the data
   postBody, _ := json.Marshal(map[string]string{
      "name":  "Toby",
      "email": "Toby@example.com",
   })
   responseBody := bytes.NewBuffer(postBody)
//Leverage Go's HTTP Post function to make request
   resp, err := http.Post("https://search-ddb-to-es-vxbv3zpxppqxmlolen5ttlkgde.us-east-1.es.amazonaws.com/lambda-index/lambda-type", "application/json", responseBody)
//Handle Error
   if err != nil {
      log.Fatalf("An Error Occured %v", err)
   }
   defer resp.Body.Close()
//Read the response body
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
   sb := string(body)
   log.Printf(sb)
}