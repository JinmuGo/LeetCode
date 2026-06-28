func maximumElementAfterDecrementingAndRearranging(arr []int) int {
   sort.Ints(arr)
   curNum := 0
   prev := 0

   for i := 0; i < len(arr); i++ {
        elem := arr[i]

        if curNum == elem && elem != prev {
            curNum++
            prev = elem
        }

        if curNum < elem {
            curNum++
            prev = curNum
        }
   }

   return curNum
}