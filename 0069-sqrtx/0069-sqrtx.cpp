class Solution {
public:
    int mySqrt(int x) {
        if (x <= 1) {
            return x;
        }
        int left = 0;
        int right = x;

        while (left < right) {
            long mid = left + (right - left) / 2;

            if (mid * mid <= x) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return left - 1;
    }
};