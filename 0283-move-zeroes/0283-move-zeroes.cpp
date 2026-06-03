class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int left = 0;
        int right = nums.size();

        while (left < right) {
            int left_elem = nums.at(left);
            if (left_elem == 0) {
                nums.erase(nums.begin() + left);               
                nums.push_back(0);
                right--;
                left--;
            }
            left++;
        }
    }
};