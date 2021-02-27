package main

import (
    "fmt"
    "testing"
)

func TestMinimumPasses(t *testing.T) {
    // input 0
    res := minimumPasses(3, 1, 2, 12)

    if res != 3 {
        t.Error(fmt.Sprintf("Incorrect result, %v != %v", 3, res))
    }
    // input 1
    res = minimumPasses(1, 1, 6, 45)
    if res != 16 {
        t.Error(fmt.Sprintf("Incorrect result, %v != %v", 16, res))
    }

    // input 2
    res = minimumPasses(5184889632, 5184889632, 20, 10000)
    if res != 1 {
        t.Error(fmt.Sprintf("Incorrect result, %v != %v", 1, res))
    }

}
