package utils
import (
    "regexp"
    "strings"
)

// ToSnakeCase converts "MyMethodName" or "myMethodName" to "my_method_name"
func ToSnakeCase(input string) string {
    var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
    var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

    snake := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
    snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}


func JoinWithPrefix(prefix string, parts ...string) string {
    return strings.Join(append([]string{prefix}, parts...), ".")
}