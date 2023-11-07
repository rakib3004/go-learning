package main

import "fmt"

type Person struct {
    Name   string
    Id    string

}

func (p Person) intro() {
    fmt.Printf("Intro of %s", p.Name)
}

func main() {
    person := Person{
        Name:  "Rakib Trofder",
        Id:   "a4b3c2012443",
    }

    person.intro()
}
