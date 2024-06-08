func merge(nums1 []int, m int, nums2 []int, n int)  {
    /*
    for i, _ := range nums1 {
        for j:=0; j<n; j++ {
            if(i < m && nums2[j] < nums1[i] ){
                var tmp = nums1[i]
                nums1[i] = nums2[j]
                nums2[j] = tmp
            }
            if(i >= m){ 
                if(nums1[i] == 0){
                    nums1[i] = nums2[j]
                    nums2 = nums2[1:]
                    n = len(nums2)
                    j--
                    continue
                } 
                if(nums2[j] < nums1[i]){
                    var tmp = nums1[i]
                    nums1[i] = nums2[j]
                    nums2[j] = tmp
                }
            }
        }
    }
    */

    for i:=m; i < m+n; i++{
        nums1[i] = nums2[i-m] 
    }
    //var val = 0
    for i:=0;i<len(nums1);i++{
        for j:=i+1; j<len(nums1);j++ {
            if(nums1[j] < nums1[i]){
                val := nums1[i]
                nums1[i] = nums1[j]
                nums1[j] = val
            }
        }
    }
}
