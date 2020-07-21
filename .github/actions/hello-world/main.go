package main

import (
	"fmt"
	"os"
)

func main() {

  // Access Inputs as environment vars
  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONDGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")

  // Use those inputs in the actions logic
  fmt.Println("Hello " + firstGreeting)

  if secondGreeting != "" {
    fmt.Println("Hello " + secondGreeting)
  } else {
    fnt.Println("No second greeting")
  }

  // Someimes inputs are not "required" and we can build around that
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
  }


}
