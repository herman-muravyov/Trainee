package main

import (
	"fmt"
	"task/palindrome"
	"task/mapping"
	"task/variables"
	"task/stringers"
)

func main() {
	fmt.Println("Palindrome solution:")
	fmt.Println(palindrome.IsPalindromeFirst(variables.PalindromeStr))
	fmt.Println(palindrome.IsPalindromeSecond(variables.PalindromeStr))
	fmt.Println(palindrome.IsPalindromeThird(variables.PalindromeStr))

	fmt.Println("\nMap task solution:")
	fmt.Println(mapping.Mapping(variables.Slice))

	fmt.Println("\nString task solutions:")
	fmt.Println(stringers.ReverseWordsFirst(variables.Str))
	fmt.Println(stringers.ReverseWordsSecond(variables.Str))
}