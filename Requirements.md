# Book Management System REST Backend Requirements
## Description
---
The goal of this project is to create a Go application that implements a rest API for book management. This API will serve as a demonstration for the use of various technologies with Go. 
## To Implement
---
- [x] Create the Database
  - [x] Create "BookDB" Database with MySQL Docker image
- [x] Create the Models
  - [x] Create Book Model
- [x] Create the Controllers
  - [x] Create and open the Go Database Connection
  - [x] Automigrate the Models
  - [x] Method to Get All Books
  - [x] Method to Get a Book by ISBN
  - [x] Method to Get a Book by Id
  - [x] Method to Update a Book
  - [x] Method to Delete a Book
  - [x] Method to Add a Book
- [x] Create the Server Component
  - [x] Create Method to launch, bind routes and run the server
- [x] Create the Routes Component
  - [x] Route to Get All Books
  - [x] Route to Get a Book by ISBN
  - [x] Route to Get a Book by ID
  - [x] Route to Update a book
  - [x] Route to Delete a Book
  - [x] Route to Add a book
- [x] Create Unit Tests
  - [x] Test Method to Get All Books
  - [x] Test Method to Get a Book by ISBN
  - [x] Test Method to Get a Book by Id
  - [x] Test Method to Update a Book
  - [x] Test Method to Delete a Book
  - [x] Test Method to Add a Book
- [x] Manual Test with Postman
  - [x] Test Route to Get All Books
  - [x] Test Route to Get a Book by ISBN
  - [x] Test Route to Get a Book by ID
  - [x] Test Route to Update a Book
  - [x] Test Route to Delete a Book
  - [x] Test Routte to Add a Book

## Technologies Used
---
- Language: [Go 1.18](https://go.dev/dl/)
- Database: [MySql Docker Image :8.0.29](https://hub.docker.com/_/mysql)
- ORM: [Gorm.io](https://gorm.io/index.html)