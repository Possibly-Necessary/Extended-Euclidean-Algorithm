// Go script that calculates the Extended Euclidean Algorithm

package main

import (
	"fmt"
)

// function that computes the Extended Euclidian Algorithm
func EXTD_EUCLID(a, b int) (r, s, t int, err error) {

	var (
		k int // For the index
		q int // Temp storage for q
	)

	k = 2 //index

	if (a >= 0) && (b > 0) && (a >= b) {

		// Define the arrays/slices  r, s, t with length k
		r := make([]int, k)
		s := make([]int, k)
		t := make([]int, k)

		// Insert both a and b into the array `r'`
		r[0] = a
		r[1] = b
		s[0] = 1
		s[1] = 0
		t[0] = 0
		t[1] = 1

		// while rk-1 (which is b) is not 0
		// Go lang does not have while-loops
		// for loops are used instead
		//index := k-1

		for r[k-1] != 0 {
			q = r[k-2] / r[k-1]
			r = append(r, r[k-2]%r[k-1])
			s = append(s, s[k-2]-(q*s[k-1]))
			t = append(t, t[k-2]-(q*t[k-1]))
			k += 1
		}

		return r[len(r)-2], s[len(s)-2], t[len(t)-2], nil

	} else {

		return 0, 0, 0, fmt.Errorf("Conditions Are Not Satisfied...")

	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to handle user input
func userInpt() (a, b int) {

	fmt.Println("Enter two integers, ")

	// Enforce only integer value input for a
	for true {
		fmt.Print("a: ")

		_, err1 := fmt.Scanln(&a)
		if err1 == nil {
			break
		}
		fmt.Println("Invalid input. Enter a strictly integer value.")
		var dump1 string
		fmt.Scanln(&dump1) // Clear STDIN buffer
	}

	// Enforce only integer value input for b
	for true {
		fmt.Print("b: ")
		_, err2 := fmt.Scanln(&b)
		if err2 == nil {
			break
		}
		fmt.Println("Invalid input. Enter a strictly integer value.")
		var dump2 string
		fmt.Scanln(&dump2) // Clear STDIN buffer
	}

	return a, b
}

func main() {

	a, b := userInpt()
	fmt.Println(a, b)

	// Call the Extended Euclidean Algorithm
	r, s, t, err := EXTD_EUCLID(a, b)
	if err != nil { //should return nill if there is no error
		fmt.Println(err)
	} else {
		fmt.Printf("r: %d, s: %d, t: %d\n", r, s, t)
	}

	// check with the Extended Euclidean equation if the values hold
	// gcd(a,b) = s.a + t.b
	gcdResult := gcd(a, b)

	if gcdResult == s*a+t*b {

		fmt.Printf("gcd(%d, %d) = (%d) * %d + (%d) * %d\n", a, b, s, a, t, b)
		fmt.Printf("%d = %d\n", gcdResult, s*a+t*b)
		fmt.Println("Equation Holds.")

	} else {
		fmt.Printf("gcd(%d, %d) = (%d) * %d + (%d) * %d\n", a, b, s, a, t, b)
		fmt.Printf("%d = %d\n", gcdResult, s*a+t*b)
		fmt.Println("Equation Does Not Hold.")
	}

}
