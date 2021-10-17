// We often need our programs to perform operations on collections of data, like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.

// In some languages it’s idiomatic to use generic data structures and algorithms. Go does not support generics; in Go it’s common to provide collection functions if and when they are specifically needed for your program and data types.

// Here are some example collection functions for slices of strings. You can use these examples to build your own functions. Note that in some cases it may be clearest to just inline the collection-manipulating code directly, instead of creating and calling a helper function.
package main

// Index returns the first index of the target string t, or -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
// Include returns true if the target string t is in the slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}
// Any returns true if one of the strings in the slice satisfies the predicate f.
