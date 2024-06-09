func removeElement(nums []int, val int) int {
	var l = len(nums)
	var k = 0
	var lasti = l - 1
	var pstart = 0
	//if len(nums) == 1 && nums[0] == val {
	//   k++
	//	}
	for i := l - 1; i >= 0 && pstart <= i; i-- {
		// fmt.Println("s nums: ", nums, " i: ", i, " k: ", k, " lasti: ", lasti, " pstart: ", pstart)

		// if i == pstart {
		// 	//continue
		// }

		if nums[i] == val { // need to remove -> should be at tail
			if nums[lasti] != val { //follower is worng -> swap
				tmp := nums[i]
				nums[i] = nums[lasti]
				nums[lasti] = tmp
			} //else fine tail good
			lasti--
			k++
			if nums[pstart] != val { // head good move to next
				pstart++
			}
			continue
		} else { // current ok
			if nums[pstart] != val { // head good move to next
				pstart++
				if nums[lasti] == val { // keep follower move to  next
					lasti--
				}
			} else { // swap head to tail
				nums[pstart] = nums[i]
				nums[i] = val
				pstart++
				k++

				if nums[lasti] != val { //follower is worng -> swap
					tmp := nums[i]
					nums[i] = nums[lasti]
					nums[lasti] = tmp
				}
				lasti--
			}
		}

		// if nums[lasti] != val{
		//     tmp := nums[i]
		//     nums[i] = nums[lasti]
		//     nums[lasti] = tmp
		//     //lasti--
		// }

		// if nums[i] == val{
		//     k++
		//     //lasti--
		//     continue
		// }

		// if nums[i] != val && nums[pstart] == val{
		//    nums[pstart] = nums[i]
		//     nums[i] = val
		//     pstart++
		//     k++
		// }

		// lasti = i
		// fmt.Println("e nums: ", nums, " i: ", i, " k: ", k, " lasti: ", lasti, " pstart: ", pstart)

		// for j:=i+1; j<len(nums); j++{
		//     if nums[j] != val {
		//         nums[i] = nums[j]
		//         nums[j] = val
		//         break
		//     }
		// }
	}
	return l - k
}
