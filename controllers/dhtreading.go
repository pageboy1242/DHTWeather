package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
    "DHTWeather/repositories"
    "DHTWeather/models"
)

func SayhelloName(w http.ResponseWriter, r *http.Request) {
    // TODO: Setup a dictionary with http request types and function pointers
    switch r.Method {
        case http.MethodGet:
            getResponse(w, r)
        case http.MethodPost:
            fmt.Println("Received POST request")
            postRequest(w, r)
        default:
            w.WriteHeader(http.StatusNotImplemented)
        }
}

func getResponse(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // parse arguments, you have to call this by yourself
    fmt.Println(r.Form) // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "<h1>Hello, World!</h1>") // send data to client side
}

func postRequest(w http.ResponseWriter, r *http.Request) {
    reading, err := parseRequest(r)
    
    rep:= repositories.NewDhtReadingRepository()
    fmt.Println("Saving")
    rerr := rep.Save(reading)
    if rerr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }
    return
}

func parseRequest(r *http.Request) (models.Reading, error) {
    fmt.Printf("Got Body %s\n", r.Body)
    dec := json.NewDecoder(r.Body)
    var dht models.Reading
    err := dec.Decode(&dht)
    fmt.Println(dht)
    
    if err != nil {
        return models.Reading{}, err
    }
    return dht, nil
}

func ReadingsHandler(w http.ResponseWriter, r *http.Request) {
    rep := repositories.NewDhtReadingRepository()
    rep.Close()
    
    fmt.Fprintf(w, "Readings")
}
