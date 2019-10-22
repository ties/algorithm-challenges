package main
/**
Given an array of integers, calculate the fractions of its elements that are
positive, negative, and are zeros. Print the decimal value of each fraction on
a new line.

Note: This challenge introduces precision problems. The test cases are scaled
to six decimal places, though answers with absolute error of up to  are
acceptable.

For example, given the array $arr = [1,1,0,-1,-1] there are $5$ elements, two
positive, two negative and one zero. Their ratios would be $0.4000000$,
$0.400000$ and $0.2000000$. It should be printed as
```
0.400000
0.400000
0.200000
```
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
    var pos, neg, zero int = 0, 0, 0
    // Create a float for float division.
    count := float64(len(arr))

    for _, val := range arr {
        if val < 0 {
            neg++
        } else if val > 0 {
            pos++
        } else {
            zero++
        }
    }

    fmt.Printf("%.6f\n", float64(pos)/count)
    fmt.Printf("%.6f\n", float64(neg)/count)
    fmt.Printf("%.6f\n", float64(zero)/count)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

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

    plusMinus(arr)
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
