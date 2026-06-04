class Solution {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        double windowVal = 0;
        double rv;

        for (int i = 0; i < k; i++) {
            windowVal += nums[i];
        }
        
        rv = windowVal / k;

        for (int j = 1; j + k <= nums.size(); j++) {
            windowVal = windowVal - nums[j - 1] + nums[j + k - 1];
            double val = windowVal / k;

            if (rv < val) {
                rv = val;
            }
        }

        return rv;
    }
};