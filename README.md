# Book Management System
## Goal
The goal of this project is to create a Go application that implements a rest 
API for book management. This API will serve as a demonstration and training 
for the use of various technologies with Go. 

## Methods Available
- Delete a book by ID
- Add a book
- Update a book
- Get all books 
- Get a book by ID
- Get a book by ISBN

## Routes Requests
- [DELETE]
  - localhost:8080/books/deletebook/:bookid &rarr; Delete book by ID
- [POST]
  - localhost:8080/books &rarr; Add a new book
  - localhost:8080/books/updatebook &rarr; Update existing book
- [GET] 
	- localhost:8080/books &rarr; Get all books
	- localhost:8080/books/getbookbyid/:bookid &rarr; Get book by ID
	- localhost:8080/books/getbookbyisbn/:bookisbn &rarr; Get book by ISBN

## Libraries Used
- [Gorm, ORM Management Library](https://gorm.io/index.html)
- [Gin, HTTP WebFramework](https://github.com/gin-gonic/gin)

## Docker Image
- [MySql Docker Image :latest](https://hub.docker.com/_/mysql)
- User "root" / Password "secret" (if needed)

## How to Use ?
1. Launch the container with the command "docker-container up". 
2. Run the program in Go with the command "go run main.go". 
3. Query the webserver (for example with postman) for the different available requests.
