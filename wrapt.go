package wrapt

import (
	"strings"
)

// text: string to be wrapped
// terminalWidth: terminal width
// Returns a string, with line breaks for each line
func Wrap(text string, terminalWidth int) string {
	if len(text) < terminalWidth {
		return text
	}

	pos := 0
	// we're going to set an upper bound for this loop
	for range 10000 {
		// if we don't contain a breaking character, return early
		if !strings.Contains(text[pos:], " ") {
			break
		}
		if pos + terminalWidth > len(text) - 1 {
			break
		} 	
		pos += terminalWidth

		pos = lastSpaceBeforePosition(text, pos)
		text = insertBreak(text, pos + 1)
	}

	return text
}

// text: string to be wrapped
// marginSize: the margin to be applied on either side of the printed string
// terminalWidth: terminal width
// Returns an array of strings, representing each line to be printed
func WrapArray(text string, marginSize int, terminalWidth int) []string {
    terminalWidth -= marginSize

    arr := []string{}
    if len(text) < terminalWidth {
        return append(arr, text)
    }

    startingPos := 0
    
    for range 10000 {
        if startingPos >= len(text) {
            break
        }
        
        endingPos := startingPos + terminalWidth
        if endingPos >= len(text) {
            arr = append(arr, text[startingPos:])
            break
        }
        
        endingPos = lastSpaceBeforePosition(text, endingPos)
        if endingPos == -1 || endingPos <= startingPos {
            arr = append(arr, text[startingPos:])
            break
        }
        
        line := text[startingPos:endingPos]
        arr = append(arr, line)
        startingPos = endingPos + 1
    }
    
    return arr
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
