package main
/**
You are given an unordered array consisting of consecutive integers 
$[1, 2, 3, ..., n]$ without any duplicates. You are allowed to swap any two
elements. You need to find the minimum number of swaps required to sort the
array in ascending order.

For example, given the array $arr = [7,1,3,2,4,5,6]$ we perform the following
steps:

i   arr                     swap (indices)
0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
5   [1, 2, 3, 4, 5, 6, 7]
It took $5$ swaps to sort the array.

Approach: A* with distance from sorted as metrinc.

Take a priority queue, sorted by the edit distance and return
when the result is definitely (edit distance/2 + steps taken > cur).
Add all new elements to the priorityqueue [0].

[0]: https://golang.org/pkg/container/heap/
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


func swap(arr []int32, lookup []int32, lhs int32, rhs int32) {
    /** Directly swap two items. */
    // First update lookup
    lookup[arr[rhs]] = lhs
    lookup[arr[lhs]] = rhs

    // Afterwards, swap
    tmp := arr[lhs]
    arr[lhs] = arr[rhs]
    arr[rhs] = tmp
}



// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
    numItems := len(arr)
    edits := int32(0)

    // Lookup table from value -> array index
    lookup := make([]int32, numItems+1)
    for idx, val := range(arr) {
        lookup[val] = int32(idx)
    }

    // Do swaps. Note that:
    // C A B D
    //   C<->B, BACD A<->B -> ABCD
    //   C<->A, ACBD C<->B -> ABCD
    // A direct swap for a pair followed by a normal swap commutes.
    for idx, _ := range(arr) {
        if arr[idx] != int32(idx)+1 {
            edits += 1
            // Swap, using the target _value at the index_ instead of the
            // index itself.
            swap(arr, lookup, int32(idx), lookup[idx+1])
        }
    }

    return edits
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

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

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

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
