package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func main() {
	fmt.Printf("%s\n", html.UnescapeString("&nbsp;"))
}
