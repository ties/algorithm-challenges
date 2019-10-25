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


type intarr []int32

/**
 * Sort protocol for int32
 */
func (f intarr) Len() int {
	return len(f)
}

func (f intarr) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f intarr) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}


func median(arr []int32) int32 {
	l := len(arr)
	data := make(intarr, l)
	copy(data, arr)
	// First step: sort array
	sort.Sort(data)

	// Odd number -> middle
	if l % 2 != 0 {
		return data[l/2]
	}

	return (data[l/2]+data[(l/2)+1])/2
}


// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	notifications := int32(0)

	for i:=int(d); i < len(expenditure); i++ {
		// Keep the trailing d transactions
		med := median(expenditure[i-int(d):i])

		if expenditure[i] >= 2*med {
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
