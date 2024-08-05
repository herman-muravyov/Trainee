package stringers

import (
	"bytes"
	"strings"
)

func ReverseWordsFirst(s string) string {
    if len(s) < 1 {
        return ""
    }

    runes := []rune(s)
    i, j := 0, 0
    for i < len(runes) && j < len(runes) {
        for i < len(runes) && runes[i] == ' ' {
            i++
        }

        for j < len(runes) && runes[j] != ' ' {
            j++
        }
        
        k := j - 1
        for i < k {
            runes[i], runes[k] = runes[k], runes[i]
            i++
            k--
        }
        
        for j < len(runes) && runes[j] == ' ' {
            j++
        }
        i = j
    }

    return string(runes)
}

func ReverseWordsSecond(s string) string {
	var buffer bytes.Buffer
	
	str := strings.Split(s, " ")
	for _, v := range str {
		v = " " + v
		for i := len(v) - 1; i >= 0; i-- {
			buffer.WriteByte(v[i])
		}
	}

	return strings.TrimSpace(buffer.String())
}