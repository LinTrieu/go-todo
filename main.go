package main

import (
    "github.com/LinTrieu/go-todo/router"
    "log"
    "net/http"
    "fmt"
)

func main() {
    r := router.Router()

    fmt.Println("Starting server on port 8080.")

  	log.Fatal(http.ListenAndServe(":8080", r))
}
