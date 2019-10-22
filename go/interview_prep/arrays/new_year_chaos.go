package main
/**
It's New Year's Day and everyone's in line for the Wonderland rollercoaster
ride! There are a number of people queued up, and each person wears a sticker
indicating their initial position in the queue. Initial positions increment by 
from  at the front of the line to  at the back.

Any person in the queue can bribe the person directly in front of them to swap
positions. If two people swap positions, they still wear the same sticker
denoting their original places in line. One person can bribe at most two others.
For example, if $n=8$ and $Person 5$ bribes $Person 4$, the queue will look like
this: $[1,2,3,5,4,6,7,8]$.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func findMinimumBribes(state []int32, numBribes []int32, goal []int32) {
}

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    /**
    Naïeve approach: Number of bribes needed is total number of positions
    people shifted in front of their original position.

    Does not work: 
    [1,2,3] => [3,2,1]
    [1,2,3] => [3,1,2] => [3,2,1]

    Three bribes while only one number moved by two.

    Alternative:
    |moved forward| + |moved backward|
    */
    var cp []int32
    var bribes []int32

    // Copy the bribes list
    for _, val := range a {
        cp = append(cp, val)
        bribes = append(bribes, 0)
    }

    minBribes := findMinimumBribes(q, bribes, goal)

    fmt.Printf("%d", bribes)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
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
