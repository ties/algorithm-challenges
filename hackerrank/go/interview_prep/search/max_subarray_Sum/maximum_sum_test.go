package main

import (
    "fmt"
    "testing"
)

func TestMaximumSum(t *testing.T) {
    res := maximumSum([]int64{3,3,9,9,5}, 7)

    if res != 6 {
        t.Error(fmt.Sprintf("sample input 0 result (%v) != 6", res))
    }
}
