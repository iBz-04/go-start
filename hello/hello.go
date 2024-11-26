package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // slice of names 
    names := []string{"Ibrah", "fauzan", "Sidik"}


    // request a greeting message.
    messages, err := greetings.Hellos(names)


    if err != nil {
        log.Fatal(err)
    }

    // Get a greeting message and print it.
    fmt.Println(messages)
}