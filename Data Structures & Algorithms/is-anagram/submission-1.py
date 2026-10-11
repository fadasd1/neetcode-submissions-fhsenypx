class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        dictS = dict()
        

        for char in s:
            c = dictS.get(char)
            if c is None:
                dictS[char] = 1
            else:
                dictS[char] = c + 1

        for char in t:
            c = dictS.get(char)
            if c is None or c < 0:
                return False
            else:
                dictS[char] = c - 1

        for value in dictS.values():
            if value != 0:
                return False
        return True        

        