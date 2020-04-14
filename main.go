package main

import (
    "net/http"
    "log"
    "DHTWeather/controllers"
)

func main() {
    http.HandleFunc("/", controllers.SayhelloName) // set router
    http.HandleFunc("/readings", controllers.ReadingsHandler)

    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
          log.Fatal("ListenAndServe: ", err)
    }
}

