class Solution {
public:
    int tribonacci(int n) {
        vector<int> dp;
        dp.push_back(0);
        dp.push_back(1);
        dp.push_back(1);

        if (n <= 2) {
            return dp[n];
        }

        for (int i = 3; i < n; i++) {
            dp.push_back(dp[i - 1] + dp[i - 2] + dp[i - 3]);
        }

        return dp[n - 1] + dp[n - 2] + dp[n - 3];
    }
};