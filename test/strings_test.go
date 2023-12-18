package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestHasPrefix(t *testing.T) {
	fmt.Println(strings.HasPrefix("my string", "prefix")) // false
	fmt.Println(strings.HasPrefix("my string", "my"))     // true
	fmt.Println(strings.HasPrefix("/api/typicode/getList", "/api/typicode"))

}
