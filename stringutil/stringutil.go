// File: stringutil/stringutil.go

package stringutil

import "strings"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ToUpper(s string) string {
    return strings.ToUpper(s)
}

func Concat(strs ...string) string {
    return strings.Join(strs, "")
}

func IsEmpty(s string) bool {
    return len(s) == 0
}

func FirstRune(s string) rune {
    if len(s) == 0 {
        return 0
    }
    return []rune(s)[0]
}
