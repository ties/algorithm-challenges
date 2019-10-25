package main
/**
You will be given an array of integers and a target value. Determine the number
of pairs of array elements that have a difference equal to a target value.

For example, given an array of [1, 2, 3, 4] and a target value of 1, we have
three values meeting the condition: $2-1=1$, $3-2=1$, and $4-3=1$.

Function Description
--------------------

Complete the pairs function below. It must return an integer representing the
number of element pairs having the required difference.

pairs has the following parameter(s):
  k: an integer, the target difference
  arr: an array of integers

Input Format
------------

The first line contains two space-separated integers $n$ and $k$, the size of
$arr$ and the target value.
The second line contains $n$ space-separated integers of the array .
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

type intarr []int32

/**
 * Sort protocol for int32
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

// Complete the pairs function below.
func pairs(k int32, arr []int32) int32 {
  // each integer add[i] will be unique
  // sort them so that it is invariant that when the higher
  // value is encountered the lower value has been seen.
  sort.Sort(intarr(arr))
  seen := make(map[int32]bool)

  res := int32(0)

  for _, val := range(arr) {
    if seen[val] {
      res += 1
    }
    seen[val+k] = true
  }

  return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := pairs(k, arr)

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
