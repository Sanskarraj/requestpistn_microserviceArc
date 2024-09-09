package dart

import (
	// "fmt"
)



// func main() {
// 	// Example JSON input
// 	jsonData1 := `[
//         [2, 7, 11, 15],
//         [3, 2, 4],
//         [1, 5, 5],
//         [1, 2, 3, 4, 5],
//         [5, 3, 5, 7]
//     ]`

// 	jsonData2 := `
//         [2, 7, 11, 15]`


// // s:=testcaseCode(jsonData1,jsonData2)
// // fmt.Println(s)
// function_code:=`List<int> twoSum(List<int> nums, int target) {
//   Map<int, int> numMap = {};
//   for (int i = 0; i < nums.length; i++) {
//     int complement = target - nums[i];
//     if (numMap.containsKey(complement)) {
//       return [numMap[complement]!, i];
//     }
//     numMap[nums[i]] = i;
//   }
//   return [];
// }`

// s:=dartCode(jsonData1,jsonData2,function_code)
// fmt.Println(s)

// }





func processJSON(jsonData string,in int) string {


	name:="nums"
	if in ==1 {
		name="targets"
	}

	return "var " + name+" = "+jsonData + ";"
}

func testcaseCode(json1 string , json2 string) string {
	s:=processJSON(json1,0)+`

`+processJSON(json2,1)
	return s;
}


func DartCode(json1 string, json2 string, function_code string) string {

	header:=`import 'dart:math';
import 'dart:collection';
import 'dart:convert';
`

timeComplexity:=`

Pair<double, int> calculateComplexity(List<int> nums, int target) {
  int startTime = DateTime.now().microsecondsSinceEpoch;
  List<int> result = twoSum(nums, target);
  int endTime = DateTime.now().microsecondsSinceEpoch;
  double timeComplexity = (endTime - startTime) / 1000; // convert to milliseconds
  // Space complexity (in bytes)
   int spaceComplexity = (4 * nums.length) + (4 * nums.length) + (nums.length * 32); // Approximate size of Map
   return Pair(timeComplexity, spaceComplexity);
}
`
mainClass:=`
void main() {
`

mainCode:=`
List<List<int>> results = [];
  List<Pair<double, int>> complexities = [];

  for (int i = 0; i < nums.length; i++) {
    int target = targets[i];
    Pair<double, int> complexity = calculateComplexity(nums[i], target);
    complexities.add(complexity);
    List<int> result = twoSum(nums[i], target);
    results.add(result);
    print("Test case ${i + 1}:");
    print("Input: num = [${nums[i].join(',')}], target = $target");
    print("Output: [${result.join(',')}]");
    print("Time complexity: ${complexity.key.toStringAsFixed(6)} ms");
    print("Space complexity: ${complexity.value} bytes\n");
  }
}
`

utility:=`
class Pair<K, V> {
  final K key;
  final V value;

  Pair(this.key, this.value);
}
`


code := header+ function_code+timeComplexity+mainClass+testcaseCode(json1,json2)+mainCode+utility
return code
}