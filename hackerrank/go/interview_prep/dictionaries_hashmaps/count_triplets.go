package main
/**
You are given an array and you need to find number of tripets of indices
$(i, j, k)$ such that the elements at those indices are in geometric progression
for a given common ratio $r$ and $i < j < k$.

For example, $arr=[1,4,16,64]$. If $r=4$, we have $[1,4,16]$ and $[4,16,64]$ at
indices $(0,1,2)$ and $(1,2,3)$.

Function Description
--------------------

Complete the countTriplets function in the editor below. It should return the
number of triplets forming a geometric progression for a given $r$ as an
integer.

countTriplets has the following parameter(s):

    arr: an array of integers
    r: an integer, the common ratio
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
    // Number of ways this number can be created by single mult
    pair := make(map[int64]int64)
    // By second multiple
    triple := make(map[int64]int64)
    
    triplets := int64(0)

    // Invariant: i < j < k (single iteration)
    for _, val := range(arr) {
        triplets += triple[val]

        if pair[val] > 0 {
            // val can be the result of a pair in pair[val] ways
            triple[r*val] += pair[val]
        }

        pair[r*val] += 1
    }

    return triplets
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nr[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    r, err := strconv.ParseInt(nr[1], 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int64

    for i := 0; i < int(n); i++ {
        arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arr = append(arr, arrItem)
    }

    ans := countTriplets(arr, r)

    fmt.Fprintf(writer, "%d\n", ans)

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
