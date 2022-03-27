# Golang-codefood
Golang endpoint about food recipe and its ingredients (Still working in progress)

## Requirement
1. Go: 1.17
2. MySql Server
3. Endpoint testing tools (Eg: Postman)

# How to run:
Build the project
`go build`

Run the main script
`go run main/main.go`

# Testing
1. HTTP GET `localhost:10000/recipe-categories`
2. HTTP POST `localhost:10000/recipe-categories`.
   Body = Json `{
    "name" : "Testing add"
   }`
3. HTTP PUT `localhost:10000/recipe-categories/:id`
    Body = Json `{
    "name" : "Testing update"
    }`
4. HTTP DELETE `localhost:10000/recipe-categories/:id`
