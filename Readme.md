# Sample Code with Documentation
* **Author:** Alan Rondoleanu
* **Created:** 10/10/2024

This is an example *Markdown* formatted Readme file.  It will be formatted correctly on Github and most (All?) git repository websites and also within any good text editor or IDE.

### Licence Details
This program is free software: you can redistribute it and/or modify  
it under the terms of the GNU General Public License as published by  
the Free Software Foundation, either version 3 of the License, or (at  
your option) any later version.
 
This program is distributed in the hope that it will be useful, but  
WITHOUT ANY WARRANTY; without even the implied warranty of  
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU  
General Public License for more details.  
You should have received a copy of the GNU General Public License  
along with GNU Emacs.  If not, see the [GNU licence web page](http://www.gnu.org/licenses/).  

##  Description:
This program uses two functions.
One to translate roman numerals in string format to this numerical value.
The other removes duplicates from an array of numbers.

### Example 1:
Input: "MMMDCCXXIV"
// Output: 3726
###  Example 2:
Input: nums = [1,2,2,4,4,9]
Output: [1,2,4,9]
###  Constraints Roman:
1. 1 <= s.length <= 15
2. contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
3. roman numeral in the range [1, 3999].
###  Constraints Remove Duplicates:
1. 1 <= nums.length <= 810,000
2. -100 <= nums[i] <= 100
3. nums is sorted in non-decreasing order.

## Project Files
1. **GoAttempt.go** source code for solution
2. **Readme.md** this Readme file
3. **Go_test.go** code containing Unit Tests

## Running and Installation Code
To run the code do:
> go run GoAttempt.go

To run the tests do:
> go test 

or to see more output from the tests:
> go test -v

To install first create an executable with:
> go build GoAttempt.go

Then copy the executable (**GoAttempt** or **GoAttempt.exe**) into a directory listed in your **PATHS**

