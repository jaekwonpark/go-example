package main

import "fmt"

func main() {
  var discriminator *string

  var str = "test"
  discriminator = new(string)
  *discriminator = str
  fmt.Printf("dscriminator:%s\n", discriminator)

}
