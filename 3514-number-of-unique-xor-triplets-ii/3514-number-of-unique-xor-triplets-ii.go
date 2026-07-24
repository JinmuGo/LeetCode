func uniqueXorTriplets(nums []int) int {
    hasNum := make([]bool, 2048)
    uniqueNums := []int{}
    for _, num := range nums {
        if !hasNum[num] {
            hasNum[num] = true
            uniqueNums = append(uniqueNums, num)
        }
    }

    hasPairXor := make([]bool, 2048)
    n := len(uniqueNums)
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            hasPairXor[uniqueNums[i] ^ uniqueNums[j]] = true
        }
    }

    hasTripletXor := make([]bool, 2048)
    for x := 0; x < 2048; x++ {
        if !hasPairXor[x] {
            continue
        }
        for _, c := range uniqueNums {
            hasTripletXor[x ^ c] = true
        }
    }

    ans := 0
    for _, ok := range hasTripletXor {
        if ok {
            ans++
        }
    }

    return ans
}