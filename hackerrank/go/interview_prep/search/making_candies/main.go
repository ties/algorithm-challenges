package main
/**
Karl loves playing games on social networking sites. His current favorite is
CandyMaker, where the goal is to make candies.

Karl just started a level in which he must accumulate  candies starting with $m$
machines and $w$ workers. In a single pass, he can make $m\timesw$ candies.
After each pass, he can decide whether to spend some of his candies to buy more
machines or hire more workers. Buying a machine or hiring a worker costs $p$
units, and there is no limit to the number of machines he can own or workers he
can employ.

Karl wants to minimize the number of passes to obtain the required number of
candies at the end of a day. Determine that number of passes.

For example, Karl starts with $m=1$ machine and $w=2$ workers. The cost to
purchase or hire, $p=1$ and he needs to accumulate $60$ candies. He executes the
following strategy:

  1. Make $m \times w = 1 \times 2 = 2$ candies. Purchase two machines.
  2. Make $3 \times 2 = 6$ candies. Purchase $3$ machines and hire $3$ workers.
  3. Make $6 \times 5 = 30$ candies. Retain all $30$ candies.
  4. Make $6 \times 5 = 30$ candies. With yesterday's production, Karl has 
     $60$ candies.
It took $4$ passes to make enough candies.

Function Description
--------------------
Complete the minimumPasses function in the editor below. The function must
return a long integer representing the minimum number of passes required.

minimumPasses has the following parameter(s):
  m: long integer, the starting number of machines
  w: long integer, the starting number of workers
  p: long integer, the cost of a new hire or a new machine
  n: long integer, the number of candies to produce
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Naieve recursive version:
//
//  c: candies in stock
//  m: long integer, the starting number of machines
//  w: long integer, the starting number of workers
//  p: long integer, the cost of a new hire or a new machine
//  n: long integer, the number of candies to produce
//
func minimumPasses(c int64, m int64, w int64, p int64, n int64) int64 {
  // Produce for this round
  c := m*w
  if n-c <= 0 {
    return 0
  }
  // Now choose how many to buy
  budget := int(c/p)

  minRounds := int64(0)

  for i:=0;i<=budget;i++ {
    // workers to buy
    for j:=0;j<=budget-i;j++{
    }
  }
}

// Complete the minimumPasses function below.
//  m: long integer, the starting number of machines
//  w: long integer, the starting number of workers
//  p: long integer, the cost of a new hire or a new machine
//  n: long integer, the number of candies to produce
func minimumPasses(m int64, w int64, p int64, n int64) int64 {
  // naieve approach: recursive:
  return MinimumPasses(n, m, w, p)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    mwpn := strings.Split(readLine(reader), " ")

    m, err := strconv.ParseInt(mwpn[0], 10, 64)
    checkError(err)

    w, err := strconv.ParseInt(mwpn[1], 10, 64)
    checkError(err)

    p, err := strconv.ParseInt(mwpn[2], 10, 64)
    checkError(err)

    n, err := strconv.ParseInt(mwpn[3], 10, 64)
    checkError(err)

    result := minimumPasses(m, w, p, n)

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
