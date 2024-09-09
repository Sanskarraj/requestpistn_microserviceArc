package typescript

import (
    "bytes"
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
)

// ConvertJSONToTypeScript converts JSON data to TypeScript array format.
func ConvertJSONToTypeScript(input interface{}, depth int) string {
    var buffer bytes.Buffer
    switch value := input.(type) {
    case []interface{}:
        if depth == 1 {
            buffer.WriteString("[")
            for i, v := range value {
                buffer.WriteString(ConvertJSONToTypeScript(v, depth-1))
                if i < len(value)-1 {
                    buffer.WriteString(", ")
                }
            }
            buffer.WriteString("]")
        } else {
            buffer.WriteString("[\n")
            for i, v := range value {
                buffer.WriteString(strings.Repeat("  ", depth))
                buffer.WriteString(ConvertJSONToTypeScript(v, depth-1))
                if i < len(value)-1 {
                    buffer.WriteString(",\n")
                }
            }
            buffer.WriteString("\n" + strings.Repeat("  ", depth-1) + "]")
        }
    case float64:
        buffer.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
    case string:
        buffer.WriteString(fmt.Sprintf("\"%s\"", value))
    }
    return buffer.String()
}

func GetNesting(depth int) string {
    return strings.Repeat("[]", depth)
}

// GetDepth determines the depth of the nested array.
func GetDepth(input interface{}) int {
    if array, ok := input.([]interface{}); ok && len(array) > 0 {
        return 1 + GetDepth(array[0])
    }
    return 0
}

// GetType determines the TypeScript type of the input.
func GetType(input interface{}) string {
    switch input.(type) {
    case []interface{}:
        if len(input.([]interface{})) > 0 {
            return GetType(input.([]interface{})[0])
        }
        return "any"
    case float64:
        return "number"
    case string:
        return "string"
    default:
        return "any"
    }
}

// func main() {
//     // Example JSON input
//     jsonData1 := ` [
//         "abc",
//         "def",
//         "ghi"
//     ]`
// 	jsonData2 := ` [
//         "abc",
//         "def",
//         "ghi"
//     ]`

// 	function_code:=`function twoSum(nums: number[], target: number): number[] {
//     const numMap: { [key: number]: number } = {};
//     for (let i = 0; i < nums.length; i++) {
//         const complement: number = target - nums[i];
//         if (complement in numMap) {
//             return [numMap[complement], i];
//         }
//         numMap[nums[i]] = i;
//     }
//     return [];
// }
// `
// s:=typescriptCode(jsonData1,jsonData2,function_code)
// fmt.Println(s)




   
// }

func processJSON(jsonData string, in int) string {
    var jsonParsed interface{}
    err := json.Unmarshal([]byte(jsonData), &jsonParsed)
    if err != nil {
        return fmt.Sprintf("Error parsing JSON: %s", err)
    }

    depth := GetDepth(jsonParsed)
    converted := ConvertJSONToTypeScript(jsonParsed, depth)
    typeStr := GetType(jsonParsed)
    name := "numsList"
    if in == 1 {
        name = "targets"
    }

    return fmt.Sprintf("const %s: %s%s = %s; ", name, typeStr, GetNesting(depth), converted)
}

func testcaseCode(json1 string,json2 string)string{
	q:= processJSON(json1,0)+processJSON(json2,1)
	return q;
}


func TypescriptCode(json1 string, json2 string, function_code string) string {
	// header:=``
	timeComplexity:=`
	function calculateComplexity(nums: number[], target: number): Pair<number, number> {
    const startTime: number = performance.now();
    
    const result: number[] = twoSum(nums, target);
    
    const endTime: number = performance.now();
    const timeComplexity: number = endTime - startTime; // in milliseconds
    
    // Space complexity (in bytes)
    const spaceComplexity: number = (4 * nums.length) + (4 * nums.length) + (nums.length * 32); // Approximate size of the object
    
    return { key: timeComplexity, value: spaceComplexity };
}
`
mainCode:=`
for (let i = 0; i < numsList.length; i++) {
    const nums: number[] = numsList[i];
    const target: number = targets[i];
    
    const complexity: Pair<number, number> = calculateComplexity(nums, target);
    
    const result: number[] = twoSum(nums, target);
    
    console.log("Test case " + (i + 1) + ":");
    console.log("Input: num = [" + nums.join(',') + "]");
    console.log("target = " + target);
    console.log("Output: [" + result.join(',') + "]");
    console.log("Time complexity: " + complexity.key + " ms");
    console.log("Space complexity: " + complexity.value + " bytes\n");
}
`
utility:=`
type Pair<K, V> = {
    key: K;
    value: V;
};
`
code := function_code+timeComplexity+testcaseCode(json1,json2)+mainCode+utility

return code
}