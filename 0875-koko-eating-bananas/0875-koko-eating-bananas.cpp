class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        int k = 0;
        int left = 1;
        int right = 1e9;

        while (left <= right) {
            int mid = left + (right - left) / 2;
            long long sum = 0;
            for (int i : piles) {
                sum += i / mid;
                if (i % mid) {
                    sum++;
                }
            }
            if (sum <= h) {
                right = mid - 1;
                k = mid;
            } else {
                left = mid + 1;
            }
        }
        return k;
    }
};