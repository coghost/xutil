package xutil

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// Stringify returns a string representation
func Stringify(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Structify returns the original representation
func Structify(data string, value interface{}) error {
	return json.Unmarshal([]byte(data), value)
}

func MustStructify(data string, op interface{}) {
	err := Structify(data, op)
	PanicIfErr(err)
}

func MustStringify(data interface{}) string {
	v, err := Stringify(data)
	PanicIfErr(err)
	return v
}

func MustStructToMap(data interface{}) map[string]interface{} {
	v := structs.Map(data)
	return v
}

// Map2StructureLoosely
// wrapper of mapstructure.NewDecoder
//
// this will convert map[interface{}]interface{} to struct
// WARN: remember to pass struct in config.Result
// e.g.
//
// WeaklyTypedInput: true,
// Result:           &st,
// TagName:          "structs", // not required, but you can always customize it
func Map2StructureLoosely(dat interface{}, config *mapstructure.DecoderConfig) {
	decoder, err := mapstructure.NewDecoder(config)
	PanicIfErr(err)
	err = decoder.Decode(dat)
	PanicIfErr(err)
}

// Map2StructureByDefault
//
// just a wrapper of commonly usage to mapstructure.DecoderConfig
func Map2StructureByDefault(dat interface{}, dstStruct interface{}) {
	cfg := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "structs",
		Result:           dstStruct,
	}
	Map2StructureLoosely(dat, cfg)
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// A2B
//
// convert any types to byte
func A2B(val interface{}) []byte {
	switch val := val.(type) {
	case []byte:
		return val
	default:
		return []byte(fmt.Sprintf("%v", val))
	}
}

// B2A
//
// convert byte to other type
func B2A(v []byte, args ...interface{}) interface{} {
	str := cast.ToString(v)
	if len(args) < 1 {
		return str
	}

	switch t := args[0].(type) {
	case bool:
		return IfaceAorB(cast.ToBool(str), t)
	case string:
		return IfaceAorB(cast.ToString(str), t)
	case nil:
		return cast.ToString(str) // no Default value supplied
	case int:
		return IfaceAorB(cast.ToInt(str), t)
	case int64:
		return IfaceAorB(cast.ToInt64(str), t)
	case float32:
		return IfaceAorB(cast.ToFloat32(str), t)
	case float64:
		return IfaceAorB(cast.ToFloat64(str), t)
	case []uint8:
		return v
	default:
		log.Printf("Unknown Type: %v\n", t)
		return FirstNonZero(string(str), "")
	}
}

// ArrS2B
//
// convert string array to byte array
func ArrS2B(fields ...string) [][]byte {
	bf := [][]byte{}
	for _, f := range fields {
		bf = append(bf, []byte(f))
	}
	return bf
}

// ArrB2A
//
// convert byte array to interface array
func ArrB2A(args ...[]byte) []interface{} {
	arr := []interface{}{}
	for _, v := range args {
		if v == nil {
			arr = append(arr, nil)
		} else {
			arr = append(arr, string(v))
		}
	}
	return arr
}

// ArrB2S
//
// convert byte array to string
func ArrB2S(args ...[]byte) []string {
	arr := []string{}
	for _, v := range args {
		arr = append(arr, string(v))
	}
	return arr
}

// ArrA2B
//
// convert interface array to byte array
func ArrA2B(fields ...interface{}) [][]byte {
	bf := [][]byte{}
	for _, f := range fields {
		bf = append(bf, A2B(f))
	}
	return bf
}
