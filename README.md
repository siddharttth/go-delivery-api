# go-delivery-api


# 🚚 Go Delivery Cost Calculator API

This is a backend REST API built with **Go** and **Fiber** that calculates the **minimum delivery cost** of transporting a set of products from multiple distribution centers (`C1`, `C2`, `C3`) to a common destination (`L1`). It considers product weights, inter-location distances, and delivery cost brackets.


You can test the /calculate endpoint using the following cURL command:
```sh
curl --location 'https://twf-siddharth.up.railway.app/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "products": {
    "A": 1,
    "D": 1,
    "I": 3
  }
}'
```
✅ This will send a POST request with the specified product quantities to your deployed API and return the minimum delivery cost based on the logic you've implemented.
