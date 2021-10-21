/**
The interface{} type

The interface{} type, the empty interface, is the source of much confusion. The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface. That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value. So, this function:

 */
package main

import "fmt"

func DoSomething(v interface{})  {
	// ...
}
// will accept any parameter whatsoever.

// Here’s where it gets confusing: inside of the DoSomething function, what is v’s type? Beginner gophers are led to believe that “v is of any type”, but that is wrong. v is not of any type; it is of interface{} type. Wait, what? When passing a value into the DoSomething function, the Go runtime will perform a type conversion (if necessary), and convert the value to an interface{} value. All values have exactly one type at runtime, and v’s one static type is interface{}.

// This should leave you wondering: ok, so if a conversion is taking place, what is actually being passed into a function that takes an interface{} value (or, what is actually stored in an []Animal slice)? An interface value is constructed of two words of data; one word is used to point to a method table for the value’s underlying type, and the other word is used to point to the actual data being held by that value. I don’t want to bleat on about this endlessly. If you understand that an interface value is two words wide and it contains a pointer to the underlying data, that’s typically enough to avoid common pitfalls. If you are curious to learn more about the implementation of interfaces, I think Russ Cox’s description of interfaces is very, very helpful.

// In our previous example, when we constructed a slice of Animal values, we did not have to say something onerous like Animal(Dog{}) to put a value of type Dog into the slice of Animal values, because the conversion was handled for us automatically. Within the animals slice, each element is of Animal type, but our different values have different underlying types.

// So… why does this matter? Well, understanding how interfaces are represented in memory makes some potentially confusing things very obvious. For example, the question “can I convert a []T to an []interface{}” is easy to answer once you understand how interfaces are represented in memory. Here’s an example of some broken code that is representative of a common misunderstanding of the interface{} type:

func PrintAll(vals []interface{})  {
	for _, val := range vals{
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	//PrintAll(names) // cannot use names (type []string) as type []interface {} in argument to PrintAll
	// By running this, you can see that we encounter the following error: cannot use names (type []string) as type []interface {} in function argument. If we want to actually make that work, we would have to convert the []string to an []interface{}:
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
	// Run it here:http://play.golang.org/p/Dhg1YS6BJS
	// That’s pretty ugly, but c'est la vie. Not everything is perfect. (in reality, this doesn’t come up very often, because []interface{} turns out to be less useful than you would initially expect)
}