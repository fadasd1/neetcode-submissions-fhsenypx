class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        let signature: number[] = new Array(26).fill(0);
        const map: Map<string, string[]> = new Map<string, string[]>();

        for(const str of strs) {
            for(const letter of str) {
                const code: number = letter.charCodeAt(0) - 97;
                signature[code] += 1;
            }
            const signatureStr: string = signature.join(",")
            
            const fromMap: string[] | undefined = map.get(signatureStr);
            if(fromMap) {
                fromMap.push(str)
            } else {
                map.set(signatureStr, [str])
            }

            signature = new Array(26).fill(0)
        }

        const result: string[][] = []

        for(const array of map.values()) {
            result.push(array)
        }

        return result;
    }
}
