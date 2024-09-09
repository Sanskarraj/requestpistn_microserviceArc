package swift

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
// static func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
//         var numMap = [Int: Int]()
//         for i in 0..<nums.count {
//             let complement = target - nums[i]
//             if let complementIndex = numMap[complement] {
//                 return [complementIndex, i]
//             }
//             numMap[nums[i]] = i
//         }
//         return []
//     }

//   `
// s:=swiftCode(jsonStr,jsonStr,function_code)
//   fmt.Println(s)

// }

  func processJSON (jsonData string,in int ) string {



name:="nums"
if in == 1 {
  name="targets"
}

	s:="let "+ name+" = "+jsonData
  return s;
  }



func testcaseCode(json1 string,json2 string)string{
  s:=processJSON(json1,0)+`
  `+processJSON(json2,1);
  return s;
}

func SwiftCode(json1 string,json2 string,function_code string)string{
header:=`
import Foundation

`
classCode:=`
class TwoSum {

`

timeComplexity:=`

static func calculateComplexity(_ nums: [Int], _ target: Int) -> (Double, Int) {
        let startTime = DispatchTime.now()
        
        let _ = twoSum(nums, target)
        
        let endTime = DispatchTime.now()
        let timeComplexity = Double(endTime.uptimeNanoseconds - startTime.uptimeNanoseconds) / 1_000_000 // convert to milliseconds
        
        // Space complexity (in bytes)
        let spaceComplexity = MemoryLayout<Int>.size * nums.count * 2 + nums.count * 32 // Approximate size of Dictionary
        
        return (timeComplexity, spaceComplexity)
    }

`

mainClass:=`
    static func main() {
`

mainCode:=`

var results = [[Int]]()
        var complexities = [(Double, Int)]()
        
        for i in 0..<nums.count {
            let target = targets[i]
            let complexity = calculateComplexity(nums[i], target)
            complexities.append(complexity)
            
            let result = twoSum(nums[i], target)
            results.append(result)
            
            print("Test case \(i + 1):")
            print("Input: num = [", terminator: "")
            for j in 0..<nums[i].count {
                print(nums[i][j], terminator: j < nums[i].count - 1 ? "," : "")
            }
            print("], target = \(target)")
            
            print("Output: [", terminator: "")
            for j in 0..<result.count {
                print(result[j], terminator: j < result.count - 1 ? "," : "")
            }
            print("]")
            
            print("Time complexity: \(complexity.0) ms")
            print("Space complexity: \(complexity.1) bytes\n")
        }
    }
}

// Run the main function
TwoSum.main()

`


code:=header+classCode+function_code+timeComplexity+mainClass+testcaseCode(json1,json2)+mainCode
return code;
}