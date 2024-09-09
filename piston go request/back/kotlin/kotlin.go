package kotlin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("\"%v\"", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func parseArray(array []interface{}) string {
	var buffer bytes.Buffer

	buffer.WriteString("listOf(")
	for i, element := range array {
		switch elem := element.(type) {
		case []interface{}:
			buffer.WriteString(parseArray(elem))
		default:
			buffer.WriteString(formatValue(elem))
		}
		if i != len(array)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func parseJSON(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		return parseArray(v)
	default:
		return formatValue(v)
	}
}

// func main() {
// 	jsonData1 := `
// 	[
//         "abc",
//         "def",
//         "ghi"
//     ]
	
// 	`

// 	jsonData2 := `[
// 	[
// 		[
// 			[2, 7, 11, 15, 20, 30],
// 			[3, 2, 4, 8, 16],
// 			[1, 5, 5, 2, 9],
// 			[3, 2, 4, 8, 16],
// 			[1, 5, 5, 2, 9]
// 		],
// 		[
// 			[2, 7, 11, 15, 20, 30],
// 			[3, 2, 4, 8, 16],
// 			[1, 5, 5, 2, 9]
// 		]

// 	],
// 	[
// 		[	
// 			[1, 5, 5, 2, 9],
// 			[3, 2, 4, 8, 16],
// 			[1, 5, 5, 2, 9]
// 		],
// 		[
// 			[2, 7, 11, 15, 20, 30],
// 			[3, 2, 4, 8, 16],
// 			[1, 5, 5, 2, 9],
// 			[2, 7, 11, 15, 20, 30],
// 			[3, 2, 4, 8, 16]
// 		]
// 	]		
// 	]`

	

// 	function_code:=`
// fun twoSum(nums: List<Int>, target: Int): List<Int> {
//     val numMap = mutableMapOf<Int, Int>()
//     for (i in nums.indices) {
//         val complement = target - nums[i]
//         if (numMap.containsKey(complement)) {
//             return listOf(numMap[complement]!!, i)
//         }
//         numMap[nums[i]] = i
//     }
//     return emptyList()
// }
// `



// 	fmt.Println(kotlinCode(jsonData1,jsonData2,function_code))

// 	// fmt.Println(strings.ReplaceAll(result2, "listOf(", "\nlistOf("))




// }


func combinedAccumulation (jsonInput string, in int ) string {
	var parsedData interface{}

	if err := json.Unmarshal([]byte(jsonInput), &parsedData); err != nil {
		log.Fatal(err)
	}

	name := "nums"
	if in == 1 {
		name= "targets"
	}
	// result := "val nums = " + parseJSON(parsedData)

	result := "val "+name +" = "+ parseJSON(parsedData)


	// Properly format the final output for display
	s:=strings.ReplaceAll(result, "listOf(", "\nlistOf(")
	return s;
}


func testcaseCode(json1 string,json2 string) string {
	q:= combinedAccumulation(json1,0)+`
`+combinedAccumulation(json2,1)
	return q;
}


func KotlinCode(json1 string, json2 string, function_code string) string {
	header:=`import kotlin.math.*  
	import kotlin.collections.*  
	import kotlin.text.*
	import kotlin.ranges.*
	import kotlin.comparisons.* 
	import java.util.*
	import java.math.*;               
	import java.text.*;               
	import java.util.function.*;      
	import java.util.regex.*;        
	import java.util.stream.*;     
	import java.lang.Math;
	import java.awt.geom.Point2D;
	import java.awt.Point;
	import java.awt.geom.Line2D;
	`

    timeComplexity:=`
	fun calculateComplexity(nums: List<Int>, target: Int): Pair<Double, Long> {
    val startTime = System.nanoTime()

    val result = twoSum(nums, target)

    val endTime = System.nanoTime()
    val timeComplexity = (endTime - startTime) / 1e6 // convert to milliseconds

    // Space complexity (in bytes)
    val spaceComplexity: Long = (Integer.BYTES.toLong() * nums.size) + (Integer.BYTES.toLong() * nums.size) + (nums.size.toLong() * 32) // Approximate size of HashMap

    return Pair(timeComplexity, spaceComplexity)
}
`

mainClass:=`fun main() {`

mainCode:=`
val results = mutableListOf<List<Int>>()
val complexities = mutableListOf<Pair<Double, Long>>()

    for (i in nums.indices) {
        val target = targets[i]
        val complexity = calculateComplexity(nums[i], target)
        complexities.add(complexity)

        val result = twoSum(nums[i], target)
        results.add(result)

        println("Test case ${i + 1}:")
        print("Input: num = [")
        for (j in nums[i].indices) {
            print(nums[i][j])
            if (j < nums[i].size - 1) print(",")
        }
        println("], target = $target")

        print("Output: [")
        for (j in result.indices) {
            print(result[j])
            if (j < result.size - 1) print(",")
        }
        println("]")

        println("Time complexity: ${complexity.first} ms")
        println("Space complexity: ${complexity.second} bytes\n")
    }
}
`


code := header + function_code + timeComplexity + mainClass + testcaseCode(json1,json2) + mainCode ;
return code;


}