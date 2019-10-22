package main
/**
Given five positive integers, find the minimum and maximum values that can be
calculated by summing exactly four of the five integers. Then print the
respective minimum and maximum values as a single line of two space-separated
long integers.

For example, $arr=[1,3,5,7,9]. Our minimum sum is $1=3+5+7=16$ and our maximum
sum is $3+5+7+9=24$. We would print:
```
16 24
```

**Sane solution**:
Iterate over list, find minimum and maximum value, highest sum is
$total - min$, minimum sum is $total - max$.

But preferred to re-implement the unrank function of the k-subset algorithm.
*/
import (
    "bufio"
    "fmt"
    "io"
    "math"
    "math/big"
    "os"
    "strconv"
    "strings"
)

/**
Unrank: Return subset with rank $r$ of size $k$ from $arr$

Implementation based on my earlier Java implementation at [0].
First time using `math/big`.

[0]: https://github.com/ties/commoncrawl-jsusage-v2/blob/master/collections/src/main/java/nl/tdk/collections/KSubsetLexRank.java
*/
func unrank(r *big.Int, k int64, arr []int32) []int32 {
    // fmt.Printf("unrank(%v, %v, %v)\n", r, k, arr)
    res := []int32{}

    y := new(big.Int)

    x := int64(1)
    n := int64(len(arr))

    for i:=int64(1); i <=k; i++ {
        y := y.Binomial(n-x, k-i)
        for y.Cmp(r) < 0 {
            r = r.Sub(r, y)
            x++
            y = y.Binomial(n-x, k-i)
        }

        // 0-based vs 1-based
        // hidden in `objects(...)` in Java implementation.
        res = append(res, arr[x-1])
        x++
    }

    // fmt.Printf("%v", res)

    return res
}

func sumUnrank(r int64, k int64, arr []int32) int64 {
    var res int64 = 0

    for _, val := range unrank(big.NewInt(r), k, arr) {
        res += int64(val)
    }

    return res
}

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
    minSum := int64(math.MaxInt64)
    maxSum := int64(math.MinInt64)

    k := int64(len(arr)-1)

    setCount := new(big.Int).Binomial(int64(len(arr)),k)
    // fmt.Printf("setCount %v\n", setCount)

    for i:=int64(1); setCount.Cmp(big.NewInt(i)) >= 0; i++ {
        // fmt.Printf("i = %d\n", i)
        sum := sumUnrank(i, k, arr)

        if sum < minSum {
            minSum = sum
        }
        if sum > maxSum {
            maxSum = sum
        }
    }

    fmt.Printf("%v %v\n", minSum, maxSum)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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
