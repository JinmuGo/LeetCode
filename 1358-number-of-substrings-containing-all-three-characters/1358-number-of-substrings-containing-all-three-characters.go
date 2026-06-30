func numberOfSubstrings(s string) int {
    arr := []int{-1, -1, -1}
    ans := 0

    for right, c := range s {
        switch c {
            case 'a':
                arr[0] = right
            case 'b':
                arr[1] = right
            case 'c':
                arr[2] = right
        }

        if arr[0] != -1 && arr[1] != -1 && arr[2] != -1 {
            minIdx := min(arr[0], arr[1], arr[2])
            ans += minIdx + 1
        }
    }

    return ans
}