func uniqueXorTriplets(nums []int) int {
    n := len(nums)

    if n < 3 {
        return n
    }

    bitLength := bits.Len(uint(n))

    return 1 << bitLength
}