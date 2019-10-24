package main
/**
Given two strings, determine if they share a common substring. A substring may
be as small as one character.

For example, the words "a", "and", "art" share the common substring . The words
"be" and "cat" do not share a substring.

Function Description
--------------------

Complete the function twoStrings in the editor below. It should return a string,
either YES or NO based on whether the strings share a common substring.

twoStrings has the following parameter(s):
    s1, s2: two strings to analyze .

Input Format

The first line contains a single integer $p$, the number of test cases.
The following $p$ pairs of lines are as follows:

    The first line contains string $s1$.
    The second line contains string $s2$.

Constraints
-----------

$s1$ and $s2$ consist of characters in the range ascii[a-z].

Output Format
-------------

For each pair of strings, return YES or NO.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
    // Note that since one-character substring are also a match,
    // we only need to match on this case (since one char is a sub-case of
    // any longer string.

    // Using a map since that is the topic of this folder,
    // could have used [26]bool
    m := make(map[byte]bool)

    for i:=0; i<len(s1); i++ {
        m[s1[i]] = true
    }

    for i:=0; i<len(s2); i++ {
        if m[s2[i]] {
            return "YES"
        }
    }

    return "NO"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s1 := readLine(reader)

        s2 := readLine(reader)

        result := twoStrings(s1, s2)

        fmt.Fprintf(writer, "%s\n", result)
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
