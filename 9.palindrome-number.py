# Given an integer x, return true if x is a palindrome and false otherwise.
# https://leetcode.com/problems/palindrome-number

class Solution(object):
    def isPalindrome(self, x):
        return list(str(x)) == list(reversed(str(x)))