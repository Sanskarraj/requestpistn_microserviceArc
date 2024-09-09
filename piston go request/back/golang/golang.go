package golang

import (
	"encoding/json"
	"fmt"
	"strings"
)



func processValue(v interface{}) string {
	switch val := v.(type) {
	case []interface{}:
		items := make([]string, len(val))
		for i, item := range val {
			items[i] = processValue(item)
		}
		return fmt.Sprintf("{%s}", strings.Join(items, ", "))
	case map[string]interface{}:
		// Handle objects (not needed in this case, but included for completeness)
		items := make([]string, 0, len(val))
		for _, v := range val {
			items = append(items, processValue(v))
		}
		return fmt.Sprintf("{%s}", strings.Join(items, ", "))
	case float64:
		return fmt.Sprintf("%d", int(val))
	case string:
		return fmt.Sprintf("%q", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}


func GetBaseType(input interface{}) string {
	switch input.(type) {
	case []interface{}:
		if len(input.([]interface{})) > 0 {
			return GetBaseType(input.([]interface{})[0])
		}
		return "object"
	case float64:
		return "int"
	case string:
		return "string"
	default:
		return "object"
	}
}

func GetDepth(input interface{}) int {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return 1 + GetDepth(array[0])
	}
	return 0
}

func GetNesting(depth int) string {
  return strings.Repeat("[]", depth)
}


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
// function_code:=`func twoSum(nums []int, target int) []int {
//     numMap := make(map[int]int)
//     for i, num := range nums {
//         complement := target - num
//         if idx, found := numMap[complement]; found {
//             return []int{idx, i}
//         }
//         numMap[num] = i
//     }
//     return []int{}
// }`
// s:=GoCode(jsonStr,jsonStr,function_code)
//   fmt.Println(s)

// }

  func processJSON (jsonData string,in int ) string {

    var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println(err)
	}

	result := processValue(data)
  depth:=GetDepth(data)
	
nesting:= GetNesting(depth)

baseType:= GetBaseType(data)

name:="numsList"
if in == 1 {
  name="targets"
}

	s:= name+" := "+nesting+baseType+result
  return s;
  }



func testcaseCode(json1 string,json2 string)string{
  s:=processJSON(json1,0)+`
  `+processJSON(json2,1);
  return s;
}

func GoCode(json1 string,json2 string,function_code string)string{
  header:=`
package main

import (
    "container/heap"
    "container/list"
    "errors"
    "fmt"
    "math"
    "math/rand"
    "sort"
    "strconv"
    "strings"
    "sync"
    "time"
)
`
timeComplexity:=`
func calculateComplexity(nums []int, target int) (float64, int64) {
    start := time.Now()
    
    twoSum(nums, target)
    
    elapsed := time.Since(start)
    timeComplexity := float64(elapsed.Nanoseconds()) / 1e6 // Convert nanoseconds to milliseconds
    
    // Space complexity (in bytes)
    spaceComplexity := int64(len(nums)*8 + len(nums)*8 + 24) // Approximate size of map
    
    return timeComplexity, spaceComplexity
}
`
mainClass:=`
func main() {
`
mainCode:=`
 _ = heap.Interface(nil)       // Minimal usage of container/heap
    _ = list.New()                // Minimal usage of container/list
    _ = errors.New("sample error") // Minimal usage of errors
    _ = math.Pi                   // Minimal usage of math
    _ = rand.Intn(10)             // Minimal usage of math/rand
    _ = sort.Ints([]int{1})       // Minimal usage of sort
    _ = strconv.Itoa(10)          // Minimal usage of strconv
    _ = strings.Contains("abc", "a") // Minimal usage of strings
    _ = sync.Mutex{}              // Minimal usage of sync
    _ = time.Now()                // Minimal usage of time


for i := 0; i < len(numsList); i++ {
        target := targets[i]
        timeComplexity, spaceComplexity := calculateComplexity(numsList[i], target)
        
        result := twoSum(numsList[i], target)
        
        fmt.Printf("Test case %d:\n", i+1)
        fmt.Printf("Input: nums = %v, target = %d\n", numsList[i], target)
        fmt.Printf("Output: %v\n", result)
        fmt.Printf("Time complexity: %.6f ms\n", timeComplexity)
        fmt.Printf("Space complexity: %d bytes\n\n", spaceComplexity)
    }
`
code:=header+function_code+timeComplexity+mainClass+testcaseCode(json1,json2)+mainCode
return code;
}