package main

import (
	"fmt"
	"strings"

	"github.com/SantiSite/module-example/slices"
)

func main() {
	list := []string{"EDteam", "gophers", "golang", "hello"}

	list = slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

	fmt.Println("list:", list)
}
