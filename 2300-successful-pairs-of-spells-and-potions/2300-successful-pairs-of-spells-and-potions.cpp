class Solution {
public:
    vector<int> successfulPairs(vector<int>& spells, vector<int>& potions, long long success) {
        const int n = spells.size();   
        const int m = potions.size();
        
        std::sort(potions.begin(), potions.end());

        for (int i = 0; i < n; i++) {
            long long spell = spells[i];
            
            int left = 0;
            int right = m - 1;
            int k = m;

            while (left <= right) {
                int mid = left + (right - left) / 2;

                long long val = spell * (long long)potions[mid];

                if (val >= success) {
                    k = mid;
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            }
            spells[i] = m - k;
        }

        return spells;
    }
};