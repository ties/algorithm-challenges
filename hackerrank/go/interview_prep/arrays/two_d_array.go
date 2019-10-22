package main

import (
    "bufio"
    "fmt"
    "math"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the hourglassSum function below.
// We define an hourglass in  to be a subset of values with indices falling in
// this pattern in 's graphical representation:
// 
// a b c
//   d
// e f g
// There are  hourglasses in , and an hourglass sum is the sum of an hourglass'
// values. Calculate the hourglass sum for every hourglass in , then print the
// maximum hourglass sum.
func hourglassSum(arr [][]int32) int32 {
    var maxSum int32 = math.MinInt32

    for i := 0; i <= len(arr) - 3; i++ {
        for j := 0; j <= len(arr[0]) - 3; j++ {
            // Get the hour glass
            var total int32 = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

            // fmt.Printf("%v\n", total)
            if total > maxSum {
                maxSum = total
            }
        }
    }
    return maxSum
}
/*
-9 -9 -9 1 1 1
0 -9 0 4 3 2
-9 -9 -9 1 2 3
0 0 8 6 6 0
0 0 0 -2 0 0
0 0 1 2 4 0
*/
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

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
