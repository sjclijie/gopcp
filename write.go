package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var b bytes.Buffer

	b.Write([]byte( "hello" ))

	fmt.Fprint(&b, " world")

	b.WriteTo( os.Stdout )
}
