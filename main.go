package main

import "fmt"

func main() {
  
  fmt.Println("1+1=", 1+1)

  fmt.Println("7/3", 7/3)
  fmt.Println("7.0/3.0", 7.0/3.0)
  
  fmt.Println("AND")
  fmt.Println(true && false)
  fmt.Println(true && true)
  fmt.Println(false && false)

  fmt.Println("OR")
  fmt.Println(true || false)
  fmt.Println(true || true)
  fmt.Println(false || false)

  fmt.Println("NOT")
  fmt.Println(!true)
  fmt.Println(!false)

}