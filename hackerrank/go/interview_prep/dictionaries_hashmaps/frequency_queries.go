package main
/**
You are given $q$ queries. Each query is of the form two integers described
below:
- $1 x$: Insert x in your data structure.
- $2 y$: Delete one occurence of y from your data structure, if present.
- $3 z$: Check if any integer is present whose frequency is exactly . If yes,
         print 1 else 0.

The queries are given in the form of a 2-D array $queries$ of size $q$ where
$queries[i][0]$ contains the operation, and $queries[i][1]$ contains the data
element. For example, you are given array
$queries = [(1,1),(2,2),(3,2),(1,1),(1,1),(2,1),(3,2)]$. The results of each
operation are:
```
Operation   Array   Output
(1,1)       [1]
(2,2)       [1]
(3,2)                   0
(1,1)       [1,1]
(1,1)       [1,1,1]
(2,1)       [1,1]
(3,2)                   1
```
Return an array with the output: [0,1].

Function Description
--------------------

Complete the freqQuery function in the editor below. It must return an array of
integers where each element is a $1$ if there is at least one element value with
the queried number of occurrences in the current array, or $0$ if there is not.

freqQuery has the following parameter(s):
    queries: a 2-d array of integers
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
    //
    // Type:
    // 1: add(n)
    // 2: delete(n)
    // 3: count(n)
    //
    res := []int32{}
    state := make(map[int32]int32)

    // How often a value is present
    presentCount := make(map[int32]int32)

    for _, query := range(queries) {
        // Type 1
        cmd := query[0]
        arg := query[1]

        if cmd == 1 {
            if state[arg] > 0 {
                presentCount[state[arg]] -= 1
            }
            state[arg] += 1
            
            presentCount[state[arg]] += 1
        } else if cmd == 2 {
            if state[arg] > 0 {
                presentCount[state[arg]] -= 1
                state[arg] -= 1
            }
            presentCount[state[arg]] += 1
        } else if cmd == 3 {
            // Find query[1] in the state
            if presentCount[arg] > 0 {
                res = append(res, int32(1))
            } else {
                res = append(res, int32(0))
            }
        }
    }


    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    var queries [][]int32
    for i := 0; i < int(q); i++ {
        queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != 2 {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    ans := freqQuery(queries)

    for i, ansItem := range ans {
        fmt.Fprintf(writer, "%d", ansItem)

        if i != len(ans) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
