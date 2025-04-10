package wrapt

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)


func Wrap(s string) string {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
	}	

	if len(s) < w {
		return s
	}

	pos := 0
	// we're going to set an upper bound for this loop
	for range 10000 {
		// if we don't contain a breaking character, return early
		if !strings.Contains(s[pos:], " ") {
			break
		}
		if pos + w > len(s) - 1 {
			break
		} 	
		pos += w

		pos = lastSpaceBeforePosition(s, pos)
		s = insertBreak(s, pos + 1)
	}

	return s
}

func lastSpaceBeforePosition(s string, pos int) int {
    for i := pos - 1; i >= 0; i-- {
        if s[i] == ' ' {
            return i
        }
    }
    
    return -1
}

func insertBreak(s string, pos int) string {
    before := s[:pos]
    after := s[pos:]

	return before + "\n" + after
}
