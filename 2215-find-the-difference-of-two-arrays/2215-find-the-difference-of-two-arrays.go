import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for _, val := range nums1 {
		map1[val] = true
	}

	for _, val := range nums2 {
		map2[val] = true
	}

	for _, val := range nums2 {
		if map1[val] {
			map1[val] = false
		}
	}

	for _, val := range nums1 {
		if map2[val] {
			map2[val] = false
		}
	}

	arr := make([][]int, 2)


	for key, val := range map1 {
		if val {
            arr[0] = append(arr[0], key)
		}
	}

	for key, val := range map2 {
		if val {
            arr[1] = append(arr[1], key)
		}
}

	return arr
}