package di

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s\n", name)
}
