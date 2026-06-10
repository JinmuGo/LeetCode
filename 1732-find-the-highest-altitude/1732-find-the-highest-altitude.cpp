class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        int hi = 0;
        int val = 0;

        for (int elem : gain) {
            val += elem;

            if (hi < val) {
                hi = val;
            }
        }

        return hi;
    }
};