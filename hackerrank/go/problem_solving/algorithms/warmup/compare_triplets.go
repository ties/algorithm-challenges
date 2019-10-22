package main
/**
Alice and Bob each created one problem for HackerRank. A reviewer rates the two
challenges, awarding points on a scale from $1$ to $100$ for three categories:
problem clarity, originality, and difficulty.

We define the rating for Alice's challenge to be the triplet
$a = (a[0], a[1], a[2])$ and the rating for Bob's challenge to be the triplet 
$b = (b[0], b[1], b[2])$.

Your task is to find their comparison points by comparing  $a[0]$ with $b[0]$,
$a[1]$ with $b[1]$, and $a[2]$ with $b[2]$.

If $a[i] > b[i]$ then Alice is awarded $1$ point.
If $a[i] < b[i]$, then Bob is awarded $1$ point.
If $a[i] = b[i]$, then neither person receives a point.

Comparison points is the total points a person earned.

Given $a$ and $b$, determine their respective comparison points.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
    res := []int32{0,0}

    for idx := range a {
        if (a[idx] < b[idx]) {
            res[1] += 1
        } else if (a[idx] > b[idx]) {
            res[0] += 1
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

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < 3; i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var b []int32

    for i := 0; i < 3; i++ {
        bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
        checkError(err)
        bItem := int32(bItemTemp)
        b = append(b, bItem)
    }

    result := compareTriplets(a, b)

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
