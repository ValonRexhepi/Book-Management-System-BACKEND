package main

import (
	"github.com/ValonRexhepi/Book-Management-System-BACKEND/controllers"
	"github.com/ValonRexhepi/Book-Management-System-BACKEND/server"
)

// Main function of the application, connect to the database, migrate the
// schemas and launch the Web Server.
func main() {
	controllers.Connect()
	controllers.Migrate()
	server.LaunchServer()
}
