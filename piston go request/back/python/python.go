package python

import (
	// "fmt"
)




// func main() {
// 	jsonStr := `[
//   [
//         [2, 7, 11, 15],
//         [3, 2, 4],
//         [1, 5, 5],
//         [1, 2, 3, 4, 5],
//         [5, 3, 5, 7]
//     ],
//     [
//     [1, 5, 5],
//         [1, 2, 3, 4, 5],
//         [5, 3, 5, 7]
//     ]
//   ]
//     `
// function_code:=`def two_sum(nums, target):
//     num_map = {}
//     for i, num in enumerate(nums):
//         complement = target - num
//         if complement in num_map:
//             return [num_map[complement], i]
//         num_map[num] = i
//     return []
// `
// s:=pythonCode(jsonStr,jsonStr,function_code)
//   fmt.Println(s)

// }

  func processJSON (jsonData string,in int ) string {

name:="nums_list"
if in == 1 {
  name="targets"
}

	s:= name+" = "+jsonData
  return s;
  }



func testcaseCode(json1 string,json2 string)string{
  s:=`      
  `+processJSON(json1,0)+`
  `+processJSON(json2,1);
  return s;
}

func PythonCode(json1 string,json2 string,function_code string)string{
  header:=`
import math      
import heapq    
import bisect   
import time
from collections import defaultdict, deque, Counter  
from itertools import permutations, combinations, product 

`
timeComplexity:=`
def calculate_complexity(nums, target):
    start_time = time.time()
    
    result = two_sum(nums, target)
    
    end_time = time.time()
    time_complexity = (end_time - start_time) * 1000  # Convert to milliseconds
    
    # Space complexity (in bytes)
    space_complexity = sys.getsizeof(nums) + sys.getsizeof({}) + len(nums) * sys.getsizeof(int())  # Approximate size of list and dict
    
    return time_complexity, space_complexity

`
mainClass:=`
def main():
`
mainCode:=`

    for i, nums in enumerate(nums_list):
        target = targets[i]
        complexity = calculate_complexity(nums, target)
        
        result = two_sum(nums, target)
        
        print(f"Test case {i + 1}:")
        print(f"Input: num = [{', '.join(map(str, nums))}], target = {target}")
        
        print(f"Output: [{', '.join(map(str, result))}]")
        
        print(f"Time complexity: {complexity[0]:.5f} ms")
        print(f"Space complexity: {complexity[1]} bytes\n")

if __name__ == "__main__":
    main()

`
code:=header+function_code+timeComplexity+mainClass+testcaseCode(json1,json2)+mainCode
return code;
}