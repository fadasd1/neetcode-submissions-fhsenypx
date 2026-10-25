class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for word in strs:
            res += f"{len(word)}#{word}"
        return res

    def decode(self, s: str) -> List[str]:
        res = []
        i = 0
        num = ""
        
        while i < len(s):
            if s[i] != "#":
                num += s[i]
            elif s[i] == "#":
                if not num:
                    length = 0
                else:
                    length = int(num)
                    num = ""
                
                res.append(s[i+1:i+1+length])
                i += length
            
            i += 1
        return res
