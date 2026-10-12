class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        alphabet = [0] * 26
        result = dict()

        for string in strs:
            for c in string:
                value = ord(c) - 97
                alphabet[value] += 1
            alphabet = tuple(alphabet)
            result.setdefault(alphabet, []).append(string)
            alphabet = [0] * 26
        final = []
        for value in result.values():
            final.append(value)
        return final
        