class Solution {
public:
    int maxVowels(string s, int k) {
        int sum = 0;

        for (int i = 0; i < k; i++) {
            sum += isVowel(s[i]);
        }

        int max = sum;

        for (int j = k; j < s.size(); j++) {
            sum = sum - isVowel(s[j - k]) + isVowel(s[j]);

            if (max < sum) {
                max = sum;
            }
        }

        return max;
    }

    bool isVowel(char c) {
        if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
            return true;
        }
        return false;
    }
};