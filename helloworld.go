package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
        log.Print("Hello world received a request.")
        target := os.Getenv("TARGET")
        port := os.Getenv("PORT")
        if target == "" {
                target = "World"
        }
        fmt.Fprintf(w, "Hello %s on %s!\n", target, port)
}

func main() {
        log.Print("Hello world sample started.")

        http.HandleFunc("/", handler)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
