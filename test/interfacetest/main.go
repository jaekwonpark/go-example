package main

import "fmt"

type AbstractModel struct {
	objectType string
	tenantId string
}


type ApiLink struct {
	href string
	rel string
	timeseries Timeseries
}


type ApiResponseMetadata struct {
	links ApiLink
}


type BoolVal struct {
	bool_val bool
	objectType string
	tenantId string
}


type DoubleVal struct {
	double_val float64
	objectType string
	tenantId string
}


type Hello struct {
	message string
}


type HelloApiResponse struct {
	data Hello
	metadata ApiResponseMetadata
}


type IntVal struct {
	int_val int32
	objectType string
	tenantId string
}


type Point struct {
	objectType string
	tenantId string
	timestamp_epoch int32
	value OneOfIntValStrValBoolValDoubleVal
}


type StrVal struct {
	objectType string
	str_val string
	tenantId string
}


type Timeseries struct {
	end_time_sec int32
	objectType string
	sampling_interval_secs int32
	start_time_sec int32
	tenantId string
	values Point
}


type OneOfIntValStrValBoolValDoubleVal interface {
	OneOfIntValStrValBoolValDoubleVal()
}
func (*BoolVal) OneOfIntValStrValBoolValDoubleVal() {}
func (*StrVal) OneOfIntValStrValBoolValDoubleVal() {}
func (*IntVal) OneOfIntValStrValBoolValDoubleVal() {}
func (*DoubleVal) OneOfIntValStrValBoolValDoubleVal() {}


func myfunc(o OneOfIntValStrValBoolValDoubleVal) {
	switch v := o.(type) {
	case *BoolVal:
		fmt.Printf("BoolVal: %v", v)
	case *StrVal:
		fmt.Printf("StrVal: %v", v)
	default:
		fmt.Printf("Unknown")
	}
}

type Message struct {
	categories []string
}
type mapMessages map[string]Message
type Myarray []string

var a mapMessages
var b Myarray



func main() {

	r := StrVal{"a", "b", "c"}
	c := BoolVal{ true, "d", "e"}

	myfunc(&r)
	myfunc(&c)

	a = make(mapMessages)
	a["abc"] = Message{
		categories: []string{"a","b"}}

	b = append(b, "aaa")
	b = append(b, "bbb")

	fmt.Printf("a:%v\n", a)
	fmt.Printf("b:%v\n", b)


}

