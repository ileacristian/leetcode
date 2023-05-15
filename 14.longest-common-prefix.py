# Write a function to find the longest common prefix string amongst an array of strings.
# If there is no common prefix, return an empty string "".
# https://leetcode.com/problems/longest-common-prefix/


class Solution(object):
    def longestCommonPrefix(self, strs):

        def commonPrefix(str1, str2):
            i=0
            j=0
            cs=str("")
            stop=False
            while i in range(0,len(str1)) and j in range(0,len(str2)) and stop==False:
                i+=1
                j+=1
                S1 = str1[i-1]
                S2 = str2[j-1]
                if S1==S2:
                    cs=cs+S1
                else:
                    stop=True
            else:
                return(cs)


        if len(strs) == 0:
            return None

        if len(strs) == 1:
            return strs[0]

        s = sorted(strs)
        return commonPrefix(s[0], s[-1])
        """
        :type strs: List[str]
        :rtype: str
        """



        return minim(strs)