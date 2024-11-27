package main

// 1963. Minimum Number of Swaps to Make the String Balanced
// Description:
// You are given a 0-indexed string s of even length n.
// The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.
//
// A string is called balanced if and only if:
//
// * It is the empty string, or
// * It can be written as AB, where both A and B are balanced strings, or
// * It can be written as [C], where C is a balanced string.
//
// You may swap the brackets at any two indices any number of times.
//
// Return the minimum number of swaps to make s balanced.
//
//
// Example 1:
// Input: s = "][]["
// Output: 1
// Explanation: You can make the string balanced by swapping index 0 with index 3.
// The resulting string is "[[]]".
//
//
// Example 2:
// Input: s = "]]][[["
// Output: 2
// Explanation: You can do the following to make the string balanced:
// - Swap index 0 with index 4. s = "[]][][".
// - Swap index 1 with index 5. s = "[[][]]".
// The resulting string is "[[][]]".
//
//
// Example 3:
// Input: s = "[]"
// Output: 0
// Explanation: The string is already balanced.
//
//
// Constraints:
// n == s.length
// 2 <= n <= 10^6
// n is even.
// s[i] is either '[' or ']'.
// The number of opening brackets '[' equals n/2, and the number of closing brackets ']' equals n/2.

// *********************************************************************************************************
//
// Solution:
// First, we count closing brackets, that don't have coresponding opening brackets
// in other words, imagine if we removed all balanced brackets,
// we would end up with string like "]]]...]]][[[...[[["
// Each swap in this string will balance 2 closing brackets ("[]]...]]][[[...[[]"),
// The minimum number of swaps to make string balanced is round up of (closingBrackets/2),
// e.g.:
// "]["           - 1 swaps
// "]][["         - 1 swaps
// "]]][[["       - 2 swaps
// "]]]][[[["     - 2 swaps
// "]]]]][[[[["   - 3 swaps
// "]]]]]][[[[[[" - 3 swaps
// etc.
// Runtime complexity O(n)
// Memory complexity  O(1)

func minSwaps(s string) int {
	balance := 0
	closingBrackets := 0
	for _, c := range s {
		if c == '[' {
			balance++
		} else if c == ']' {
			balance--
		}
		if balance < 0 {
			closingBrackets++
			balance = 0
		}
	}
	return (closingBrackets + 1) / 2
}
