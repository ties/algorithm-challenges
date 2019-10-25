package main
/**
Each time Sunny and Johnny take a trip to the Ice Cream Parlor, they pool their
money to buy ice cream. On any given day, the parlor offers a line of flavors.
Each flavor has a cost associated with it.

Given the value of $money$ and the $cost$ of each flavor for $t$ trips to the
Ice Cream Parlor, help Sunny and Johnny choose two distinct flavors such that
they spend their entire pool of money during each visit. ID numbers are the 1-
based index number associated with a $cost$. For each trip to the parlor, print
the ID numbers for the two types of ice cream that Sunny and Johnny purchase as
two space-separated integers on a new line. You must print the smaller ID first
and the larger ID second.

For example, there are $n=5$ flavors having $cost=[2,1,3,5,6]$. Together they
have $money=5$ to spend. They would purchase flavor ID's $1$ and $3$ for a cost
of $2+3+5$. Use  based indexing for your response.

Note:
  * Two ice creams having unique IDs $i$ and $j$ may have the same cost (i.e.,
    $cost[i] == cost[j]$).
  * There will always be a unique solution.

Function Description
--------------------
Complete the function whatFlavors in the editor below. It must determine the two
flavors they will purchase and print them as two space-separated integers on a
line.

whatFlavors has the following parameter(s):
    cost: an array of integers representing price for a flavor
    money: an integer representing the amount of money they have to spend
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) {
  //
  // Problem comes down to: find two integers in cost that sum
  // to money
  //

  // map (remainder) -> index of first choice
  remainder := make(map[int32]int)

  for idx, val := range(cost) {
	  if remainder[val] > 0 {
		  fmt.Printf("%d %d\n", remainder[val], idx+1)
	  }
	  remainder[money-val] = idx+1
  }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        money := int32(moneyTemp)

        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        costTemp := strings.Split(readLine(reader), " ")

        var cost []int32

        for i := 0; i < int(n); i++ {
            costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
            checkError(err)
            costItem := int32(costItemTemp)
            cost = append(cost, costItem)
        }

        whatFlavors(cost, money)
    }
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
