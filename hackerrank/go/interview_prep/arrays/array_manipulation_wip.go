package main
/**
Starting with a 1-indexed array of zeros and a list of operations, for each
operation add a value to each of the array element between two given indices,
*inclusive*.

Once all operations have been performed, return the maximum value in your array.

For example, the length of your array of zeros $n=10$. Your list of queries is
as follows:
```
    a b k
    1 5 3
    4 8 7
    6 9 1
```
Add the values of $k$ between the indices $a$ and $b$ inclusive:
```
index->	 1 2 3  4  5 6 7 8 9 10
	[0,0,0, 0, 0,0,0,0,0, 0]
	[3,3,3, 3, 3,0,0,0,0, 0]
	[3,3,3,10,10,7,7,7,0, 0]
	[3,3,3,10,10,8,8,8,1, 0]
```

The largest value is $10$ after all operations are performed.

Approaches taken:
  * Naieve array based approach - filling every value in interval: slow
  * (Python) list of integer intervals with queries splitting these intervals:
    too slow.
  * (Python) arrays of [lhs, rhs, val] for integer intervals instead of object
    based, bisecting array to find intervals to evaluate: too slow.

Final approach: Value at x_i is the sum over x_0...x_i.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
    // slice is initialized with zeroes
    // 1-based indices + another element for the decrease
    // if upper bound == n
    arr := make([]int64, n+2)

    for _, row := range queries {
        var lhs, rhs, k = row[0], row[1], row[2]

        arr[lhs] += int64(k)
        arr[rhs+1] -= int64(k)
    }

    // Find maximum by adding up the changes into an accumulator.
    // all k are positive => invariant: acc >= 0.
    var acc, max int64 = 0, 0

    for _, val := range arr {
        acc += val
        if acc > max {
            max = acc
        }
    }

    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    var queries [][]int32
    for i := 0; i < int(m); i++ {
        queriesRowTemp := strings.Split(readLine(reader), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != int(3) {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := arrayManipulation(n, queries)

    fmt.Fprintf(writer, "%d\n", result)

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
