package reflection

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Profile Profile
	Game    Game
}

type Profile struct {
	Age  int
	City string
}

type Game struct {
	Title string
	Genre string
}

// func walk(x interface{}, fn func(input string)) {
// 	// Get the value of the 'x' variable
// 	val := reflect.ValueOf(x)
// 	fmt.Println("the reflect.ValueOf()", val)

// 	// extract the first field from the reflection value
// 	field := val.Field(0)
// 	fmt.Println("the val.Field()", val.Field(0))

// 	// fn("Ever since I was a jit, you know am the shit")

// 	// Display the field value as string
// 	fn(field.String())
// }

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// Cleaner, we do the thing based on the reflected type data
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	// Get the value of the 'x' variable
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	fmt.Println("the reflect.ValueOf()", val)
	return val
}
