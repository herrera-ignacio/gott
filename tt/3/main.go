package main

import (
	"fmt"
)

/*
Problem

Chef appeared for a placement test.

There is a problem worth XXX points. Chef finds out that the problem has exactly 101010 test cases. It is known that each test case is worth the same number of points.

Chef passes NNN test cases among them. Determine the score Chef will get.

NOTE: See sample explanation for more clarity.
Input Format

    First line will contain TTT, number of test cases. Then the test cases follow.
    Each test case contains of a single line of input, two integers XXX and NNN, the total points for the problem and the number of test cases which pass for Chef's solution.

Output Format

For each test case, output the points scored by Chef.
*/

func main() {
	var t, points, amountSolved int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d%d", &points, &amountSolved)
		fmt.Println(eval(points, amountSolved))
	}
}

func eval(points, amountSolved int) int {
	pointsPerCase := points / 10
	totalPoints := amountSolved * pointsPerCase
	return totalPoints
}
