/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function(s) {
    const arr = s.split(" ").map(s => s.trim()).filter(Boolean).reverse();

    return arr.join(" ");
};