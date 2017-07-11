package main

import (
	"flag"
)

func main() {
	var key = flag.String("key", "", "API key for the ifttt maker channel.")
	flag.Parse()

	maker := NewMaker(*key)
	maker.Send("test4", map[string]string{"foo": "bar"})
}
