package main
/**
You are planning production for an order. You have a number of machines that
each have a fixed number of days to produce an item. Given that all the machines
operate simultaneously, determine the minimum number of days to produce the
required order.

For example, you have to produce $goal=10$ items. You have three machines that
take $machines = [2,3,2]$ days to produce an item. The following is a schedule
of items produced:
```
Day Production  Count
2   2               2
3   1               3
4   2               5
6   3               8
8   2              10
```
It takes $8$ days to produce $10$ items using these machines.

Function Description
--------------------

Complete the minimumTime function in the editor below. It should return an
integer representing the minimum number of days required to complete the order.

minimumTime has the following parameter(s):
-------------------------------------------
  machines: an array of integers representing days to produce one item per
            machine
  goal:     an integer, the number of items required to complete the order
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


func builtAt(machines []int64, t int64) int64 {
    built := int64(0)
    for _, val := range(machines) {
        built += int64(t/val)
    }

    return built
}


// Complete the minTime function below.
func minTime(machines []int64, goal int64) int64 {
    /*
    Approach with calculating the lcm of the build time of
    all machines and then n*lcm time steps building
    the number built in lcm items at once
    does not work: This loop timed out.

    m: |machines|
    n: goal

    Alternative:
      * Calculate how many are built at t=n (O(machines))
      * Binary search: log(goal)
      => O(machines log goal)
    */
    l := int64(0)
    // Find r by searching until number produced > goal
    r := int64(1)
    for builtAt(machines, r) < goal {
        // Previous r was too low -> make it the lower bound.
        l = r
        r *= 2
    }

    // Binary search for first insertion position
    for l<r {
        mid := l+(r-l)/int64(2)

        builtMid := builtAt(machines, mid)

        if goal <= builtMid {
            r = mid
        } else {
            l = mid+1
        }
    }

    return l
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nGoal := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nGoal[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    goal, err := strconv.ParseInt(nGoal[1], 10, 64)
    checkError(err)

    machinesTemp := strings.Split(readLine(reader), " ")

    var machines []int64

    for i := 0; i < int(n); i++ {
        machinesItem, err := strconv.ParseInt(machinesTemp[i], 10, 64)
        checkError(err)
        machines = append(machines, machinesItem)
    }

    ans := minTime(machines, goal)

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
