/**
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function(word1, word2) {
let result = [];
    let i = 0;
    let j = 0;

    while (i < word1.length && j < word2.length) {
        result.push(word1[i]);
        result.push(word2[j]);
        i++;
        j++;
    }

    if (i < word1.length) {
        result.push(word1.slice(i));
    }
    if (j < word2.length) {
        result.push(word2.slice(j));
    }

    return result.join('');
};