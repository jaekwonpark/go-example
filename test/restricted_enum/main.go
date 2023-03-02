package main

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
)

type OpenDays struct {
  ObjectType_ *string `json:"$objectType,omitempty"`
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  WorkDays *WrapperWeekDays `json:"workDays,omitempty"`
}

type WrapperWeekDays struct {
  WeekDays
}

func (e *WrapperWeekDays) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for WrapperWeekDays:%s", err))
  }

  wd := new(WeekDays)
  if err := json.Unmarshal(b, &wd); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for WeekDays:%s", err))
  }
  switch *wd {
  case WEEKDAYS_SATURDAY:
    return errors.New(fmt.Sprintf("Invalid WeekDays:%v", *wd))
  case WEEKDAYS_SUNDAY:
    return errors.New(fmt.Sprintf("Invalid WeekDays:%v", *wd))

  }
  e.WeekDays = *wd
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
  byt := []byte(`{"WorkDays": "SATURDAY"}`)
  var openDays OpenDays
  if err := json.Unmarshal(byt, &openDays); err != nil {
    panic(err)
  }
  fmt.Println(openDays.WorkDays)


}
