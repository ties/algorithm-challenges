package main
/**
Two strings are anagrams of each other if the letters of one string can be
rearranged to form the other string. Given a string, find the number of pairs of
substrings of the string that are anagrams of each other.

For example $s= mom$, the list of all anagrammatic pairs is $[m,m], [mo,om]$ at
positions $[[0], [2]], [[0,1][1,2]]$ respectively.

Function Description
--------------------

Complete the function sherlockAndAnagrams in the editor below. It must return an
integer that represents the number of anagrammatic pairs of substrings in $s$.

sherlockAndAnagrams has the following parameter(s):
    s: a string .

Input Format
------------

The first line contains an integer $q$, the number of queries.
Each of the next $q$ lines contains a string $s$ to analyze.

Constraints
-----------
$$1 <= q <= 10$$
$$2 <= |s| <= 100$$

String $s$ contains only lowercase letters  ascii[a-z].

Output Format
-------------

For each query, return the number of unordered anagrammatic pairs.
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

// https://mariadesouza.com/2018/01/01/string-manipulation-in-go/
type sortRunes []rune 

func (s sortRunes) Less(i, j int) bool {     
    return s[i] < s[j] 
} 

func (s sortRunes) Len() int{    
    return len(s) 
} 

func (s sortRunes) Swap(i, j int) {   
    s[i], s[j] = s[j], s[i] 
} 

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
    // We work with positions
    // That are not the identity,
    // 
    // So for every pair of positions,
    // for every pair of the same length
    //
    // We check if it is an anagram with the current pair
    //
    // (will not be fast)
    anagrams := int32(0)

    m := make(map[string]int32)

    // For every string of every length
    // Sort it and count how often it occurs.
    for i:=0; i<=len(s); i++ {
        for j:=i+1; j<=len(s); j++ {
            st := []rune(s[i:j])
            sort.Sort(sortRunes(st))
            elem := string(st)

            m[elem] += 1
        }
    }

    for i:=0; i<=len(s); i++ {
        for j:=i+1; j<=len(s); j++ {
            // Sort outer slice
            o := []rune(s[i:j])
            sort.Sort(sortRunes(o))
            outer := string(o)

            // Look up how often this fragment was found
            // and substract its own occurence if len(outer) == 1
            anagrams += m[outer]-1
        }
    }

    return anagrams/2
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
        s := readLine(reader)

        result := sherlockAndAnagrams(s)

        fmt.Fprintf(writer, "%d\n", result)
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
