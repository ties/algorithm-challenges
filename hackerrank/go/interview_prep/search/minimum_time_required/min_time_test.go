package main

import (
    "fmt"
    "testing"
)

func TestMinTime(t *testing.T) {
    res := minTime([]int64{2,3}, 5)

    if res != 6 {
        t.Error(fmt.Sprintf("sample input 0 result (%v) != 6", res))
    }

    res = minTime([]int64{1,3,4}, 10)

    if res != 7 {
        t.Error(fmt.Sprintf("sample input 1 (%v) != 7", res))
    }

    res = minTime([]int64{4,5,6}, 12)

    if res != 12 {
        t.Error(fmt.Sprintf("sample input 2 result (%v) != 20", res))
    }
}

