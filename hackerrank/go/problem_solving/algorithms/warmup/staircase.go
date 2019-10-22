package main
/**
Consider a staircase of size :
```
   #
  ##
 ###
####
```
Observe that its base and height are both equal to $n$, and the image is drawn
using # symbols and spaces. The last line is not preceded by any spaces.

Write a program that prints a staircase of size $n$.
*/

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the staircase function below.
func staircase(n int) {
    // Each line: $n-i$ spaces, $i$ hashes
    for i:=1; i <= n; i++ {
        fmt.Printf("%s%s\n", strings.Repeat(" ", n-i), strings.Repeat("#", i))
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int(nTemp)

    staircase(n)
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
