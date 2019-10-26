package main
/**
HackerLand National Bank has a simple policy for warning clients about possible
fraudulent account activity. If the amount spent by a client on a particular day
is greater than or equal to $2x$ the client's median spending for a trailing
number of days, they send the client a notification about potential fraud. The
bank doesn't send the client any notifications until they have at least that
trailing number of prior days' transaction data.

Given the number of trailing days $d$ and a client's total daily expenditures
for a period of $n$ days, find and print the number of times the client will
receive a notification over all $n$ days.

For example, $d=3$ and $expenditures=[10,20,30,40,50]$. On the first three days,
they just collect spending data. At day $4$, we have trailing expenditures of 
$[10,20,30]$. The median is $20$ and the day's expenditure is $40$. Because
$40 >= 2*20$, there will be a notice. The next day, our trailing expenditures
are $[20,30,40] and the expenditures are $50$. This is less than $2x30$ so no
notice will be sent. Over the period, there was one notice sent.

Note: The median of a list of numbers can be found by arranging all the numbers
from smallest to greatest. If there is an odd number of numbers, the middle one
is picked. If there is an even number of numbers, median is then defined to be 
the average of the two middle values. (Wikipedia)

Function Description
--------------------

Complete the function activityNotifications in the editor below. It must return
an integer representing the number of client notifications.

activityNotifications has the following parameter(s):
	expenditure: an array of integers representing daily expenditures
	d: an integer, the lookback days for median spending
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


/**
 * Median of a sorted array.
 */
func Median(sortedArr []int) float32 {
	l := len(sortedArr)
	// Odd number -> middle
	if l % 2 != 0 {
		return float32(sortedArr[l/2])
	}

	return float32(sortedArr[(l/2)-1]+sortedArr[(l/2)])/2
}


func RemoveAndAdd(window []int, remove int, add int) []int {
    d := len(window)
    // find index
    idx := sort.SearchInts(window, remove)
    // remove from window by omitting
    if idx < d-2 {
	window = append(window[:idx], window[idx+1:]...)
    } else {
	window = window[:d-1]
    }
    // And now insert
    // Find insertion index
    insertIndex := sort.SearchInts(window, add)
    // Insert at index
    // 1: Save part of window to keep
    trail := make([]int, len(window[insertIndex:]))
    copy(trail, window[insertIndex:])
    // 2: Append at index
    window = append(window[:insertIndex], add)

    // If it was not the last pos, add the remainder
    if (insertIndex < d-1) {
	window = append(window, trail...)
    }
    return window
}


// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, depth int32) int32 {
    d := int(depth)

    notifications := int32(0)

    // Create window and sort it.
    window := make([]int, d)
    for idx, val := range(expenditure[:d]) {
	window[idx] = int(val)
    }
    sort.Ints(window)

    for i:=int(d); i < len(expenditure); i++ {
	// Remove oldest element from window
	// Only do this after iteration d (window) + 1 (current)
	if i-d-1 >= 0 {
	    window = RemoveAndAdd(window, int(expenditure[i-d-1]), int(expenditure[i-1]))
	}

	// Median of (sorted) window
	med := Median(window)

	if len(window) != d {
	    panic(fmt.Sprintf("Window size %d != %d", len(window), d))
	}
	// expenditure[i-1] is in window
	if window[sort.SearchInts(window, int(expenditure[i-1]))] != int(expenditure[i-1]) {
	    panic("Value not in window.")
	}

	if float32(expenditure[i]) >= 2*med {
		notifications += 1
	}
    }

    return notifications
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nd := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nd[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    dTemp, err := strconv.ParseInt(nd[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    expenditureTemp := strings.Split(readLine(reader), " ")

    var expenditure []int32

    for i := 0; i < int(n); i++ {
        expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
        checkError(err)
        expenditureItem := int32(expenditureItemTemp)
        expenditure = append(expenditure, expenditureItem)
    }

    result := activityNotifications(expenditure, d)

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
