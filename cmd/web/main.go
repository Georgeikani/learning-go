package main

import (
	"fmt"
	"net/http"
	"github.com/georgeikani/hotel-booking/pkg/handlers"
)

// // Function and Handlers


var portNumber = ":8080"

func main(){

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

	 
//Making our application module-ready
