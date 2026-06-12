import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	hashMap := make(map[int]bool)
    arr := make([][]int, 2)

    for _, val := range nums1 {
        hashMap[val] = true
    }

    for _, val := range nums2 {
        if _, ok := hashMap[val]; !ok {
            arr[1] = append(arr[1], val)
        }
        hashMap[val] = false
    }

    for key, val := range hashMap {
        if val {
            arr[0] = append(arr[0], key)
        }
    }

	return arr
}