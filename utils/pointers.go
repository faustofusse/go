package utils

func AddressOf[T any](value T) *T {
    return &value
}

func Dereference[T any](value *T, fallback T) T {
    if value == nil {
        return fallback
    } else {
        return *value
    }
    // switch value.(type) {
    //     case *string:
    //         if value == nil { return "" }
    //         return *(value.(*string))
    //     case string: return value.(string)
    //     case *int64:
    //         if value == nil { return "" }
    //         return fmt.Sprint(*(value.(*int64)))
    //     case int64: return fmt.Sprint(value.(int64))
    //     default: return ""
    // }
}
