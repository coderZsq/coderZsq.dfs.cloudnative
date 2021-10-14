/**
Go's interfaces—static, checked at compile time, dynamic when asked for—are, for me, the most exciting part of Go from a language design point of view. If I could export one feature of Go into other languages, it would be interfaces.

This post is my take on the implementation of interface values in the “gc” compilers: 6g, 8g, and 5g. Over at Airs, Ian Lance Taylor has written two posts about the implementation of interface values in gccgo. The implementations are more alike than different: the biggest difference is that this post has pictures.


Before looking at the implementation, let's get a sense of what it must support.
*/
package main

import (
	"net"
)

/**
Usage

Go's interfaces let you use duck typing like you would in a purely dynamic language like Python but still have the compiler catch obvious mistakes like passing an int where an object with a Read method was expected, or like calling the Read method with the wrong number of arguments. To use interfaces, first define the interface type (say, ReadCloser):
*/
type ReadCloser interface {
	Read(b []byte) (n int, err net.Error)
	Close()
}

// and then define your new function as taking a ReadCloser. For example, this function calls Read repeatedly to get all the data that was requested and then calls Close:
func ReadAndClose(r ReadCloser, buf []byte) (n int, err net.Error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}
