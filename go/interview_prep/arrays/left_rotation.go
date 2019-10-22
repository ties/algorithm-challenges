package main
/**
A left rotation operation on an array shifts each of the array's elements  unit
to the left. For example, if 1 left rotations are performed on array [1,2,3,4],
then the array would become: [2,3,4,1]

Given an array $a$ of $n$ integers and a number, $d$, perform $d$ left rotations
on the array. Return the updated array to be printed as a single line of
space-separated integers.
*/

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
    aLen := len(a)
    // Calculate number of complete rotations
    leftRotations := int(d)%aLen
    // Ugly O(N) memory solution

    // Create temporary slice
    tmp := make([]int32, aLen)
    copy(tmp, a)

    // fmt.Printf("%v\n", tmp)
    // fmt.Printf("====\n")

    for i:=0; i < aLen; i++ {
        // Calculate two positions:
        // a[i] and a[i-d]
        src := i + leftRotations
        if src >= aLen {
            src = src % aLen
        }

        // fmt.Printf("%d<-%d %v %v\n", i, src, a[i], tmp[src])
        a[i] = tmp[src]
    }

    return a
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nd := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nd[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    dTemp, err := strconv.ParseInt(nd[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    aTemp := strings.Split(readLine(reader), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    result := rotLeft(a, d)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
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
