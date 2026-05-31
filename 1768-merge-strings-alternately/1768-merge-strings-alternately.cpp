class Solution {
public:
    string mergeAlternately(string word1, string word2) {
        int i = 0;
        int j = 0;
        string str = "";

        while (i < word1.length() && j < word2.length()) {
           str += word1[i++]; 
           str += word2[j++]; 
        }

        str += word1.substr(i);
        str += word2.substr(j);

        return str;
    }
};