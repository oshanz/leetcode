/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	max := n
	min := 1
	for {
		number := ((max - min) / 2) + min
		switch guess(number) {
		case -1:
			max = number
		case 0:
			return number
		case 1:
			min = number + 1
		}
	}
}
