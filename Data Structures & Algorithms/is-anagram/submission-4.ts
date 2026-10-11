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
            
                map.set(letter, (map.get(letter) ?? 0) + 1)
        }

        for(let i = 0; i < t.length; i++) {
            const letter: string = t[i]
            const count = map.get(letter) ?? 0
            if(count <= 0) {
                return false;
            }
            map.set(letter, count - 1);
        }

        for(const [key, value] of map) {
            if(value != 0) {
                return false;
            }
        }
        return true;
    }
}
