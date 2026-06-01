class Solution {
public:
    string reverseVowels(string s) {
        int left = 0;
        int right = s.length() - 1;

        while (left <= right) {
            bool leftVowel = isVowel(s[left]);

            if (leftVowel) {
                while (!isVowel(s[right])) {
                    right--;
                }

                std::swap(s[left], s[right]);
                right--;
            }
            left++;
        }

        return s;
    }

    bool isVowel(char c) {
        int lower_c = tolower(c);
        if (lower_c == 'a' || lower_c == 'e' || lower_c == 'i' ||
            lower_c == 'o' || lower_c == 'u') {
            return true;
        }
        return false;
    }
};