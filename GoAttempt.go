// Author: Alan Rondoleanu
// Created: 10/10/2024
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at
// your option) any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
// You should have received a copy of the GNU General Public License
// along with GNU Emacs.  If not, see <http://www.gnu.org/licenses/>.
//
// ----------------------------------------------
// Description:
// This program uses two functions.
// One to translate roman numerals in string format to this numerical value.
// The other removes duplicates from an array of numbers.
// Example 1:
// Input: "MMMDCCXXIV"
// Output: 3726
// Example 2:
// Input: nums = [1,2,2,4,4,9]
// Output: [1,2,4,9]
// Constraints Roman:
// 1 <= s.length <= 15
// contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// roman numeral in the range [1, 3999].
// Constraints Remove Duplicates:
// 1 <= nums.length <= 810,000
// -100 <= nums[i] <= 100
// nums is sorted in non-decreasing order.
// ----------------------------------------------

package main

import (
	"fmt"
)

func Roman(t_year string) int {

	var year = t_year

	// User Validation
	if len(year) >= 15 && len(year) <= 0 {
		fmt.Println("1 - 15 Characters only")
	}

	var pass = true
	if pass == true {
		for i := 0; i <= len(year)-1; i++ {
			if year[i] != 'I' &&
				year[i] != 'V' &&
				year[i] != 'X' &&
				year[i] != 'L' &&
				year[i] != 'C' &&
				year[i] != 'D' &&
				year[i] != 'M' {

				fmt.Println("Roman Numerals Only.")
				pass = false
				break
			}
		}
	}

	var count = 0

	// Converts numerals to integers
	for i := 0; i <= len(year)-1; i++ {
		switch string(year)[i] {
		case 'M':
			count += 1000
			break
		case 'D':
			count += 500
			break
		case 'C':
			count += 100
			break
		case 'L':
			count += 50
			break
		case 'X':
			count += 10
			break
		case 'V':
			count += 5
			break
		case 'I':
			count += 1
			break
		default:
			count += 0
		}
	}

	if count < 4000 {
		return count
	} else {
		fmt.Println("Number is too high, requires too much processing power.")
		return -1
	}
}

func removeDuplicates(t_array []int) []int {
	var array = t_array

	// Constraint
	if len(array) >= 1 && len(array) <= 810000 {

		// While loop created to keep checking array till all duplicates are gone
		var pass = false
		for !pass {
			pass = true

			// Constraints
			for i := 0; i < len(array); i++ {
				var next = i + 1
				if next == len(array) {
					next = len(array) - 1
				}
				if array[i] > array[next] {
					fmt.Println("Array of numbers must be in decreasing order")
					return []int{}
				}
				if array[i] > 100 || array[i] < -100 {
					fmt.Println("Numbers in array must be between -100 and 100")
					return []int{}
				}
			}

			// Nested for statement to check array against itself for duplicates
			for i := 0; i < len(array); i++ {
				for o := 0; o < len(array); o++ {
					if array[i] == array[o] && i != o {
						array = append(array[:o], array[o+1:]...)
						pass = false
					}
				}
			}
		}
		return array
	} else {
		return []int{}
	}
}

func main() {

	var array = removeDuplicates([]int{1, 2, 3, 3, 5, 7, 8, 9})
	var test = removeDuplicates(array)
	fmt.Println(test)
}
