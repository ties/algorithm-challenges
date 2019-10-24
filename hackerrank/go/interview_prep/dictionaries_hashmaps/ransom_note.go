package main
/**
Harold is a kidnapper who wrote a ransom note, but now he is worried it will be
traced back to him through his handwriting. He found a magazine and wants to
knowaf he can cut out whole words from it and use them to create an untraceable
replica of his ransom note. The words in his note are case-sensitive and he must
use only whole words available in the magazine. He cannot use substrings or
concatenation to create the words he needs.

Given the words in the magazine and the words in the ransom note, print Yes if
he can replicate his ransom note exactly using whole words from the magazine;
otherwise, print No.

For example, the note is "Attack at dawn". The magazine contains only "attack
at dawn". The magazine has all the right words, but there's a case mismatch.
The answer is "No".

Function Description
--------------------

Complete the checkMagazine function in the editor below. It must print "Yes" if
the note can be formed using the magazine, or "No".

checkMagazine has the following parameters:
	magazine: an array of strings, each a word in the magazine
	note: an array of strings, each a word in the ransom note
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
    // Build the map
    m := make(map[string]int)

    for _, val := range(magazine) {
        v := m[val]
        m[val] = v + 1
    }

    // Check if the strings are in the magazine
    // exit early if it fails.
    for _, val := range(note) {
        v := m[val]
        if v == 0 {
            fmt.Println("No")
            return
        }

        m[val] = v-1
    }
    fmt.Println("Yes")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    mn := strings.Split(readLine(reader), " ")

    mTemp, err := strconv.ParseInt(mn[0], 10, 64)
    checkError(err)
    m := int32(mTemp)

    nTemp, err := strconv.ParseInt(mn[1], 10, 64)
    checkError(err)
    n := int32(nTemp)

    magazineTemp := strings.Split(readLine(reader), " ")

    var magazine []string

    for i := 0; i < int(m); i++ {
        magazineItem := magazineTemp[i]
        magazine = append(magazine, magazineItem)
    }

    noteTemp := strings.Split(readLine(reader), " ")

    var note []string

    for i := 0; i < int(n); i++ {
        noteItem := noteTemp[i]
        note = append(note, noteItem)
    }

    checkMagazine(magazine, note)
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
