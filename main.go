package main

import (
	"fmt"
)

// Join takes two strings, str1 and str2, and appends characters from str2 to str1 if they don't already exist in str1.
// Create a new string, str3, initialized with str1.
// Convert str1 into a slice of strings, str1Arr, to facilitate character comparison.
// Loop through each character in str2.
// Initialize a boolean variable, check, to track if the character from str2 exists in str1.
// Loop through each character in str1 to check for a match with the current character from str2.
// If the character from str2 matches a character in str1, set check to false.
// If check is true, meaning the character from str2 does not exist in str1, append it to str3.
// Also add the character to str1Arr to prevent duplicates.
// Return the modified string, str3
func Join(str1 string, str2 string) string {

	str3 := str1

	var str1Arr []string
	for _, v := range str1 {
		str1Arr = append(str1Arr, string(v))
	}

	for _, v := range str2 {
		
		check := true
		
		for i := 0; i < len(str1Arr); i++ {
			
			if string(v) == str1Arr[i] {
				check = false
			}
		}
	
		if check {
			str3 += string(v)
			
			str1Arr = append(str1Arr, string(v))
		}
	}

	return str3
}

func main() {

	s := "abcdef"
	s2 := "xyzaxtxxxxxxxxxxxxxxxxxxxxxxxxxxxxn"
	fmt.Println(Join(s, s2))
}
