# shopify-order-automation-pipedream-workflow


## Overview
Automated workflow built on Pipedream which receives a Shopify order webhook to send emails to customer if specific conditions are met.

## Steps

1
|
HTTP Webhook Trigger
|
Receives Shopify order JSON
|
|
2
|
GO Code
|
Check 3 conditions
|
|
3
|
Email
|
Send greeting email to the cutomer immediately
|
|
4
|
Delay
|
Wait for 5 minutes
|
|
5
|
Gmail
|
Send discount email to the customer
|

## Conditions to check
- Order Tag = "MakeOrder"
- Customer tag = "ColdCustomer"
- Order Amount > 2500

## Emails to send
- Email 1 (sent immediately):** Greeting email confirming order is places successfully
- Email 2 (After 5 mins):** Email notification for limited period discount

## Tech Stack
- PipeDream (workflow automation)
- Golang (condition logic)
- Gmail (email delivery)

## Testing
- Send below payload to the webhook URL:
```json
{
  "tags": "MakeOrder",
  "total_price": "3000.00",
  "customer": {
    "first_name": "Ankit",
    "email": "ankit.787as@gmail.com",
    "tags": "ColdCustomer"
  }
}
```

## Live Workflow
https://eoyi8okzenkyx7v.m.pipedream.net

## Workflow Screenshots
