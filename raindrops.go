// Package raindrops is kind of like fizz buzz
package raindrops

/*


*/

import(
	"strconv"
)

// Convert a number to a string
// The rules of raindrops are that if a given number return a string that
func Convert(num int) string {

	res := ""
	number := strconv.Itoa(num)

	//has 3 as a factor, add 'Pling' to the result.
//has 5 as a factor, add 'Plang' to the result.
//has 7 as a factor, add 'Plong' to the result.
	if num%3 == 0 {
		res += "Pling"
	}
	if num%5 == 0 {
		res += "Plang"
	}
	if num%7 == 0 {
		res += "Plong"
	}

	//does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
	if len(res)==0 {
		res = number
	}

	return res
}
