package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 	str := "alfa alfa alfa alfa alfa"
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
	*/

	//dealing with case
	myString := " This is a clear EXAMPLE of why we seach in one case only."

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Did not find it!")

	}

	// other case
	fmt.Println(strings.ToLower(myString))                //всі букви малі
	fmt.Println(strings.ToUpper(myString))                // всі букви великі
	fmt.Println(strings.Title(strings.ToLower(myString))) // всі слова починаються з велико\

}
