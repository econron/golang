package main

import (
	"fmt"
	"reflect"
)

func inspectFields(input interface{}) {
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	
	switch val.Kind() {
	case reflect.Struct:
		fmt.Println("Struct fields and values:")
		// Iterate over the fields of the struct
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			value := val.Field(i)

			// Handle pointer fields
			if value.Kind() == reflect.Ptr && !value.IsNil() {
				value = value.Elem()
			}

			fmt.Printf("%s: %v\n", field.Name, value)
		}
	case reflect.Map:
		fmt.Println("Map keys and values:")
		// Iterate over the keys of the map
		for _, key := range val.MapKeys() {
			value := val.MapIndex(key)
			fmt.Printf("%v: %v\n", key, value)
		}
	default:
		fmt.Println("Unsupported type. Please provide a struct or a map.")
	}
}

func main() {
	// Test with a struct containing pointer fields
	type TestStruct struct {
		Name  *string
		Age   int
		Email *string
	}

	name := "John"
	// email := "john@example.com"
	structInstance := &TestStruct{
		Name:  &name,
		Age:   30,
		Email: nil,
	}

	// Test with a map
	mapInstance := map[string]interface{}{
		"Name":  "Jane",
		"Age":   25,
		"Email": nil,
	}

	// Inspect the struct
	fmt.Println("Inspecting struct:")
	inspectFields(structInstance)

	// Inspect the map
	fmt.Println("\nInspecting map:")
	inspectFields(mapInstance)
}