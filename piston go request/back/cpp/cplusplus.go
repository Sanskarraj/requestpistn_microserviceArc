
package cpp

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// "log"
	"strings"
)

// func main(){
//   json1:=`
//   [
//         "abc",
//         "def",
//         "ghi"
//   ]
// `


// func_code:=`
// vector<int> twoSum(vector<int>& nums, int target) {
//     unordered_map<int, int> numMap;
//     for (int i = 0; i < nums.size(); ++i) {
//         int complement = target - nums[i];
//         if (numMap.find(complement) != numMap.end()) {
//             return {numMap[complement], i};
//         }
//         numMap[nums[i]] = i;
//     }
//     return {};
// }
// `

// fmt.Println(codeCplusCplus(func_code,json1,json1))



// }



func CodeCplusCplus( testcaseJSON1 string,testcaseJSON2 string,func_code string ,) string {


	header:=`
#include <bits/stdc++.h>
#include <chrono>
using namespace std;
`

// replace with func_code from parameter


complexity_func:=
`pair<double, size_t> calculateComplexity(vector<int>& nums, int target) {
    auto start = chrono::high_resolution_clock::now();
    
    vector<int> result = twoSum(nums, target);
    
    auto end = chrono::high_resolution_clock::now();
    chrono::duration<double, std::milli> duration = end - start;
    
    // Time complexity (in milliseconds)
    double timeComplexity = duration.count();
    
    // Space complexity (in bytes)
    size_t spaceComplexity = sizeof(int) * nums.size() + sizeof(unordered_map<int, int>);
    
    return {timeComplexity, spaceComplexity};
}


int main() {`

testcase:=converterCplusplus(testcaseJSON1,0)+ converterCplusplus(testcaseJSON2,1)

main_code:=`
    vector<vector<int>> res;
    vector<pair<double, size_t>> complexities;

    for (size_t i = 0; i < nums.size(); ++i) {
        int target = targets[i];
        auto [timeComplexity, spaceComplexity] = calculateComplexity(nums[i], target);
        complexities.push_back({timeComplexity, spaceComplexity});
        
        vector<int> result = twoSum(nums[i], target);
        res.push_back(result);
        
        cout << "Test case " << i + 1 << ":\n";
        cout << "Input: num = [";
        for (size_t j = 0; j < nums[i].size(); ++j) {
            cout << nums[i][j];
            if (j < nums[i].size() - 1) cout << ",";
        }
        cout << "], target = " << target << "\n";
        
        cout << "Output: [";
        for (size_t j = 0; j < result.size(); ++j) {
            cout << result[j];
            if (j < result.size() - 1) cout << ",";
        }
        cout << "]\n";
        
        cout << "Time complexity: " << timeComplexity << " ms\n";
        cout << "Space complexity: " << spaceComplexity << " bytes\n\n";
    }

    return 0;
}`


code := header + func_code +complexity_func+ testcase+ main_code

return code;

}





// formatToCppVector converts a JSON structure to a C++ vector string representation
func formatToCppVector(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		elements := make([]string, len(v))
		for i, elem := range v {
			elements[i] = formatToCppVector(elem)
		}
		return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case float64:
		if v == float64(int(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%g", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// getVectorType determines the C++ vector type based on the content
func getVectorType(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		if len(v) == 0 {
			return "auto"
		}
		innerType := getVectorType(v[0])
		if strings.HasPrefix(innerType, "std::vector") {
			return fmt.Sprintf("std::vector<%s>", innerType)
		}
		return fmt.Sprintf("std::vector<%s>", innerType)
	case string:
		return "std::string"
	case float64:
		return "int"
	}
	return "auto"
}

// parseAndCleanJSON removes outer curly braces and parses the JSON input
func parseAndCleanJSON(jsonInput string) (interface{}, error) {
	// Remove the outer curly braces and any whitespace
	jsonInput = strings.TrimSpace(jsonInput)
	jsonInput = strings.TrimPrefix(jsonInput, "{")
	jsonInput = strings.TrimSuffix(jsonInput, "}")
	jsonInput = strings.TrimSpace(jsonInput)

	var data interface{}
	if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
		return nil, err
	}
	return data, nil
}

// converterCplusplus converts a JSON string to a C++ vector representation
func converterCplusplus(input string,i int ) string {
	data, err := parseAndCleanJSON(input)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return ""
	}

	name := "nums"
	if(i ==1){name="targets"}

	vectorType := getVectorType(data)
	result := fmt.Sprintf("%s%s%s%s", vectorType,name,formatToCppVector(data),";")
	return result
}
