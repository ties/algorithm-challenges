package main
/**
A binary tree is a tree which is characterized by one of the following
properties:

  * It can be empty (null).
  * It contains a root node only.
  * It contains a root node with a left subtree, a right subtree, or both.
    These subtrees are also binary trees.

In-order traversal is performed as
  * Traverse the left subtree.
  * Visit root.
  * Traverse the right subtree.

For this in-order traversal, start from the left child of the root node and keep
exploring the left subtree until you reach a leaf. When you reach a leaf, back
up to its parent, check for a right child and visit it if there is one. If there
is not a child, you've explored its left and right subtrees fully. If there is a
right child, traverse its left subtree then its right in the same manner. Keep
doing this until you have traversed the entire tree. You will only store the
values of a node as you visit when one of the following is true:

  * it is the first node visited, the first time visited
  * it is a leaf, should only be visited once
  * all of its subtrees have been explored, should only be visited once while
    this is true
  * it is the root of the tree, the first time visited

**Tree based on [0]**
[0]: https://golang.org/doc/play/tree.go

Swapping:
---------
Swapping subtrees of a node means that if initially node has left subtree L and
right subtree R, then after swapping, the left subtree will be R and the right
subtree, L.

For example, in the following tree, we swap children of node 1.
```
                                Depth
    1               1            [1]
   / \             / \
  2   3     ->    3   2          [2]
   \   \           \   \
    4   5           5   4        [3]
```
In-order traversal of left tree is 2 4 1 3 5 and of right tree is 3 5 1 2 4.

Swap operation:
---------------

We define depth of a node as follows:

The root node is at depth 1.
If the depth of the parent node is d, then the depth of current node will be
d+1.

Given a tree and an integer, k, in one operation, we need to swap the subtrees
of all the nodes at each depth h, where h ∈ [k, 2k, 3k,...]. In other words, if
h is a multiple of k, swap the left and right subtrees of that level.

You are given a tree of n nodes where nodes are indexed from [1..n] and it is
rooted at 1. You have to perform t swap operations on it, and after each swap
operation print the in-order traversal of the current state of the tree.

Function Description
--------------------

Complete the swapNodes function in the editor below. It should return a
two-dimensional array where each element is an array of integers representing
the node indices of an in-order traversal after a swap operation.

swapNodes has the following parameter(s):
  * indexes: an array of integers representing index values of each , beginning
    with , the first element, as the root.
  * queries: an array of integers, each representing a  value.

Input Format
------------
  * The first line contains n, number of nodes in the tree.
  * Each of the next n lines contains two integers, a b, where a is the index
    of left child, and b is the index of right child of ith node.

Note: -1 is used to represent a null node.

The next line contains an integer, t, the size of .
Each of the next t lines contains an integer , each being a value .

Output Format
-------------
For each k, perform the swap operation and store the indices of your in-order
traversal to your result array. After all swap operations have been performed,
return your result array for printing.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int32
	Right *Tree
}

// Walk traverses a tree in-order,
// sending each Value on a channel.
func Walk(t *Tree, ch chan int32) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan int32 {
        ch := make(chan int32)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}


func InOrder(root *Tree) []int32 {
  res := []int32{}
  /*In-order traversal*/
  c1 := Walker(root)
  for {
    v1, ok1 := <- c1
    if !ok1 {
      break
    }
    res = append(res, v1)
  }

  return res
}

/**
Swap children if current depth == k, depth(root) = 1
*/
func SwapAtLevel(root *Tree, k int32, curDepth int32) {
  if curDepth % k == 0 {
    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp
  }
  // If has lhs/rhs: recurse
  if root.Left != nil {
    SwapAtLevel(root.Left, k, curDepth + 1)
  }
  if root.Right != nil {
    SwapAtLevel(root.Right, k, curDepth + 1)
  }
}


/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
  // Build the tree
  root := &Tree{nil, 1, nil}

  stack := []*Tree{root}
  for _, row := range(indexes) {
    l := len(stack)
    if l == 0 {
      break
    }
    cur := stack[0]
    stack = stack[1:]

    if row[0] != -1 {
      cur.Left = &Tree{nil, row[0], nil}
    }
    if row[1] != -1 {
      cur.Right = &Tree{nil, row[1], nil}
    }

    // fmt.Printf("%v with lhs(%v) rhs(%v)\n", cur, cur.Left, cur.Right)

    if row [0] != -1 {
      stack = append(stack, cur.Left)
    }
    if row[1] != -1 {
      stack = append(stack, cur.Right)
    } 
  }

  // Construction of tree works.
  // Swap operations:
  out := [][]int32{}

  for _, k := range(queries) {
    // Swap at multiples of K
    // depth(root) = 1
    SwapAtLevel(root, k, 1)
    // Store in-order traversal
    out = append(out, InOrder(root))
  }

  return out
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var indexes [][]int32
    for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
        indexesRowTemp := strings.Split(readLine(reader), " ")

        var indexesRow []int32
        for _, indexesRowItem := range indexesRowTemp {
            indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
            checkError(err)
            indexesItem := int32(indexesItemTemp)
            indexesRow = append(indexesRow, indexesItem)
        }

        if len(indexesRow) != int(2) {
            panic("Bad input")
        }

        indexes = append(indexes, indexesRow)
    }

    queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    var queries []int32

    for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
        queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        queriesItem := int32(queriesItemTemp)
        queries = append(queries, queriesItem)
    }

    result := swapNodes(indexes, queries)

    for resultRowItr, rowItem := range result {
        for resultColumnItr, colItem := range rowItem {
            fmt.Fprintf(writer, "%d", colItem)

            if resultColumnItr != len(rowItem) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        if resultRowItr != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
