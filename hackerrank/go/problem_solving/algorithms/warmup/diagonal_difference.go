package main
/**
Given a square matrix, calculate the absolute difference between the sums of its
diagonals.

For example, the square matrix  is shown below:
```
1 2 3
4 5 6
9 8 9  
```
The left-to-right diagonal is $1+5+9=15$. The right to left diagonal is
$3+5+9=17$. Their absolute difference is $|15-17|=2$.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
    // 1: Calculate the two diagonals
    var left, right int32 = 0, 0

    // Assume square matrix
    arrDim := len(arr)
    if (arrDim != len(arr[0])) {
        panic("Not a square matrix")
    }

    for idx := range arr {
        left += arr[idx][idx]
        right += arr[idx][arrDim - idx - 1]
    }

    // 2: absolute difference
    diff := left-right
    if diff < 0 {
        return -diff
    }
    return diff

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var arr [][]int32
    for i := 0; i < int(n); i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(n) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := diagonalDifference(arr)

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
