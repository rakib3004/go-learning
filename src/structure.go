package main

import "fmt"

type Person struct {
    Name   string
    Age    int
    Email  string
}

func (p Person) introduce() {
    fmt.Printf("Hi, my name is %s. I am %d years old. You can reach me at %s.\n", p.Name, p.Age, p.Email)
}

func main() {
    person := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john.doe@example.com",
    }

    person.introduce()
}
