func removeDuplicates(nums []int) int {
    var k = 0
    for i :=1; i < len(nums); i++{

        if nums[k] != nums[i] {
            nums[k+1] = nums[i]
            k++

        }
    }
    return k+1
}
