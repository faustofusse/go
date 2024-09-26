package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintFloat(value *float64, fallback string) string {
    if value == nil {
        return fallback
    } else if *value - float64(int(*value)) == 0 {
        return fmt.Sprintf("%v", int(*value))
    } else {
        return fmt.Sprintf("%.2f", *value)
    }
}

func PrintFullFloat(value *float64, fallback string) string {
    if value == nil {
        return fallback
    } else if *value - float64(int(*value)) == 0 {
        return fmt.Sprintf("%v", int(*value))
    } else {
        return fmt.Sprintf("%f", *value)
    }
}

type Int64 int64

func (value *Int64) UnmarshalJSON(data []byte) error {
    // Ignore null, like in the main JSON package.
    if string(data) == "null" || string(data) == `""` {
        *value = 0
        return nil
    }
    str := strings.ReplaceAll(string(data), "\"", "")
    parsed, err := strconv.ParseInt(str, 10, 64)
    if err != nil {
        *value = 0
        return nil
    }
    *value = Int64(parsed)
    return nil
}

func (value *Int64) Equals(other int64) bool {
    return *value == Int64(other)
}

type Float64 float64

func (value *Float64) Reference() *float64 {
    if value == nil {
        return nil
    }
    parsed := float64(*value)
    return &parsed
}

func (value *Float64) UnmarshalJSON(data []byte) error {
    // Ignore null, like in the main JSON package.
    if string(data) == "null" || string(data) == `""` {
        *value = 0
        return nil
    }
    str := strings.ReplaceAll(string(data), "\"", "")
    parsed, err := strconv.ParseFloat(str, 64)
    if err != nil {
        *value = 0
        return nil
    }
    *value = Float64(parsed)
    return nil
}

func (value *Float64) Equals(other float64) bool {
    return *value == Float64(other)
}
