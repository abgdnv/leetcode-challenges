package main

//921. Minimum Add to Make Parentheses Valid
// A parentheses string is valid if and only if:
// * It is the empty string,
// * It can be written as AB (A concatenated with B), where A and B are valid strings, or
// * It can be written as (A), where A is a valid string.
// You are given a parentheses string s. In one move, you can insert a parenthesis at any position of the string.

// * For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis to be "())))".

// Return the minimum number of moves required to make s valid.

// Example 1:

// Input: s = "())"
// Output: 1
// Example 2:

// Input: s = "((("
// Output: 3

// Constraints:

// 1 <= s.length <= 1000
// s[i] is either '(' or ')'.
// *************************************************************************************************************
// Solution:
// count needed closing brakets: ++ if character is '(', otherwise --
// if there is no opening bracket for closing bracket, increase the counter of needed opening brackets and reset
// the counter of needed closing brackets

// the result is the sum of two counters

// time complexity  O(n)
// space complexity O(1)

func minAddToMakeValid(s string) int {
	oMoves, cMoves := 0, 0
	for _, c := range s {
		if c == '(' {
			cMoves++
		} else {
			cMoves--
		}
		if cMoves < 0 {
			oMoves++
			cMoves = 0
		}
	}
	return oMoves + cMoves
}
