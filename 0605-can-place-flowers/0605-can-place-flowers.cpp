class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        bool ans = false;

        for (int i = 0; i < flowerbed.size(); i++) {
            if (n == 0) {
                ans = true;
                break;
            }
            if ((i == 0 || !flowerbed[i - 1]) && !flowerbed[i] &&
                (i == flowerbed.size() - 1 || !flowerbed[i + 1])) {
                n -= 1;
                flowerbed[i] = 1;
            }
        }

        if (n == 0) {
            ans = true;
        }
        return ans;
    }
};