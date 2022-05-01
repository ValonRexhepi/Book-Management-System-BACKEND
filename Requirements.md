# Book Management System REST Backend Requirements
## Description
---
The goal of this project is to create a N-Tier Go application that implements a rest API for book management. This API will serve as a demonstration for the use of various technologies with Go. 
## To Implement
---
- [ ] Database
  - [ ] Create MySQL Docker image 
- [ ] Creating the DTO
  - [ ] Create Book Model
- [ ] Creating the DAL
  - [ ] Create and open the Go Database Connection
  - [ ] Automigrate the DTO Models
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