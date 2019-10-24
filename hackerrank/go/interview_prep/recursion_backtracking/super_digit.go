package main
/**
We define super digit of an integer $x$ using the following rules:

Given an integer, we need to find the super digit of the integer.
  * If $x$ has only $1$ digit, then its super digit is $x$.
  * Otherwise, the super digit of $x$ is equal to the super digit of the sum of
    the digits of $x$.

For example, the super digit of $9875$ will be calculated as:
```
	super_digit(9875)   	9+8+7+5 = 29 
	super_digit(29) 	2 + 9 = 11
	super_digit(11)		1 + 1 = 2
	super_digit(2)		= 2 
```

You are given two numbers $n$ and $k$. The number $p$ is created by
concatenating the string $n$ $k$ times. Continuing the above example where
$n=9875$, assume your value $k=4$. Your initial
$p = 9875 9875 9875 9875$ (spaces added for clarity).

```
    superDigit(p) = superDigit(9875987598759875)
                  5+7+8+9+5+7+8+9+5+7+8+9+5+7+8+9 = 116
    superDigit(p) = superDigit(116)
                  1+1+6 = 8
    superDigit(p) = superDigit(8)
```
All of the digits of  $p$ sum to $116$. The digits of $116$ sum to $8$. $8$ is
only one digit, so it's the super digit.

Function Description
--------------------

Complete the function superDigit in the editor below. It must return the
calculated super digit as an integer.

superDigit has the following parameter(s):
    n: a string representation of an integer
    k: an integer, the times to concatenate  to make 
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the superDigit function below.
func superDigit(n string, k int32) int32 {
    // First: superDigit just for n
    // read digit-wise as number
    digitSum := int64(0)

    for _, val := range(n) {
        d, _ := strconv.ParseInt(string(val), 10, 32)
        digitSum += int64(d)
    }

    // Two cases:
    if digitSum > 10 || k > 1 {
        return superDigit(fmt.Sprintf("%d", int64(k)*digitSum), 1)
    }
    // < 10 -> fits in int32
    return int32(digitSum)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    n := nk[0]

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := superDigit(n, k)

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
