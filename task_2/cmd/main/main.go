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
	fmt.Println(palindrome.IsPalindromeFirst("hello"))
	fmt.Println(palindrome.IsPalindromeSecond("hello"))
	fmt.Println(palindrome.IsPalindromeThird("hello"))

	fmt.Println("\nMap task solution:")
	fmt.Println(mapping.Mapping(variables.Slice))

	fmt.Println("\nString task solutions:")
	fmt.Println(stringers.ReverseWordsFirst(variables.Str))
	fmt.Println(stringers.ReverseWordsSecond(variables.Str))
}