package main
/**
In an array, $arr$, the elements at indices $i$ and $j$ (where $i<j$) form an
inversion if $arr[i] > arr[j]$. In other words, inverted elements  and  are
considered to be "out of order". To correct an inversion, we can swap adjacent
elements.

For example, consider the dataset $arr=[2,4,1]$. It has two inversions: $(4,1)$
and $(2,1)$. To sort the array, we must perform the following two swaps to 
correct the inversions:

Given $d$ datasets, print the number of inversions that must be swapped to sort
each dataset on a new line.

Function Description
--------------------
Complete the function countInversions in the editor below. It must return an
integer representing the number of inversions required to sort the array.

countInversions has the following parameter(s):
	arr: an array of integers to sort.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func Merge(a []int32, iLeft int, iRight int, iEnd int, b []int32) int64 {
    swaps := int64(0)

    i, j := iLeft, iRight

    for k:=iLeft; k<iEnd;k++ {
        if (i < iRight && (j >= iEnd || a[i] <= a[j])) {
            b[k] = a[i]
            i += 1
        } else {
            b[k] =a[j]
            j += 1
        }
    }

    return swaps
}

func MergeSortInner(a []int32, b []int32, n int) int64 {
    swaps := int64(0)

    // Single element runs are already sorted.
    // Make successively longer runs (2, 4, ...)
    for width := 1; width < n; width *= 2 {
        //  A is full of runs of length width
        for i:=0; i < n; i=i + 2*width {
            // Merge two runs: A[i:i+width-1] and A[i+width:i+2*width-1] to B[]
            // or copy A[i:n-1] to B[] ( if(i+width >= n) )
            mw := n
            if i+width < n {
                mw = i+width
            }
            m2w := n
            if i+2*width < n {
                m2w = i+2*width
            }
            swaps += Merge(a, i, mw, m2w, b)
        }

        copy(a, b)
    }


    return swaps
}


/**
 * Sort the given array using merge sort,
 * Return the number of inversions
 */
func MergeSort(arr []int32) int64 {
    n := len(arr)
    // Make a buffer array of same size.
    b := make([]int32, n)

    return MergeSortInner(arr, b, n)
}

// Complete the countInversions function below.
func countInversions(arr []int32) int64 {
    return MergeSort(arr)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        arrTemp := strings.Split(readLine(reader), " ")

        var arr []int32

        for i := 0; i < int(n); i++ {
            arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arr = append(arr, arrItem)
        }

        result := countInversions(arr)

        fmt.Fprintf(writer, "%d\n", result)
    }

    writer.Flush()
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
