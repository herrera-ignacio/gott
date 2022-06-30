package main

import (
	"fmt"
)

/*
Problem

Chef wants to appear in a competitive exam. To take the exam, there are following requirements:

    Minimum age limit is XXX (i.e. Age should be greater than or equal to XXX).
    Age should be strictly less than YYY.

Chef's current Age is AAA. Find whether he is currently eligible to take the exam or not.
Input Format

    First line will contain TTT, number of test cases. Then the test cases follow.
    Each test case consists of a single line of input, containing three integers X,Y,X, Y,X,Y, and AAA as mentioned in the statement.

Output Format

For each test case, output YES if Chef is eligible to give the exam, NO otherwise.

You may print each character of the string in uppercase or lowercase (for example, the strings YES, yEs, yes, and yeS will all b
*/

func main() {
	var t, a, b, age int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d%d%d", &a, &b, &age)
		eval(a, b, age)

	}
}

func eval(a, b, age int) {
	if age < a || age >= b {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
