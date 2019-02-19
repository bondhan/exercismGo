/*
Package bob implements the test for "bob"
*/
package bob

import (
	"strings"
	"regexp"
	// "fmt"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if (remark == ""){
		return "Fine. Be that way!"
	}

	regExp := regexp.MustCompile(`^[0-9,:)\s]*$`)

	words:= remark[:len(remark)-1]
	lastChar := remark[len(remark)-1:]
	
	switch(lastChar){
	case ".":
		return "Whatever."
	case "?":
		isNumberOrComma := regExp.MatchString(words)
		// fmt.Println("processed = ", remark, "-", isNumberOrComma)
		if (strings.ToUpper(words) == words && !isNumberOrComma){
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case "!":
		if (strings.ToUpper(words) == words){
			return "Whoa, chill out!"
		}
		return "Whatever."
	default:
		isNumberOrComma := regExp.MatchString(remark)
		// fmt.Println("processed = ", remark, "-", isNumberOrComma)

		if (isNumberOrComma) {
			return "Whatever."
		} else if (strings.ToUpper(words) == words){
			return "Whoa, chill out!"
		} else {
			return "Whatever."
		}

	}

	return ""
}
