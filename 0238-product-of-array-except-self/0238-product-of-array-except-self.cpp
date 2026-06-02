class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int len = nums.size();
        vector<int> ans(len, 1);

        for (int i = 1; i < len; i++) {
            ans[i] = ans[i - 1] * nums[i - 1];
        }


        int acc_right = 1;
        for (int j = len - 2; j >= 0; j--) {
            acc_right *= nums[j + 1];
            ans[j] = ans[j] * acc_right;
        }

        return ans;
    }
};