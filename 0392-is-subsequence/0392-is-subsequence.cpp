class Solution {
public:
    bool isSubsequence(string s, string t) {
        if (s.size() == 0) {
            return true;
        }
        int s_ptr = 0;
        int t_ptr = 0;
        bool ans = false;

        while (t_ptr < t.size()) {
            if (s.at(s_ptr) == t.at(t_ptr)) {
                s_ptr++;
            }
            if (s_ptr == s.size()) {
                ans = true;
                break ;
            }
            t_ptr++;
        }

        return ans;
    }
};