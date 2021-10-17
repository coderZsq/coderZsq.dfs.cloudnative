/**
Introduction

JSON (JavaScript Object Notation) is a simple data interchange format. Syntactically it resembles the objects and lists of JavaScript. It is most commonly used for communication between web back-ends and JavaScript programs running in the browser, but it is used in many other places, too. Its home page, json.org, provides a wonderfully clear and concise definition of the standard.

With the json package it’s a snap to read and write JSON data from your Go programs.
*/
package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Given the Go data structure, Message,
type Message struct {
	Name string
	Body string
	Time int64
}

// Encoding
func init() {
	// To encode JSON data we use the Marshal function.
	// func Marshal(v interface{}) ([]byte, error)
	// and an instance of Message
	m := Message{"Alice", "Hello", 1294706395881547000}
	// we can marshal a JSON-encoded version of m using json.Marshal:
	b, err := json.Marshal(m)
	// If all is well, err will be nil and b will be a []byte containing this JSON data:
	fmt.Printf("%s, err: %v\n", b, err)
	// b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

	// Only data structures that can be represented as valid JSON will be encoded:
	// * JSON objects only support strings as keys; to encode a Go map type it must be of the form map[string]T (where T is any Go type supported by the json package).
	// * Channel, complex, and function types cannot be encoded.
	// * Cyclic data structures are not supported; they will cause Marshal to go into an infinite loop.
	// * Pointers will be encoded as the values they point to (or ‘null’ if the pointer is nil).
	// The json package only accesses the exported fields of struct types (those that begin with an uppercase letter). Therefore only the the exported fields of a struct will be present in the JSON output.

	// Decoding
	// To decode JSON data we use the Unmarshal function.
	// func Unmarshal(data []byte, v interface{}) error
	// We must first create a place where the decoded data will be stored
	var m2 Message
	// and call json.Unmarshal, passing it a []byte of JSON data and a pointer to m
	err = json.Unmarshal(b, &m2)
	// If b contains valid JSON that fits in m, after the call err will be nil and the data from b will have been stored in the struct m, as if by an assignment like:
	fmt.Printf("%+v, err: %v\n", m2, err)
	//m = Message{
	//	Name: "Alice",
	//	Body: "Hello",
	//	Time: 1294706395881547000,
	//}
	// How does Unmarshal identify the fields in which to store the decoded data? For a given JSON key "Foo", Unmarshal will look through the destination struct’s fields to find (in order of preference):
	// * An exported field with a tag of "Foo" (see the Go spec for more on struct tags),
	// * An exported field named "Foo", or
	// * An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".
	// What happens when the structure of the JSON data doesn’t exactly match the Go type?
	b = []byte(`{"Name": "Bob", "Food": "Pickle"}`)
	var m3 Message
	err = json.Unmarshal(b, &m3)
	fmt.Printf("%+v, err: %v\n", m3, err)
	// Unmarshal will decode only the fields that it can find in the destination type. In this case, only the Name field of m will be populated, and the Food field will be ignored. This behavior is particularly useful when you wish to pick only a few specific fields out of a large JSON blob. It also means that any unexported fields in the destination struct will be unaffected by Unmarshal.
	// But what if you don’t know the structure of your JSON data beforehand?
}

// Generic JSON with interface
func init() {
	// The interface{} (empty interface) type describes an interface with zero methods. Every Go type implements at least zero methods and therefore satisfies the empty interface.
	// The empty interface serves as a general container type:
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	// A type assertion accesses the underlying concrete type:
	r := i.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)
	// Or, if the underlying type is unknown, a type switch determines the type:
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v* 2)
	case float64:
		fmt.Println("the reciprocal of is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}
	// The json package uses map[string]interface{} and []interface{} values to store arbitrary JSON objects and arrays; it will happily unmarshal any valid JSON blob into a plain interface{} value. The default concrete Go types are:
	// * bool for JSON booleans,
	// * float64 for JSON numbers,
	// * string for JSON strings, and
	// * nil for JSON null.
}

// Decoding arbitrary data
func init() {
	// Consider this JSON data, stored in the variable b:
	//b := []byte{`{"Name": "Wed"}`}
}

func main() {}
