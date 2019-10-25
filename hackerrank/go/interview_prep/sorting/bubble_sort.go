package main

/**
Consider the following version of Bubble Sort:
```
for (int i = 0; i < n; i++) {

    for (int j = 0; j < n - 1; j++) {
        // Swap adjacent elements if they are in decreasing order
        if (a[j] > a[j + 1]) {
            swap(a[j], a[j + 1]);
        }
    }
}
```

Given an array of integers, sort the array in ascending order using the Bubble
Sort algorithm above. Once sorted, print the following three lines:

  1. Array is sorted in numSwaps swaps., where $numSwaps$ is the number of swaps
     that took place.
  2. First Element: firstElement, where $firstElement$ is the first element in
     the sorted array.
  3. Last Element: lastElement, where $lastElement$ is the last element in the
     sorted array.

Hint: To complete this challenge, you must add a variable that keeps a running
tally of all swaps that occur during execution.

For example, given a worst-case but small array to sort: $a=[6,4,1]$ we go
through the following steps:
```
swap    a
0       [6,4,1]
1       [4,6,1]
2       [4,1,6]
3       [1,4,6]
```
It took 3 swaps to sort the array. Output would be
```
Array is sorted in 3 swaps.
First Element: 1
Last Element: 6
```

Function Description
--------------------

Complete the function countSwaps in the editor below. It should print the three
lines required, then return.

countSwaps has the following parameter(s):
  a: an array of integers.
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type intarr []int32

/**
 * Sort protocol for int32
 */
func (f intarr) Len() int {
	return len(f)
}

func (f intarr) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f intarr) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Complete the countSwaps function below.
func countSwaps(a []int32) {
	arr := intarr(a)

	swaps := 0

	// Bubble sort
	// https://en.wikipedia.org/wiki/Bubble_sort#Implementation
	swapped := true
	n := len(arr)

	for swapped {
		swapped = false
		for i:=1; i<=n-1; i++ {
			if arr.Less(i, i-1) {
				arr.Swap(i, i-1)
				swapped = true
				swaps += 1
			}
		}
		n -= 1
	}

	// Three outputs:
	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[len(arr)-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
