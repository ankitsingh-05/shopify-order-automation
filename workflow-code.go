package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  pd "github.com/PipedreamHQ/pipedream-go"
  )

func main(){
  trigger, ok := pd.Steps["trigger"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: trigger data not found")
		os.Exit(1)
	}
  
  event, ok := trigger["event"].(map[string]interface{})
  if !ok {
    fmt.Println("Error: event data not found")
    os.Exit(1)
  }
  
  body, ok := event["body"].(map[string]interface{})
  if !ok {
    fmt.Println("Error: body not found, send a payload first")
  }

  orderTags := fmt.Sprintf("%v", body["tags"])

  customer := body["customer"].(map[string]interface{})
  customerTags := fmt.Sprintf("%v", customer["tags"])
  customerEmail := fmt.Sprintf("%v", customer["email"])
  customerName := fmt.Sprintf("%v", customer["first_name"])

  orderAmount, _ := strconv.ParseFloat(fmt.Sprintf("%v", body["total_price"]), 64)

  hasMakeOrder := strings.Contains(orderTags, "MakeOrder")
  hasColdCustomer := strings.Contains(customerTags, "ColdCustomer")
  isHighValue := orderAmount > 2500

  fmt.Printf("hasMakeOrder: %v, hasColdCustomer: %v, isHighVale: %v\n", hasMakeOrder, hasColdCustomer, isHighValue)

  if !hasMakeOrder || !hasColdCustomer || !isHighValue {
    fmt.Println("Conditions not met, stopping the workflow")
    os.Exit(1)
  }

  pd.Export("customerEmail", customerEmail)
  pd.Export("customerName", customerName)
}
