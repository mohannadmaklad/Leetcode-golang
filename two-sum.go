func twoSum(nums []int, target int) []int {
    ret := []int{}
    newArray := make([]int, len(nums))
    copy(newArray, nums)
    //Sort the input array
    sort.Ints(newArray)

    for i, n := range newArray{
        cur := n
        needed := target - cur

        // Find the needed number in the rest of the array
        needed_idx := sort.Search(len(newArray), func(idx int) bool {return newArray[idx] >= needed})
        if needed_idx < len(newArray) && newArray[needed_idx] == needed && needed_idx != i{
            found_indices := find(nums,newArray[needed_idx]);
            found_indices = append(found_indices, find(nums,cur)...)
            ret = append(ret, found_indices[0],found_indices[len(found_indices)-1])
            return ret
        }
    }
    return ret;  
}


func find(nums []int, num int) []int{
    ret := []int{}
    for i,n := range nums{
        if n == num {
            ret = append(ret, i)
        }
    }
    return ret
}