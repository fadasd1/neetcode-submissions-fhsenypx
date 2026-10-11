class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s: string, t: string): boolean {
        const arr: number[] = new Array(26).fill(0);
        for(let i = 0; i < s.length; i++) {
            const up = s.charCodeAt(i) - 97;
            arr[up]++;
        }

        for(let i = 0; i < t.length; i++) {
            const down = t.charCodeAt(i) - 97
            arr[down]--;
        }

        for(const num of arr) {
            if(num != 0) {
                return false;
            }
        }
        return true;
    }
}
