package utils

import "time"

func Now() *int64{
    now := new(int64)
    *now = time.Now().Unix()
    return now
}
