package main

import "fmt"

func isPalindrome(n int) bool {
	org_num := n
	rev_num := 0

	for n > 0 {
		reminder := n % 10
		rev_num = rev_num*10 + reminder
		n = n / 10
	}

	if org_num == rev_num {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	return org_num == rev_num
}

func main() {
	isPalindrome(121)
}
