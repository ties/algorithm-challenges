package main
/**
You are in charge of the cake for your niece's birthday and have decided the
cake will have one candle for each year of her total age. When she blows out
the candles, she’ll only be able to blow out the tallest ones. Your task is to
find out how many candles she can successfully blow out.

For example, if your niece is turning $4$ years old, and the cake will have $4$
candles of height $4$, $4$, $1$, $3$ she will be able to blow out $2$ candles
successfully, since the tallest candles are of height $4$ and there are $2$ such
candles.
*/
import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
    // Two scans
    // find max
    var max int32 = math.MinInt32
    for _, val := range ar {
        if val > max {
            max = val
        }
    }

    var count int32 = 0
    // match
    for _, val := range ar {
        if val == max {
            count++
        }
    }
    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(arCount); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := birthdayCakeCandles(ar)

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
