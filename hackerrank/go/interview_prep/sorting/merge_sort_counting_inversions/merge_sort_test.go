package main

import (
    "fmt"
    "testing"
    "reflect"
)

// func TestSortSingleElement(t *testing.T) {
//     buffer := []int32{42}

//     swaps := MergeSort(buffer)

//     if swaps != int64(0) {
//         t.Error(fmt.Sprintf("%d swaps != %d", swaps, 0))
//     }

//     if !reflect.DeepEqual(buffer, []int32{42}) {
//         t.Error(fmt.Sprintf("Wrong result"))
//     }
// }

// func TestSortPair(t *testing.T) {
//     buffer := []int32{13, 7}

//     swaps := MergeSort(buffer)

//     if swaps != int64(1) {
//         t.Error(fmt.Sprintf("%d swaps != %d", swaps, 1))
//     }

//     if !reflect.DeepEqual(buffer, []int32{7,13}) {
//         t.Error(fmt.Sprintf("Wrong result"))
//     }
    
//     buffer = []int32{7, 13}

//     swaps = MergeSort(buffer)

//     if swaps != int64(0) {
//         t.Error(fmt.Sprintf("%d swap for sorted array", swaps))
//     }
//     if !reflect.DeepEqual(buffer, []int32{7,13}) {
//         t.Error(fmt.Sprintf("Wrong result"))
//     }
// }

// func TestSortMultiple(t *testing.T) {
//     buffer := []int32{3,2,1}

//     swaps := MergeSort(buffer)

//     if swaps != int64(2) {
//         t.Error(fmt.Sprintf("%d swaps != %d", swaps, 2))
//     }

//     if !reflect.DeepEqual(buffer, []int32{1,2,3}) {
//         t.Error(fmt.Sprintf("Wrong result"))
//     }
// }

// func TestSortStable(t *testing.T) {
//     buffer := []int32{2,2,1}

//     swaps := MergeSort(buffer)

//     if swaps != int64(1) {
//         t.Error(fmt.Sprintf("%d swaps != %d", swaps, 1))
//     }

//     if !reflect.DeepEqual(buffer, []int32{1,2,2}) {
//         t.Error(fmt.Sprintf("Wrong result"))
//     }
// }

func TestSortExample(t *testing.T) {
    buffer := []int32{2,4,1}

    swaps := MergeSort(buffer)

    if swaps != int64(2) {
        t.Error(fmt.Sprintf("Not two swaps: %v", swaps))
    }

    if !reflect.DeepEqual(buffer, []int32{1,2,4}) {
        t.Error(fmt.Sprintf("Wrong result"))
    }

    swaps = MergeSort([]int32{2,1,3,1,2})

    if swaps != int64(4) {
        t.Error(fmt.Sprintf("Not 4 swaps: %v", swaps))
    }

    swaps = MergeSort([]int32{7,5,3,1})

    if swaps != int64(6) {
        t.Error(fmt.Sprintf("Not 6 swaps: %v.", swaps))
    }
}
