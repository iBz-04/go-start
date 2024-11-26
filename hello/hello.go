package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // request a greeting message.
    message, err := greetings.Hello("Ibrah")

    if err != nil {
        log.Fatal(err)
    }

    // Get a greeting message and print it.
    fmt.Println(message)
}