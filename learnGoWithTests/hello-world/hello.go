package main

import "fmt"

func main() {
	fmt.Println(Hello("Hello", "Felix"))
}

func Hello(greeting, name string) string {
	if greeting == "" {
		greeting = "Hello"
	}
	if name == "" {
		name = "World"
	}
	return greeting + " " + name
}
