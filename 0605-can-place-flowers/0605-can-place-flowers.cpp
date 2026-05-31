class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        for (int i = 0; i < flowerbed.size(); i++) {
            if ((i == 0 || !flowerbed[i - 1]) && !flowerbed[i] &&
                (i == flowerbed.size() - 1 || !flowerbed[i + 1])) {
                n -= 1;
                flowerbed[i] = 1;
            }
        }
        return n <= 0;
    }
};