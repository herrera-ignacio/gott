package main

/**
Harsh was recently gifted a book consisting of NN pages. Each page contains exactly MM words printed on it. As he was bored, he decided to count the number of words in the book.

Help Harsh find the total number of words in the book.

Input Format
The first line of input will contain a single integer TT, denoting the number of test cases.
Each test case consists of two space-separated integers on a single line, NN and MM — the number of pages and the number of words on each page, respectively.
*/

import (
	"fmt"
)

func main() {
	var t, n, m int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d%d", &n, &m)
		fmt.Println(n * m)
	}
}
