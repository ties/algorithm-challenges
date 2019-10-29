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

// int64 sort protocol
type ArrayElem struct {
  val int64
  idx int
}

func (a ArrayElem) String() string {
  return fmt.Sprintf("{idx: %v, val: %v}", a.idx, a.val)
}

type ArrayElemArr []ArrayElem

/**
 * Sort protocol for int32
 */
func (f ArrayElemArr) Len() int {
  return len(f)
}

/**
Sort by value and index.s
*/
func (f ArrayElemArr) Less(i, j int) bool {
  if f[i].val != f[j].val {
    return f[i].val < f[j].val
  }
  return f[i].idx < f[j].idx
}

func (f ArrayElemArr) Swap(i, j int) {
  f[i], f[j] = f[j], f[i]
}


func max(lhs, rhs int64) int64 {
  if lhs < rhs {
    return rhs
  }
  return lhs
}

// Build array of PrefixSums modulo m
// param a: array of numbers
func PrefixSums(a []int64, m int64) []int64 {
  prefixes := make([]int64, len(a)+1)

  prefixes[0] = 0
  for i:=1; i<len(prefixes); i++ {
    prefixes[i] = (a[i-1]%m + prefixes[i-1]%m + m)%m

    if prefixes[i] < 0 {
      panic("Overflow.")
    }
  }

  return prefixes
}

func ModSub(lhs, rhs int64, m int64) int64 {
  return (lhs - rhs + m)%m
}

// Complete the maximumSum function below.
// Bounds:
// 2 <= len(a) <= 10**5 => PowerSet(a) is too big
// max(a) == int64 -> prefix sums need to be big
//
// **contiguous segment** of a.
func maximumSum(a []int64, m int64) int64 {
  // Calculate prefix sums
  // prefixes[i] = sum up to and including i-th *element*
  // of array (not: index)
  prefixes := PrefixSums(a, m)

  // Create sorted prefix sum array
  prefixSums := make([]ArrayElem, len(prefixes))

  for idx, val := range(prefixes) {
    prefixSums[idx] = ArrayElem{val: val, idx:idx}
  }

  sort.Sort(ArrayElemArr(prefixSums))

  best := int64(-1)

  for _, val := range(prefixSums) {
    remainder := ModSub(m, val.val, m)
    best = max(val.val, best)

    // Search for first element that would wrap around again
    idx := sort.Search(len(prefixSums), func(i int) bool { return prefixSums[i].val >= remainder })

    // First element larger than remainder. Scan backwards for:
    // * index being > cur index
    // * val being < remainder
    if idx >= len(prefixSums) {
      idx = len(prefixSums)-1
    }

    for i:=idx; i>=0; i-- {
      elem := prefixSums[i]

      // Invariant: first elem smaller
      if elem.idx > val.idx && elem.val < remainder {
        best = max(ModSub(elem.val, val.val, m), best)
      }
    }

  }

  return best
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
