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
  "math"
  "os"
  "sort"
  "strconv"
  "strings"
)

// int64 sort protocol
type ArrayElem struct {
  val int64
  idx int
  orig int64
}

func (a ArrayElem) String() string {
  return fmt.Sprintf("{idx: %v, val: %v, orig: %v}", a.idx, a.val, a.orig)
}

type ArrayElemArr []ArrayElem

/**
 * Sort protocol for int32
 */
func (f ArrayElemArr) Len() int {
  return len(f)
}

/**
Sort by value and index.
  * increasing
  * lowest index first
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
  prefixes := make([]int64, len(a))

  for i:=0; i<len(a); i++ {
    if i == 0 {
      prefixes[i] = a[i]%m
    } else {
      prefixes[i] = (a[i]%m + prefixes[i-1]%m)%m
    } 

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

  best := int64(-1)

  // Create tuples to be sorted,
  // picking up highest value as we iterate.
  for idx, val := range(prefixes) {
    prefixSums[idx] = ArrayElem{val: val, idx:idx, orig: a[idx]}

    if best < val {
      best = val
    }
  }

  // Sort in increasing order by (val, idx)
  sort.Sort(ArrayElemArr(prefixSums))

  min := int64(math.MaxInt64)
  //
  // The highest result has is m-the lowest value.
  // The lowest value is the smallest interval between two consecutive values.
  //
  for i:=0; i<len(prefixSums)-1;i++ {
    lhs := prefixSums[i]
    rhs := prefixSums[i+1]

    if lhs.idx > rhs.idx {
      if rhs.val - lhs.val < min {
        min = rhs.val - lhs.val
      }
    }
  }

  return max(m-min, best)
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

/**
 * Change readLine to always return complete line.
 */
func readLine(reader *bufio.Reader) string {
  str := []byte{}

  for true {
    partial, isPrefix, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    str = append(str, partial...)

    if !isPrefix {
      break
    }
  }

  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
