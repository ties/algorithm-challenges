package main

import (
    "fmt"
    "testing"
)

func TestTiplets(t *testing.T) {
    res := triplets([]int32{1,3,5}, []int32{2,3}, []int32{1,2,3})

    if res != 8 {
        t.Error(fmt.Sprintf("Result (%v) != 8", res))
    }

    res = triplets([]int32{1,4,5}, []int32{2,3,3}, []int32{1,2,3})

    if res != 5 {
        t.Error("Result != 5")
   }

    res = triplets([]int32{1,3,5,7}, []int32{5,7,9}, []int32{7,9,11,13})

    if res != 12 {
        t.Error("Result != 12")
    }
}

