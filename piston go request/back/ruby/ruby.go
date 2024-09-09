package ruby

import (
	// "encoding/json"
	// "fmt"
	//  "strings"
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
// function_code:=`
//   def self.two_sum(nums, target)
//     num_map = {}
//     nums.each_with_index do |num, i|
//       complement = target - num
//       return [num_map[complement], i] if num_map.key?(complement)
//       num_map[num] = i
//     end
//     []
//   end
//   `
// s:=rubyCode(jsonStr,jsonStr,function_code)
//   fmt.Println(s)

// }

  func processJSON (jsonData string,in int ) string {



name:="nums"
if in == 1 {
  name="targets"
}

	s:= name+" = "+jsonData
  return s;
  }



func testcaseCode(json1 string,json2 string)string{
  s:=processJSON(json1,0)+`
  `+processJSON(json2,1);
  return s;
}

func RubyCode(json1 string,json2 string,function_code string)string{
  header:=`
  require 'set'
require 'enumerator'
require 'bigdecimal'

`
classCode:=`class TwoSum
`
timeComplexity:=`

def self.calculate_complexity(nums, target)
    start_time = Time.now

    result = two_sum(nums, target)

    end_time = Time.now
    time_complexity = (end_time - start_time) * 1000.0 # convert to milliseconds

    # Space complexity (in bytes)
    space_complexity = (nums.size * [0].pack("i").size) + (nums.size * [0].pack("i").size) + (nums.size * 32) # Approximate size of Hash

    [time_complexity, space_complexity]
  end
end

`

mainCode:=`
results = []
complexities = []

nums.each_with_index do |num, i|
  target = targets[i]
  complexity = TwoSum.calculate_complexity(num, target)
  complexities << complexity

  result = TwoSum.two_sum(num, target)
  results << result

  puts "Test case #{i + 1}:"
  print "Input: num = ["
  num.each_with_index do |n, j|
    print n
    print "," if j < num.length - 1
  end
  puts "], target = #{target}"

  print "Output: ["
  result.each_with_index do |r, j|
    print r
    print "," if j < result.length - 1
  end
  puts "]"

  puts "Time complexity: #{complexity[0]} ms"
  puts "Space complexity: #{complexity[1]} bytes\n\n"
end
`
code:=header+classCode+function_code+timeComplexity+testcaseCode(json1,json2)+mainCode
return code;
}