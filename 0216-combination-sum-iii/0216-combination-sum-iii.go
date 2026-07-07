func combinationSum3(k int, n int) [][]int {
    sum := 0
    for i := 1; i <= k; i++ {
        sum += i
    } 
    if sum > n {
        return [][]int{}
    }

    var res [][]int
    var curArr []int

    backtracking(1, k, n, curArr, &res)

    return res
}

func backtracking(L, k, target int, curArr []int, res *[][]int) {
    if len(curArr) == k  {
        if target == 0 {
            temp := make([]int, len(curArr))
            copy(temp, curArr)
            *res = append(*res, temp)
        }
        return
    }

    if target < 0 {
        return
    }
    

    for i := L; i <= 9; i++ {
        curArr = append(curArr, i)
        backtracking(i + 1, k, target - i, curArr, res)
        curArr = curArr[:len(curArr) - 1]
    }
}