package main

import (
	"fmt"
	"regexp"
)

var phoneRegexp = regexp.MustCompile(`^(1(([35][0-9])|(47)|[8][01236789]))\d{8}$`)

func main() {

	phpneNumbers := []string{"13625002107", "18725002107", "11111111"}

	for _, phoneNumber := range phpneNumbers {
		fmt.Println("begin match %v:", phoneNumber, phoneRegexp.MatchString(phoneNumber))
	}

}
