class Solution {
public:
    string removeStars(string s) {
        vector<char> stack;

        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '*') {
                stack.pop_back();
                continue;
            }
            stack.push_back(s[i]);
        }

        string rv;

        for (char c : stack) {
            rv += c;
        }
        
        return rv;
    }
};