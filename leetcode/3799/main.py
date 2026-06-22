from itertools import permutations

# 3799. Word Squares II
class Solution:
     def wordSquares(self, words: List[str]) -> List[List[str]]:
        n = len(words[0])

        res = []

        for top, left, bottom, right in permutations(words, 4):
            if (
                top[0] == left[0] and 
                top[3] == right[3] and 
                bottom[0] == left[3] and 
                bottom[3] == right[3] 
            ):
                res.append([top, left, bottom, right])
        
        res.sort()
        return res  