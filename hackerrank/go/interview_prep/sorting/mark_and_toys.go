package main
/**
Mark and Jane are very happy after having their first child. Their son loves
toys, so Mark wants to buy some. There are a number of different toys lying in
front of him, tagged with their prices. Mark has only a certain amount to spend,
and he wants to maximize the number of toys he buys with this money.

Given a list of prices and an amount to spend, what is the maximum number of
toys Mark can buy? For example, if $prices=[1,2,3,4]$ and Mark has $k=7$ to
spend, he can buy items $[1,2,3]$ for $6$ or $[3,4]$ for $7$ units of currency.
He would choose the first group of  items.

Function Description
--------------------

Complete the function maximumToys in the editor below. It should return an
integer representing the maximum number of toys Mark can purchase.

maximumToys has the following parameter(s):
	prices: an array of integers representing toy prices
	k: an integer, Mark's budget
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


// Complete the maximumToys function below.
func maximumToys(pr []int32, k int32) int32 {
	// Note that:
	// 1 < n < 10**5:   a N^2 memory solution (dynamic programming) would
	// 		    not fit.
	// 1 <= k <= 10**9: An array per amount would not fit either.
	//
	// And that being greedy works: A toy priced 1 is one toy, so is one
	// at any other price.
	//
	prices := intarr(pr)
	// Sort the array for the greedy solution.
	sort.Sort(prices)

	for idx, val := range(prices) {
		if val <= k {
			k -= val
		} else {
			// At the first item you can not buy,
			// you have bought (idx-1)+1 items
			// idx-1: last item bought, +1: 0 vs 1-indexed.
			return int32(idx)
		}
	}
	return 0
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

    pricesTemp := strings.Split(readLine(reader), " ")

    var prices []int32

    for i := 0; i < int(n); i++ {
        pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
        checkError(err)
        pricesItem := int32(pricesItemTemp)
        prices = append(prices, pricesItem)
    }

    result := maximumToys(prices, k)

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
