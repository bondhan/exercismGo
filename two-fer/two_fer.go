/* this package is a library which implement 2 for one
*/
package twofer

import (
	"strings"
	// "fmt"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	
		// if strings.TrimSpace(name) == "" {
		// 	return fmt.Sprintf("One for %s, one for me.", "you")
		// } else {
		// 	return fmt.Sprintf("One for %s, one for me.", name)
		// }

		if strings.TrimSpace(name) == "" {
			return "One for you, one for me."
		} else {
			return "One for "+name+", one for me."
		}

		
		// aName := "you"
		// if strings.TrimSpace(name) != "" {
		// 	aName = name
		// } 

		// return "One for "+aName+", one for me."
}
