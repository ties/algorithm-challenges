package main

import (
    "fmt"
    "testing"
    "reflect"
)

func TestAddRemoveEqualValue(t *testing.T) {
    buffer := []int{1,2,2,4}
    cp := make([]int, len(buffer))
    copy(cp, buffer)

    cp = RemoveAndAdd(cp, 1, 1)

    if !reflect.DeepEqual(buffer, cp) {
	t.Error(fmt.Sprintf("Different result after add+remove same value at index 0 (%v != %v)", cp, buffer))
    }

    cp = RemoveAndAdd(cp, 2, 2)

    if !reflect.DeepEqual(buffer, cp) {
	t.Error(fmt.Sprintf("Different result after add+remove same value at index 0 (%v != %v)", cp, buffer))
    }
    
    cp = RemoveAndAdd(cp, 4, 4)

    if !reflect.DeepEqual(buffer, cp) {
	t.Error(fmt.Sprintf("Different result after add+remove same value at index 0 (%v != %v)", cp, buffer))
    }
}

func TestAddRemoveDifferingValue(t *testing.T) {
    buffer := []int{1,2,2,4}
    cp := make([]int, len(buffer))
    copy(cp, buffer)

    cp = RemoveAndAdd(cp, 1, 5)

    if !reflect.DeepEqual([]int{2,2,4,5}, cp) {
	t.Error("Incorrect result.")
    }

    cp = RemoveAndAdd(cp, 2, 6)

    if !reflect.DeepEqual([]int{2,4,5,6}, cp) {
	t.Error("Incorrect result.")
    }
    
    cp = RemoveAndAdd(cp, 6, 8)

    if !reflect.DeepEqual([]int{2,4,5, 8}, cp) {
	t.Error("Incorrect result.")
    }
}
