func letterCombinations(digits string) []string {
    mapping := map[string][]string{
        "2": {"a", "b", "c"},
        "3": {"d", "e", "f"},
        "4": {"g", "h", "i"},
        "5": {"j", "k", "l"},
        "6": {"m", "n", "o"},
        "7": {"p", "q", "r", "s"},
        "8": {"t", "u", "v"},
        "9": {"w", "x", "y", "z"},
    }

    arr := []string{}
    n := len(digits)

    for _, val := range mapping[string(digits[0])] {
        dfs(1, n, val, digits, mapping, &arr)
    }

    return arr
}

func dfs(L, n int, cur, digits string, mapping map[string][]string, arr *[]string) {
    if L == n {
        *arr = append(*arr, cur)
        return 
    }

    for _, val := range mapping[string(digits[L])] {
        str := cur + val
        dfs(L + 1, n, str, digits, mapping, arr)
    }
}