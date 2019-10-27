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

// From: https://play.golang.org/p/SmzvkDjYlb
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Complete the minTime function below.
func minTime(machines []int64, goal int64) int64 {
    // // Find lcm of all machines
    // // since we can apply lcm iterations at once.
    // var lcm int64
    // if len(machines) == 1 {
    //     lcm = machines[0]
    // } else {
    //     lcm = LCM(machines[0], machines[1], machines[2:]...)
    // }

    // fmt.Printf("LCM: %v", lcm)

    // return int64(0)
    built := int64(0)
    t := int64(1)

    for built < goal {
        for i:=0; i < len(machines); i++ {
            if t % machines[i] == 0 {
                built += 1
            }
        }

        t += 1
    }

    return t-1
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
