package main

import (
  "encoding/json"
  "fmt"
  "log"
  "reflect"
)

type Foo struct {
  P1 string `json:"p1,omitempty"`
  P2 string `json:"p2,omitempty"`
}

type Bar struct {
  Name string `required:"true"`
  Age  int    `required:"true"`
}

func (b *Bar) Unmarshal(data []byte) error {
  err := json.Unmarshal(data, b)
  if nil != err {
    return err
  }

  fields := reflect.ValueOf(b).Elem()
  numField := fields.NumField()
  for i := 0; i < numField; i++ {

    required := fields.Type().Field(i).Tag.Get("required")
    name := fields.Type().Field(i).Name
    if required == "true" {
      field := fields.Field(i)
      if field.IsZero() {
        return fmt.Errorf("Required field %s is missing", name)
      }
    }
  }
  return nil
}

func main() {

  //profile1 := `{"Name":"foo", "Age":20}`
  profile2 := `{"Name": "name2", "Age2":21}`

  var profile Bar

  /*err := profile.Unmarshal([]byte(profile1))
  if err != nil {
    log.Printf("profile1 unmarshal error: %s\n", err.Error())
    return
  }*/
  //fmt.Printf("profile1 unmarshal: %v\n", profile)

  err := profile.Unmarshal([]byte(profile2))
  if err != nil {
    log.Printf("profile2 unmarshal error: %s\n", err.Error())
    return
  }
  fmt.Printf("profile2 unmarshal: %v\n", profile)

}
