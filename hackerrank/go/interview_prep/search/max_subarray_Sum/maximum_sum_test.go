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

    res = maximumSum([]int64{1,5,9}, 5)

    if res != 4 {
        t.Error(fmt.Sprintf("test case 18 result (%v) != 4", res))
    }
    res = maximumSum([]int64{3,2,7,4}, 7)
}
