class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s: string, t: string): boolean {
        const map: Map<string, number> = new Map();

        for(let i = 0; i < s.length; i++) {
            const letter: string = s[i]
            if(map.has(letter)) {
                map.set(letter, map.get(letter) + 1)
            } else {
                map.set(letter, 1)
            }
        }

        for(let i = 0; i < t.length; i++) {
            const letter: string = t[i]
            if(map.has(letter)) {
                map.set(letter, map.get(letter) - 1)
            } else {
                return false;
            }
        }

        for(const [key, value] of map) {
            if(value != 0) {
                return false;
            }
        }
        return true;
    }
}
