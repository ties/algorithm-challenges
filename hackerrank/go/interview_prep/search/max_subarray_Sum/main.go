package main
/**
We define the following:
  * A subarray of array $a$ of length $n$ is a contiguous segment from $a[i]$
    through $a[j]$ where $0 <= i <= j <= n$.
  * The sum of an array is the sum of its elements.

Given an $n$ element array of integers, $a$, and an integer, $m$, determine the
maximum value of the sum of any of its subarrays modulo $m$. For example, Assume
$a=[1,2,3]$ and $m=2$. The following table lists all subarrays and their moduli:
```
		sum	%2
[1]		1	1
[2]		2	0
[3]		3	1
[1,2]		3	1
[2,3]		5	1
[1,2,3]		6	0
```
The maximum modulus is 1.

Function Description
--------------------
Complete the maximumSum function in the editor below. It should return a long
integer that represents the maximum value of $subarray sum % m$.

maximumSum has the following parameter(s):
  * a: an array of long integers, the array to analyze
  * m: a long integer, the modulo divisor
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

type intarr []int64

/**
 * Sort protocol for int64
 */
func (f intarr) Len() int {
  return len(f)
}

func (f intarr) Less(i, j int) bool {
  return f[i] < f[j]
}

func (f intarr) Swap(i, j int) {
  f[i], f[j] = f[j], f[i]
}

/**
The maximum sum is the maximum sum of taking any
value + whether or nor you take the previous value

recursively.
*/
func MaxSum(a []int64, n int, m int64, res int64) int64 {
  fmt.Printf("MaxSum(n=%4v, res=%4v)\n", n, res)
  // Base case.
  if n == 0 {
    if (res+a[0])%m > res {
      return res+a[0] %m
    }
    return res
  }

  // Otherwise, pick or do not pick.
  with := MaxSum(a, n-1, m, (res + a[n])%m)
  withOut := MaxSum(a, n-1, m, res)

  if with > withOut {
    return with
  }
  return withOut
}

// Complete the maximumSum function below.
// Bounds:
// 2 <= len(a) <= 10**5 => PowerSet(a) is too big
func maximumSum(a []int64, m int64) int64 {
  // First step: Take all a values modulo m
  // and initialize arrays to -1
  for idx, val := range(a) {
    a[idx] = val % m
  }

  // Sort the values
  sort.Sort(intarr(a))

  return MaxSum(a, len(a)-1, m, 0)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        nm := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nm[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        m, err := strconv.ParseInt(nm[1], 10, 64)
        checkError(err)

        aTemp := strings.Split(readLine(reader), " ")

        var a []int64

        for i := 0; i < int(n); i++ {
            aItem, err := strconv.ParseInt(aTemp[i], 10, 64)
            checkError(err)
            a = append(a, aItem)
        }

        result := maximumSum(a, m)

        fmt.Fprintf(writer, "%d\n", result)
    }

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
