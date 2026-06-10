class Solution {
public:
    string removeStars(string s) {
        string rv;

        for (char c : s) {
            if (c == '*') {
                rv.pop_back();
            } else {
                rv.push_back(c);
            }
            
        }
        return rv;
    }
};