package java

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ConvertJSONToJavaList converts nested JSON arrays into a formatted Java-like string.
func ConvertJSONToJavaList(input interface{}, depth int) string {
	var buffer bytes.Buffer

	switch value := input.(type) {
	case []interface{}:
		if depth == 1 {
			buffer.WriteString("Arrays.asList(")
			for i, v := range value {
				buffer.WriteString(ConvertJSONToJavaList(v, depth-1))
				if i < len(value)-1 {
					buffer.WriteString(", ")
				}
			}
			buffer.WriteString(")")
		} else {
			buffer.WriteString(strings.Repeat("    ", depth-1) + "Arrays.asList(\n")
			for i, v := range value {
				buffer.WriteString(strings.Repeat("    ", depth))
				buffer.WriteString(ConvertJSONToJavaList(v, depth-1))
				if i < len(value)-1 {
					buffer.WriteString(",\n")
				}
			}
			buffer.WriteString("\n" + strings.Repeat("    ", depth-1) + ")")
		}
	case float64:
		buffer.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	case string:
		buffer.WriteString(fmt.Sprintf("\"%s\"", value))
	}
	return buffer.String()
}

// GetDepth determines the depth of the nested array.
func GetDepth(input interface{}) int {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return 1 + GetDepth(array[0])
	}
	return 0
}

// GetBaseType determines the base type of the nested array (Integer or String).
func GetBaseType(input interface{}) string {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return GetBaseType(array[0])
	}
	switch input.(type) {
	case float64:
		return "Integer"
	case string:
		return "String"
	default:
		return "Object"
	}
}

// GenerateJavaType generates the appropriate Java type based on the depth and base type of the nested array.
func GenerateJavaType(depth int, baseType string) string {
	if depth == 0 {
		return baseType
	}
	return fmt.Sprintf("List<%s>", GenerateJavaType(depth-1, baseType))
}

// func main() {
// 	// Example JSON inputs
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
// 	// Process integer array
// 	//processJSON(jsonDataInt,1)

// 	//fmt.Println("\n---\n")

// 	// Process string array
// 	//rocessJSON(jsonDataString,0)
// 	func_code :=`public static ArrayList<Integer> twoSum(List<Integer> nums, int target) {
//         Map<Integer, Integer> numMap = new HashMap<>();
//         for (int i = 0; i < nums.size(); i++) {
//             int complement = target - nums.get(i);
//             if (numMap.containsKey(complement)) {
//                 return new ArrayList<>(Arrays.asList(numMap.get(complement), i));
//             }
//             numMap.put(nums.get(i), i);
//         }
//         return new ArrayList<>();
//     }`

// 	code:=javaCode(jsonDataInt,jsonDataString,func_code)

// 	fmt.Println(code)

// }

func processJSON(jsonData string, in int) string {
	var jsonParsed interface{}

	// Parse the JSON string into a generic interface{}
	err := json.Unmarshal([]byte(jsonData), &jsonParsed)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err.Error() // Return the error message as a string
	}

	// Determine the name based on the input parameter 'in'
	name := "nums"
	if in == 1 {
		name = "targets"
	}

	// Determine the depth and base type of the parsed JSON structure
	depth := GetDepth(jsonParsed)
	baseType := GetBaseType(jsonParsed)

	// Generate the corresponding Java type (e.g., List<List<Integer>>)
	javaType := GenerateJavaType(depth, baseType)

	// Convert the parsed JSON into a Java-style list declaration
	converted := ConvertJSONToJavaList(jsonParsed, depth)

	// Print the Java variable declaration (optional)
	//fmt.Printf("%s %s = %s;\n", javaType, name, converted)

	// Return the Java variable declaration as a string
	s := javaType + " " + name + " = " + converted + ";"
	return s
}


func testcaseCode(json1 string,json2 string) string {
	q:= processJSON(json1,0)+processJSON(json2,1)
	return q;
}



func JavaCode(json1 string, json2 string, function_code string) string {

	header:=`import java.util.*;               
import java.math.*;               
import java.text.*;               
import java.util.function.*;      
import java.util.regex.*;        
import java.util.stream.*;     
import java.lang.Math;
import java.awt.geom.Point2D;
import java.awt.Point;
import java.awt.geom.Line2D;`


class:=`public class TwoSum {
`
timeComplexity:=` public static Pair<Double, Long> calculateComplexity(List<Integer> nums, int target) {
        long startTime = System.nanoTime();
        
        ArrayList<Integer> result = twoSum(nums, target);
        
        long endTime = System.nanoTime();
        double timeComplexity = (endTime - startTime) / 1e6; // convert to milliseconds
        
        // Space complexity (in bytes)
        long spaceComplexity = (Integer.BYTES * nums.size()) + (Integer.BYTES * nums.size()) + (nums.size() * 32); // Approximate size of HashMap
        
        return new Pair<>(timeComplexity, spaceComplexity);
    }
`

mainClass:=`public static void main(String[] args) {`


mainCode:=`List<ArrayList<Integer>> results = new ArrayList<>();
        List<Pair<Double, Long>> complexities = new ArrayList<>();

        for (int i = 0; i < nums.size(); i++) {
            int target = targets.get(i);
            Pair<Double, Long> complexity = calculateComplexity(nums.get(i), target);
            complexities.add(complexity);
            
            ArrayList<Integer> result = twoSum(nums.get(i), target);
            results.add(result);
            
            System.out.println("Test case " + (i + 1) + ":");
            System.out.print("Input: num = [");
            for (int j = 0; j < nums.get(i).size(); j++) {
                System.out.print(nums.get(i).get(j));
                if (j < nums.get(i).size() - 1) System.out.print(",");
            }
            System.out.println("], target = " + target);
            
            System.out.print("Output: " + result);
            System.out.println();
            
            System.out.println("Time complexity: " + complexity.getKey() + " ms");
            System.out.println("Space complexity: " + complexity.getValue() + " bytes\n");
        }
    }
}`

utility:=`class Pair<K, V> {
    private final K key;
    private final V value;

    public Pair(K key, V value) {
        this.key = key;
        this.value = value;
    }

    public K getKey() { return key; }
    public V getValue() { return value; }
}
`


code := header + class + function_code + timeComplexity + mainClass + testcaseCode(json1,json2) + mainCode + utility

return code;
}