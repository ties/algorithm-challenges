package main
/**
Given 3 arrays $a$, $b$, $c$ of different sizes, find the number of distinct
triplets $(p,q,r)$ where $p$ in $a$, $q$ in $b$, $r$ in $c$, satisfying the
criteria: $p <= q$ and $q >= r$

For example, given $a=[3,5,7]$, $b=[3,6]$ and $c=[4,6,9]$, we find four distinct
triplets: $(3,6,4)$,$(3,6,6)$,$(5,6,6)$.

Function Description
--------------------

Complete the triplets function in the editor below. It must return the number of
distinct triplets that can be formed from the given arrays.

triplets has the following parameter(s):
  a, b, c: three arrays of integers.
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

// Complete the triplets function below.
func triplets(a []int32, b []int32, c []int32) int64 {
  // Sort a, b, and c
  sort.Sort(intarr(a))
  sort.Sort(intarr(b))
  sort.Sort(intarr(c))

  fmt.Println("==================================")
  fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)
  fmt.Println("==================================")
  res := int64(0)
  // for every q in b:
  // add |p|*1*|q| elements to result
  pIdx := 0
  rIdx := 0

  numP := 0
  numR := 0

  lastQ := int32(-1)

  for _, qVal := range(b) {
    // Unique q value
    if qVal == lastQ {
      continue
    }
    lastQ = qVal

    // iterate forward in a until a[p_idx] > qVal
    for ;pIdx < len(a) && a[pIdx] <= qVal; pIdx++ {
      if pIdx == 0 || a[pIdx] != a[pIdx-1] {
        numP += 1
      }
    }
    // iterate forward in c until c[q_idx] > qVal
    for ;rIdx < len(c) && c[rIdx] <= qVal; rIdx++ {
      if rIdx == 0 || c[rIdx] != c[rIdx-1] {
        numR += 1
      }
    }

    fmt.Printf("q = %v, #p=%v pIdx=%v, #r=%v rIdx=%v\n", qVal, numP, pIdx, numR, rIdx)
    fmt.Printf("adding %v\n", int64(numP)*int64(numR))
    res += int64(numP)*int64(numR)
  }

  fmt.Println("==================================")

  return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    lenaLenbLenc := strings.Split(readLine(reader), " ")

    lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
    checkError(err)
    lena := int32(lenaTemp)

    lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
    checkError(err)
    lenb := int32(lenbTemp)

    lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
    checkError(err)
    lenc := int32(lencTemp)

    arraTemp := strings.Split(readLine(reader), " ")

    var arra []int32

    for i := 0; i < int(lena); i++ {
        arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
        checkError(err)
        arraItem := int32(arraItemTemp)
        arra = append(arra, arraItem)
    }

    arrbTemp := strings.Split(readLine(reader), " ")

    var arrb []int32

    for i := 0; i < int(lenb); i++ {
        arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
        checkError(err)
        arrbItem := int32(arrbItemTemp)
        arrb = append(arrb, arrbItem)
    }

    arrcTemp := strings.Split(readLine(reader), " ")

    var arrc []int32

    for i := 0; i < int(lenc); i++ {
        arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
        checkError(err)
        arrcItem := int32(arrcItemTemp)
        arrc = append(arrc, arrcItem)
    }

    ans := triplets(arra, arrb, arrc)

    fmt.Fprintf(writer, "%d\n", ans)

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
