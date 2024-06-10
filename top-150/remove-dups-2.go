func removeDuplicates(nums []int) int {
 
    var k = 1
    var l = len(nums)
    for i :=2; i < l; i++{

        if nums[k] != nums[i] {
            nums[k+1] = nums[i]
            k++

        }else if nums[k-1] != nums[k]{
            nums[k+1] = nums[i]
            k++
        }
    }
    if l < 2{
        return l
    } else {
        return k + 1
    }


}
