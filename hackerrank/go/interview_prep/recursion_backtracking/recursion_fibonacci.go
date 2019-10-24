package main
/**
The Fibonacci Sequence

The Fibonacci sequence appears in nature all around us, in the arrangement of
seeds in a sunflower and the spiral of a nautilus for example.

The Fibonacci sequence begins with  and  as its first and second terms. After
these first two elements, each subsequent element is equal to the sum of the
previous two elements.

Programmatically:
-----------------
$$fibonacci(0) == 0$$
$$fibonacci(1) == 1$$
$$fibonacci(n) == fibonacci(n-1) + fibonacci(n-2)$$

Given , return the  number in the sequence.

As an example, . The Fibonacci sequence to  is . With zero-based indexing, .

Function Description
--------------------
Complete the recursive function $fibonacci$ in the editor below. It must return
the $n^th$ element in the Fibonacci sequence.

fibonacci has the following parameter(s):
	n: the integer index of the sequence to return
*/
import "fmt"

/** Vanilla recursive fibonacci */
func fibonacci(n int) int {
    if n == 0 {
	    return 0
    } else if n == 1 {
	    return 1
    }
    return fibonacci(n-2)+fibonacci(n-1)
}

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    fmt.Println(fibonacci(n))
}

