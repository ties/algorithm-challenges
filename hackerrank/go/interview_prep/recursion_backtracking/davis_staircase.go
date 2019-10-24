package main
/**
Davis has a number of staircases in his house and he likes to climb each
staircase 1, 2, or 3 steps at a time. Being a very precocious child, he wonders
how many ways there are to reach the top of the staircase.

Given the respective heights for each of the $s$ staircases in his house, find
and print the number of ways he can climb each staircase, modulo $10^10+7$ on a
new line.

For example, there is $s=1$ staircase in the house that is $n=5$ steps high.
Davis can step on the following sequences of steps:
```
1 1 1 1 1
1 1 1 2
1 1 2 1 
1 2 1 1
2 1 1 1
1 2 2
2 2 1
2 1 2
1 1 3
1 3 1
3 1 1
2 3
3 2
```
There are $13$ possible ways he can take these $5% steps. $13%100000007 = 13$

Function Description
--------------------

Complete the stepPerms function in the editor below. It should recursively
calculate and return the integer number of ways Davis can climb the staircase,
modulo 10000000007.

stepPerms has the following parameter(s):
    n: an integer, the number of stairs in the staircase
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

const MOD = int64(10000000007)

func stepPermsInner(n int32, memoize map[int32]int64) int64 {
    // Calculate how many ways you can do the remaining n steps
    if n == 1 {
	return 1
    } else if n == 2 {
	return 2
    } else if n == 3 {
	// 3, 2+1, 1+2, 1+1+1
	return 4
    }

    if memoize[n] > 0 {
	return memoize[n]
    }

    res := stepPermsInner(n-3, memoize)+stepPermsInner(n-2, memoize)+stepPermsInner(n-1, memoize)
    memoize[n] = res
    return res
}

// Complete the stepPerms function below.
func stepPerms(n int32) int32 {
    memoize := make(map[int32]int64)
    return int32(stepPermsInner(n, memoize) % MOD)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    s := int32(sTemp)

    for sItr := 0; sItr < int(s); sItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        res := stepPerms(n)

        fmt.Fprintf(writer, "%d\n", res)
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
