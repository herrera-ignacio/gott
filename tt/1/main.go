package main

import (
	"fmt"
)

/*
Problem

Chef wants to become fit for which he decided to walk to the office and return home by walking. It is known that Chef's office is XXX km away from his home.

If his office is open on 555 days in a week, find the number of kilometers Chef travels through office trips in a week.
Input Format

    First line will contain TTT, number of test cases. Then the test cases follow.
    Each test case contains of a single line consisting of single integer XXX.
*/

func main() {
	var t, x int
	fmt.Scanf("%d", &t)

        for i := 0; i < t; i++ {
                fmt.Scanf("%d", &x)
                fmt.Println(x * 10)
        }
}
