func maxNumberOfBalloons(text string) int {
    var arr[15] int
    cnt := 0
    
    for _, val := range text {
        switch val {
            case 'b', 'a', 'l', 'o', 'n':
                arr[val - 'a']++
        }
    }

    for arr[0] > 0 && arr[1] > 0 && arr[11] > 0 && arr[13] > 0 && arr[14] > 0 {
        arr[0]--
        arr[1]--
        arr[13]--

        // 11, 14 l, o
        if arr[11] >= 2 && arr[14] >= 2 {
            arr[11] -= 2 
            arr[14] -= 2 
        } else {
            break
        }
        cnt++
    }

    return cnt
}