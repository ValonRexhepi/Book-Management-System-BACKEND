# Book Management System REST Backend Requirements
## Description
---
The goal of this project is to create a Go application that implements a rest API for book management. This API will serve as a demonstration for the use of various technologies with Go. 
## To Implement
---
- [x] Create the Database
  - [x] Create "BookDB" Database with MySQL Docker image
- [x] Creating the Models
  - [x] Create Book Model
- [x] Creating the Controllers
  - [x] Create and open the Go Database Connection
  - [x] Automigrate the Models
  - [ ] Method to Get All Books
  - [ ] Method to Get a Book by Name
  - [ ] Method to Get a Book by Id
  - [ ] Method to Update a Book
  - [ ] Method to Delete a Book
  - [ ] Method to Add a Book

## Technologies Used
---
- Language: [Go 1.18](https://go.dev/dl/)
- Database: [MySql Docker Image :8.0.29](https://hub.docker.com/_/mysql)
- ORM: [Gorm.io](https://gorm.io/index.html)