package slices

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func filter[T constraints.Ordered](list []T) {
	fmt.Println(list)
}
