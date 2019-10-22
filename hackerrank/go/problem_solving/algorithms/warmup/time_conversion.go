package main
/**
Given a time in -hour AM/PM format, convert it to military (24-hour) time.

Note: Midnight is 12:00:00AM on a 12-hour clock, and 00:00:00 on a 24-hour
clock. Noon is 12:00:00PM on a 12-hour clock, and 12:00:00 on a 24-hour clock.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "time"
    "strings"
)

/*
 * Complete the timeConversion function below.
 *
 * reference: https://gobyexample.com/time-formatting-parsing
 */
func timeConversion(s string) string {
    // Parse time
    t1, _ := time.Parse("3:04:05PM", s)
    return t1.Format("15:04:05")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

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
