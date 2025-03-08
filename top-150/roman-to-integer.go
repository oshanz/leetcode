
/*
LESSON LEARNED

Translating directly from domain logic to code will be the first step. 
Identifying the logic in a creative way will be helpful to build an efficient algorithm

*/


func romanToInt(s string) int {
	mapr := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var out = 0
	for i := 0; i < len(s)-1; i++ {

		char := string(s[i])
		next_char := string(s[i+1])

		if char == "I" && (next_char == "V" || next_char == "X") {
			out--
		} else if char == "X" && (next_char == "L" || next_char == "C") {
			out -= 10
		} else if char == "C" && (next_char == "D" || next_char == "M") {
			out -= 100
		} else {
			out += mapr[char]
		}
	}
	out += mapr[string(s[len(s)-1])]
	return out
}
