package utils

func Contains[T comparable](s []T, e T) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func Map[T any, K any](slice []T, transform func(T) K) []K {
    result := []K{}
    for _, v := range slice {
        result = append(result, transform(v))
    }
    return result
}

func Filter[T any](slice []T, condition func(T) bool) []T {
    result := []T{}
    for _, v := range slice {
        if condition(v) {
            result = append(result, v)
        }
    }
    return result
}
