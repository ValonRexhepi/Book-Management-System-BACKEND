package main

import (
	"fmt"

	"github.com/ValonRexhepi/Book-Management-System-REST/controllers"
)

func main() {
	controllers.Connect()
	controllers.Migrate()
	fmt.Println("App Started")
	for {

	}
}
