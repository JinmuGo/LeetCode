class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        int n = candies.size();
        int max = *max_element(candies.begin(), candies.end());
        vector<bool> vec(n, false);

        for (int i = 0; i < n; i++) {
            vec[i] = candies[i] + extraCandies >= max ? true : false;
        }

        return vec;
    }
};