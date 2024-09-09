package javascript

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
// function twoSum(nums, target) {
//     const numMap = new Map();
//     for (let i = 0; i < nums.length; i++) {
//         const complement = target - nums[i];
//         if (numMap.has(complement)) {
//             return [numMap.get(complement), i];
//         }
//         numMap.set(nums[i], i);
//     }
//     return [];
// }

//   `
// s:=javascriptCode(jsonStr,jsonStr,function_code)
//   fmt.Println(s)

// }

  func processJSON (jsonData string,in int ) string {



name:="numsList"
if in == 1 {
  name="targets"
}

	s:="const "+ name+" = "+jsonData
  return s;
  }



func testcaseCode(json1 string,json2 string)string{
  s:=processJSON(json1,0)+`
  `+processJSON(json2,1);
  return s;
}

func JavascriptCode(json1 string,json2 string,function_code string)string{


timeComplexity:=`

function calculateComplexity(nums, target) {
    const startTime = performance.now();
    
    // Execute twoSum here for time complexity measurement
    const result = twoSum(nums, target);
    
    const endTime = performance.now();
    const timeComplexity = endTime - startTime; // time in milliseconds
    
    // Approximate space complexity in bytes
    const spaceComplexity = (4 * nums.length) + (4 * nums.length) + (nums.length * 32); // Approximate size of Map
    
    // Return an object containing time complexity, space complexity, and result
    return {
        timeComplexity: timeComplexity,
        spaceComplexity: spaceComplexity,
        result: result
    };
}

`

mainCode:=`

function printResults(numsList, targets) {
    numsList.forEach(function(nums, index) {
        var target = targets[index];
        
        // Get the result and complexity in one step
        var resultData = calculateComplexity(nums, target);
        var timeComplexity = resultData.timeComplexity;
        var spaceComplexity = resultData.spaceComplexity;
        var result = resultData.result;
        
        console.log('Test case ' + (index + 1) + ':');
        console.log('Input: num = [' + nums.join(',') + '], target = ' + target);
        console.log('Output: [' + result.join(',') + ']');
        console.log('Time complexity: ' + timeComplexity.toFixed(6) + ' ms');
        console.log('Space complexity: ' + spaceComplexity + ' bytes\n');
    });
}

`

utility:=`

printResults(numsList, targets);
`


code:=function_code+timeComplexity+mainCode+testcaseCode(json1,json2)+utility
return code;
}