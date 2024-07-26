package palindrome

/* Number 1 (Two pointers pattern) - 1.8ns */
func IsPalindromeFirst(s string) bool {
    left, right := 0, len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return false
        }

        left++
        right--
    }

    return true
}

/* Number 2 (LeetCode version) - 5.3ns */
func IsPalindromeSecond(s string) bool {
    left, right := 0, len(s)-1

    for left < right {
        for left < right && !isAlphanumeric(s[left]) {
            left++
        }
 
        for left < right && !isAlphanumeric(s[right]) {
            right--
        }

        if toLower(s[left]) != toLower(s[right]) {
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphanumeric(ch byte) bool {
    return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
    if ch >= 'A' && ch <= 'Z' {
        return ch + ('a' - 'A')
    }
    return ch
}

/* Number 3 (Dequeue: removing from a queue) - 35.1ns */
func IsPalindromeThird(s string) bool {
    deque := []rune{}
    
    for _, char := range s {
        deque = append(deque, char)
    }
    
    for len(deque) > 1 {
        if deque[0] != deque[len(deque)-1] {
            return false
        }

        deque = deque[1 : len(deque)-1]
    }
    
    return true
}