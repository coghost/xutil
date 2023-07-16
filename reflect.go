package xutil

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/fatih/structs"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

// IsZeroVal check if any type is its zero value
func IsZeroVal(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// IsDefaultVal alias of IsZeroVal
func IsDefaultVal(x interface{}) bool {
	return IsZeroVal(x)
}

func GetKind(o interface{}) reflect.Kind {
	return reflect.TypeOf(o).Elem().Kind()
}

// GetAttr beware `o` should be the pointer
func GetAttr(o interface{}, key string) reflect.Value {
	v := reflect.ValueOf(o).Elem().FieldByName(key)
	return v
}

// SetAttr will set o.key to val
// same as GetAttr, o should be a pointer
func SetAttr(o interface{}, key string, val interface{}) {
	f := GetAttr(o, key)
	setVal(f, val)
}

// SetField: Set field to val with val's type dynamically
func setVal(field reflect.Value, val interface{}) {
	switch vt := val.(type) {
	case bool:
		field.SetBool(val.(bool))
	case []byte:
		field.SetBytes(val.([]byte))
	case uint64:
		field.SetUint(val.(uint64))
	case int:
		field.SetInt(int64(val.(int)))
	case int64:
		field.SetInt(val.(int64))
	case float64:
		field.SetFloat(val.(float64))
	case string:
		field.SetString(val.(string))
	default:
		panic(fmt.Sprintf("Unsupported type %v", vt))
	}
}

// Caller wraps runtime.Caller and returns file and line number information
//
// Returns: filename, linenum
func Caller(skip int) (string, int) {
	pc, file, l, _ := runtime.Caller(skip)

	lst := strings.Split(file, "/")
	file = lst[len(lst)-1]

	funcName := runtime.FuncForPC(pc).Name()
	lst = strings.Split(funcName, "/")
	funcName = fmt.Sprintf("%s:%s", file, lst[len(lst)-1])

	return funcName, l
}

func GetAllTags(structIn interface{}, tag string) (tags []string) {
	fields := structs.Fields(structIn)
	for _, field := range fields {
		tag := field.Tag(tag)
		tags = append(tags, tag)
	}
	return tags
}

func CallerStack() []byte {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return buf[:n]
}

func QuietGetKeys(obj interface{}) []string {
	if obj == nil {
		return []string{}
	}
	keys, ok := funk.Keys(obj).([]string)
	if !ok {
		return []string{}
	}
	return keys
}

func MustGetKeys(obj interface{}) []string {
	if obj == nil {
		return []string{}
	}
	keys, ok := funk.Keys(obj).([]string)
	if !ok {
		log.Fatal().Msg("cannot get key from object")
	}
	return keys
}
