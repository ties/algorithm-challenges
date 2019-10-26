package main

import (
    "testing"
)

func TestMedianEvenEqual(t *testing.T) {
    buffer := []int{1,2,2,4}
    res := Median(buffer)

    if res != 2.0 {
	t.Error("Result != 2.0")
    }
}

func TestMedianEvenNotEqual(t *testing.T) {
    buffer := []int{1,2,3,4}
    res := Median(buffer)

    if res != 2.5 {
	t.Error("Result != 2.5")
    }
}

func TestMedianOdd(t *testing.T) {
    buffer := []int{1,3,5}
    res := Median(buffer)

    if res != 3.0 {
	t.Error("Result != 3.0")
    }
}
