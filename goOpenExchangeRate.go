package main
import (
"fmt"
"io/ioutil"
"net/http"
"encoding/json"
"strings"
)

// Get all currency rates from Open Exchange Rates API
func getallCurrency(){

      // Get data in JSON format from Open Exchange Rates API using a Key
      response, err := http.Get("http://openexchangerates.org/api/latest.json?app_id=YOUR_APP_ID")
      if err != nil {
          panic(err.Error())
          return
      }
      defer response.Body.Close()
      body, err := ioutil.ReadAll(response.Body)
      if err != nil {
          panic(err.Error())
          return
      }

      // Removing the leading/trailing white spaces in the data      
      data := string(body)
      data = strings.TrimSpace(data)

      // Parsing the JSON data
      parseData := make(map[string]interface{})
      err = json.Unmarshal([]byte(data), &parseData)
      if err !=nil {
		      panic(err)
      }

      // Iterate, format and display the data interactively
      for index, _ := range parseData { 
          if index != "rates" { 
                fmt.Println( index, ":", parseData[index])     
          } else {   
                currencies := parseData["rates"].(map[string]interface{})
                for currency, _ := range currencies {
                    fmt.Println( currency,  ":", currencies[currency].(float64)) 
                }
          }
      }
  
    }    

// Get currency details from Open Exchange Rates
func main() {

      getallCurrency() 
           
}