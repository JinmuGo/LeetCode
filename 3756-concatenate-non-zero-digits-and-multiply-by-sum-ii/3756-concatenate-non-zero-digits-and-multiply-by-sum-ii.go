const MOD = 1000000007

func sumAndMultiply(s string, queries [][]int) []int {
n := len(s)
    
    originalToNew := make([]int, n+1)
    
    var filteredBytes []byte
    for i := 0; i < n; i++ {
        originalToNew[i] = len(filteredBytes)
        if s[i] != '0' {
            filteredBytes = append(filteredBytes, s[i])
        }
    }
    originalToNew[n] = len(filteredBytes)

    m := len(filteredBytes)
    sumPref := make([]int, m+1)
    numPref := make([]int, m+1)
    power10 := make([]int, m+1)
    
    power10[0] = 1
    for i := 0; i < m; i++ {
        elem := int(filteredBytes[i] - '0')
        sumPref[i+1] = sumPref[i] + elem
        numPref[i+1] = (numPref[i]*10 + elem) % MOD
        power10[i+1] = (power10[i] * 10) % MOD
    }

    ans := make([]int, len(queries))
    
    for i, query := range queries {
        newL := originalToNew[query[0]]
        newR := originalToNew[query[1]+1] - 1
        
        if newL > newR {
            ans[i] = 0
            continue
        }
        
        currentSum := sumPref[newR+1] - sumPref[newL]
        currentX := (numPref[newR+1] - (numPref[newL]*power10[newR-newL+1])%MOD + MOD) % MOD
        
        ans[i] = (currentX * currentSum) % MOD
    }

    return ans
}