package utils

import "testing"

func TestToSnakeCase(t *testing.T) {
    tests := map[string]string{
        "MyMethodName":    "my_method_name",
        "simpleTest":      "simple_test",
        "HTTPRequest":     "http_request",
        "GetURLParams":    "get_url_params",
        "IDNumberCheck":   "id_number_check",
    }

    for input, expected := range tests {
        if result := ToSnakeCase(input); result != expected {
            t.Errorf("ToSnakeCase(%q) = %q; want %q", input, result, expected)
        }
    }
}
