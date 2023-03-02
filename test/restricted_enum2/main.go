package main

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
)

type Customer struct {
  Name *string
  Address *string
}
type Pet struct {
  Name *string
  Age int
  Owner *Customer
}
type OpenDays struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  Pet *Pet `json:"pet,omitempty"`
  WorkDays *WeekDays `json:"workDays,omitempty"`
}
func (e *OpenDays) UnmarshalJSON(b []byte) error {
  var raw map[string]interface{}
  if err := json.Unmarshal(b, &raw); err != nil {
    return err
  }
  for k, v := range raw {
    switch k {
    case "Pet":
      if bi, err := json.Marshal(v); err == nil {
        fmt.Printf("Marshalled:%s\n", string(bi))
        tmpPet := &Pet{}
        if err := json.Unmarshal(bi, tmpPet); err == nil {
          fmt.Printf("Unmarshalled:tmpPet:%v", tmpPet)
        } else {
          fmt.Printf("error while unmarshalling e.Pet:%s", err)
        }
      }
      return errors.New("failed to unmarshal Pet")
    case "WorkDays":
      wd := v.(string)
      if wd == "SUNDAY" || wd == "SATURDAY" {
        wd = "NOT VALID"
        fmt.Printf("%s is not valid", wd)
      }
    }
    fmt.Println("Key:%s, Val:%v\n", k, v)
  }

  return nil
}

type WeekDays int

const(
  WEEKDAYS_MONDAY WeekDays = 0
  WEEKDAYS_TUESDAY WeekDays = 1
  WEEKDAYS_WEDNESDAY WeekDays = 2
  WEEKDAYS_THURSDAY WeekDays = 3
  WEEKDAYS_FRIDAY WeekDays = 4
  WEEKDAYS_SATURDAY WeekDays = 5
  WEEKDAYS_SUNDAY WeekDays = 6
  WEEKDAYS_UNKNOWN WeekDays = 7
  WEEKDAYS_REDACTED WeekDays = 8
)

// returns the name of the enum given an ordinal number
func (e *WeekDays) name(index int) string {
  names := [...]string {
    "MONDAY",
    "TUESDAY",
    "WEDNESDAY",
    "THURSDAY",
    "FRIDAY",
    "SATURDAY",
    "SUNDAY",
    "$UNKNOWN",
    "$REDACTED",
  }
  if index < 0 || index > len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *WeekDays) index(name string) WeekDays {
  names := [...]string {
    "MONDAY",
    "TUESDAY",
    "WEDNESDAY",
    "THURSDAY",
    "FRIDAY",
    "SATURDAY",
    "SUNDAY",
    "$UNKNOWN",
    "$REDACTED",
  }
  for idx := range names {
    if names[idx] == name {
      return WeekDays(idx)
    }
  }
  return WEEKDAYS_UNKNOWN
}

func (e *WeekDays) UnmarshalJSON(b []byte) error {
  var enumStr string
  fmt.Printf("Unmarshal WeekDays\n")
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for WeekDays:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *WeekDays) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func main() {
  byt := []byte(`{"ObjectType_": "OpenDays", "WorkDays": "SATURDAY", "Pet": {"Name":"Puppy", "Age":1, "Owner": {"Name": "John", "Address": "This location"} } }`)
  var openDays OpenDays
  if err := json.Unmarshal(byt, &openDays); err != nil {
    fmt.Printf("openDays:%v\n", openDays)
    panic(err)
  }
}
