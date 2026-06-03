class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int right = 0;

        for (int left = 0; left < nums.size(); left++) {
            if (nums[left] != 0) {
                nums[right++] = nums[left];
            }
        }

        while (right < nums.size()) {
            nums[right++] = 0;
        }
    }
};