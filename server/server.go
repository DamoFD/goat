package server

import (
    "fmt"
    "net/http"
)

func Start() {
    router := NewRouter()

    fs := http.FileServer(http.Dir("./static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))


    fmt.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
