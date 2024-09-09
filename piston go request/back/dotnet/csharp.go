package dotnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ConvertJSONToCSharpList converts nested JSON arrays into a formatted C# string.
func ConvertJSONToCSharpList(input interface{}, depth int) string {
	var buffer bytes.Buffer

	switch value := input.(type) {
	case []interface{}:
		if len(value) == 0 {
			return "new List<object>()"
		}

		baseType := GetBaseType(value[0])
		switch depth {
		case 1:
			buffer.WriteString(fmt.Sprintf("new List<%s> {", baseType))
		case 2:
			buffer.WriteString(fmt.Sprintf("%snew List<List<%s>> {\n", strings.Repeat("    ", depth-1), baseType))
		case 3:
			buffer.WriteString(fmt.Sprintf("%snew List<List<List<%s>>> {\n", strings.Repeat("    ", depth-1), baseType))
		case 4:
			buffer.WriteString(fmt.Sprintf("%snew List<List<List<List<%s>>>> {\n", strings.Repeat("    ", depth-1), baseType))
		case 5:
			buffer.WriteString(fmt.Sprintf("%snew List<List<List<List<List<%s>>>>> {\n", strings.Repeat("    ", depth-1), baseType))
		default:
			buffer.WriteString(fmt.Sprintf("%snew List<object> {\n", strings.Repeat("    ", depth-1)))
		}

		for i, v := range value {
			if depth > 1 {
				buffer.WriteString(strings.Repeat("    ", depth))
			}
			buffer.WriteString(ConvertJSONToCSharpList(v, depth-1))
			if i < len(value)-1 {
				buffer.WriteString(",")
				if depth > 1 {
					buffer.WriteString("\n")
				} else {
					buffer.WriteString(" ")
				}
			}
		}

		if depth > 1 {
			buffer.WriteString("\n" + strings.Repeat("    ", depth-1))
		}
		buffer.WriteString("}")

	case float64:
		buffer.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	case string:
		buffer.WriteString(fmt.Sprintf("\"%s\"", value))
	}
	return buffer.String()
}

// GetBaseType determines the base type of the nested array (int or string).
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

// GenerateCSharpType generates the appropriate C# type based on the depth and base type of the nested array.
func GenerateCSharpType(depth int, baseType string) string {
	if depth == 0 {
		return baseType
	}
	return fmt.Sprintf("List<%s>", GenerateCSharpType(depth-1, baseType))
}

// GetDepth determines the depth of the nested array.
func GetDepth(input interface{}) int {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return 1 + GetDepth(array[0])
	}
	return 0
}

func processJSON(jsonData string, in int) string {
	var jsonParsed interface{}

	err := json.Unmarshal([]byte(jsonData), &jsonParsed)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err.Error()
	}

	name := "nums"
	if in == 1 {
		name = "targets"
	}

	depth := GetDepth(jsonParsed)
	baseType := GetBaseType(jsonParsed)

	csharpType := GenerateCSharpType(depth, baseType)

	converted := ConvertJSONToCSharpList(jsonParsed, depth)

	fmt.Printf("%s %s = %s;\n", csharpType, name, converted)


	s := csharpType + " " + name + " = " + converted + ";"
	return s
}


func testcaseCode(json1 string,json2 string) string {
	q:= processJSON(json1,0)+processJSON(json2,1)
	return q;
}


func CsharpCode(json1 string, json2 string, function_code string) string {
	header:=`package dotnet;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Math;
using System.Text;
using System.Collections;
using System.Collections.Concurrent;
`
class:=`public class TwoSum
{
`
timeComplexity:=`public static Tuple<double, long> CalculateComplexity(List<int> nums, int target)
    {
        var startTime = DateTime.Now;
        
        List<int> result = FindTwoSum(nums, target);
        
        var endTime = DateTime.Now;
        double timeComplexity = (endTime - startTime).TotalMilliseconds;
        
        // Space complexity (in bytes)
        long spaceComplexity = (sizeof(int) * nums.Count) + (sizeof(int) * numMap.Count); // Approximate size of Dictionary
        
        return Tuple.Create(timeComplexity, spaceComplexity);
    }
	`
mainClass:=`public static void Main(string[] args)
    {
`
mainCode:=`
var results = new List<List<int>>();
        var complexities = new List<Tuple<double, long>>();

        for (int i = 0; i < nums.Count; i++)
        {
            int target = targets[i];
            var complexity = CalculateComplexity(nums[i], target);
            complexities.Add(complexity);
            
            List<int> result = FindTwoSum(nums[i], target);
            results.Add(result);
            
            Console.WriteLine("Test case " + (i + 1) + ":");
            Console.Write("Input: num = [");
            for (int j = 0; j < nums[i].Count; j++)
            {
                Console.Write(nums[i][j]);
                if (j < nums[i].Count - 1) Console.Write(",");
            }
            Console.WriteLine("], target = " + target);
            
            Console.Write("Output: [");
            for (int j = 0; j < result.Count; j++)
            {
                Console.Write(result[j]);
                if (j < result.Count - 1) Console.Write(",");
            }
            Console.WriteLine("]");
            
            Console.WriteLine("Time complexity: " + complexity.Item1 + " ms");
            Console.WriteLine("Space complexity: " + complexity.Item2 + " bytes\n");
        }
    }
}
`

// utility:=``

code := header + class + function_code + timeComplexity + mainClass + testcaseCode(json1,json2) + mainCode ;
return code;

}



// func main() {
// 	// Example JSON input
// 	jsonDataInt := `
// 		[
// 			[2, 7, 11, 15],
// 			[3, 2, 4],
// 			[1, 5, 5],
// 			[1, 2, 3, 4, 5],
// 			[5, 3, 5, 7]
// 		]`

// 	jsonDataString := `
// 	[
		
// 			"abc",
// 			"def",
// 			"ghi",
// 			"jklmno",
// 			"pqrstuvwx"
		
//     ]`


// 	function_code:=`public static List<int> FindTwoSum(List<int> nums, int target)
//     {
//         var numMap = new Dictionary<int, int>();
//         for (int i = 0; i < nums.Count; i++)
//         {
//             int complement = target - nums[i];
//             if (numMap.ContainsKey(complement))
//             {
//                 return new List<int> { numMap[complement], i };
//             }
//             numMap[nums[i]] = i;
//         }
//         return new List<int>();
//     }
// 	`

// 	code:=csharpCode(jsonDataInt,jsonDataString,function_code)

// 	fmt.Println(code)

// }
