package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "alfa alfa alfa alfa alfa"
	str = repleseNth(str, "alfa", "beta", 3)
	fmt.Println(str)

}
func repleseNth(s, old, new string, n int) string {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			// we did not find it
			break
		}

		// have found it
		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i += len(old)
	}
	return s

}
