class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        arr = [0] * 26
        for char in s:
            up = ord(char) - 97
            arr[up] += 1
        for char in t:
            down = ord(char) - 97
            arr[down] -= 1
        for num in arr:
            if num is not 0:
                return False
        return True        

        