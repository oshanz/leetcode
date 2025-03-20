func climbStairs(n int) int {

	// 0 0
	// 1 1
	// 2 11,2
	// 3 111,21,12 ? 2+1,1+2,1+11
	// 4 1111,112,121,211,22 } 3+1 2+2/11 1+21/12/111 {3,2}
	// 5 4,3
	// ways n = ways n-1 + ways n-2
	// 1 = 1
	// 2 = 1 + 0 (2)
	// 3 = 1 + 2
	// 4 =  3 + 2
	// 5 = 5 + 3

	nZero := 1
	nOne := 1
	//nTwo := 2
	oneStepBackWays := nOne
	twoStepBackWays := nZero
	i := 2
	for i <= n {
		currentWays := oneStepBackWays + twoStepBackWays
		twoStepBackWays = oneStepBackWays
		oneStepBackWays = currentWays
        i++
	}
	return oneStepBackWays
}
