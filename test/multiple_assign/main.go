package main

import (
  "fmt"
)

type Type string

const (
  Service Type = "Service"
  Application Type = "Application"
)

func main() {
  var v1 int
  v1 = 4

  fmt.Printf("v1=%d(%p)\n", v1, &v1)

  v1, v2, _ := addtwothree(v1)
  /*if v1, v2 := addtwothree(v1); v2 != 0 {
    fmt.Printf("v2 is 0\n")
  }*/

  fmt.Printf("v1:%d(%p), v2:%d(%p)\n", v1, &v1, v2, &v2)

  s := Service
  a := Application

  fmt.Printf("s=%s(%T), a=%s(%T)\n", s, s, a, a)


}

func addtwothree(v1 int) (int, int, error) {
  return v1+2, v1+3, fmt.Errorf("err")
}
