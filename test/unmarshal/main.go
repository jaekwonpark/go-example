package main

import (
  "encoding/json"
  "fmt"
  "reflect"
)

type Foo struct {
  P1 *string `json:"p1,omitempty"`
  P2 *string `json:"p2,omitempty"`
}

type Bar struct {
  // use custom tag "required"
  P3 *int `json:"p3" required:"true"`
}

type BarWrapper struct {
  Bar
}

func (b *BarWrapper) UnmarshalJSON(data []byte) error {

  // use inner struct to get the value
  fields := reflect.ValueOf(&b.Bar).Elem()
  numField := fields.NumField()

  // for each field, check if "required" tag is "true" and the value is missing
  for i := 0; i < numField; i++ {
    field := fields.Type().Field(i)
    fmt.Printf("Checking field %s with type %s\n", field.Name, field.Type)
    required := field.Tag.Get("required")
    // we check if field is nil because all fields in generated code is pointer type
    if required == "true" && fields.Field(i).IsNil() {
      return fmt.Errorf("Required field %s with type %s is missing\n", field.Name, field.Type)
    }
  }

  // Unmarshal inner struct
  if err := json.Unmarshal(data, &b.Bar); nil != err {
    return fmt.Errorf("Unable to unmarshal: %s", err)
  }
  return nil
}

func main() {
  foo := new(Foo)
  foo.P1 = new(string)
  foo.P2 = new(string)
  *foo.P1 = "prop1"
  *foo.P2 = "prop2"
  strFoo, _ := json.Marshal(foo)
  fmt.Println(string(strFoo))

  bar := new(BarWrapper)

  if err := json.Unmarshal(strFoo, bar); err != nil {
    fmt.Println("\nFailure")
    fmt.Println(err)
  } else {
    fmt.Println("\nSuccess")
    fmt.Printf("bar.p3:%d\n", *bar.P3)
  }
}

