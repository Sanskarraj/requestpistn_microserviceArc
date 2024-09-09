package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
	"log"
	"encoding/base64"
	// "strconv"
	"back/cpp"
	"back/java"
	"back/python"
	"back/javascript"
	"back/golang"
	"back/rust"
	"back/swift"
	"back/kotlin"
	"back/ruby"
	"back/typescript"
	"back/dotnet"
	"back/dart"
)

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type TestCase struct {
	Input   [][]int `json:"input"`
	Targets []int   `json:"targets"`
}

type ExecuteRequest struct {
	Language           string   `json:"language"`
	Version            string   `json:"version"`
	Files              []File   `json:"files"`
	Stdin              string   `json:"stdin"`
	Args               []string `json:"args"`
	CompileTimeout     int      `json:"compile_timeout"`
	RunTimeout         int      `json:"run_timeout"`
	CompileMemoryLimit int      `json:"compile_memory_limit"`
	RunMemoryLimit     int      `json:"run_memory_limit"`
}


func main() {


	mapLanguage := map[int]string{
		111: "c++",
		121: "java",
		141: "python",
		151: "javascript",
		161: "go",
		173: "rust",
		183: "swift",
		193: "kotlin",
		177: "ruby",
		187: "typescript",
		191: "csharp",
		197: "dart",
	};


	mapVersion := map[int]string{
		111: "10.2.0",
		121: "15.0.2",
		141: "3.12.0",
		151: "20.11.1",
		161: "1.16.2",
		173: "1.68.2",
		183: "5.3.3",
		193: "1.8.20",
		177: "3.0.1",
		187: "5.0.3",
		191: "5.0.201",
		197: "2.19.6",
	};


	mapExtensions := map[int]string{
		111: ".cpp",
		121: ".java",
		141: ".py",
		151: ".js",
		161: ".go",
		173: ".rs",
		183: ".swift",
		193: ".kt",
		177: ".rb",
		187: ".ts",
		191: ".cs",
		197: ".dart",
	};
	

	base64String := "ew0KImlucHV0IjogWw0KCQkgIFsyLCA3LCAxMSwgMTUsIDIwLCAzMCwgNDAsIDUwLCA2MCwgODAsIDkwLCAxMDAsIDEyMCwgMTMwLCAxNDAsIDE1MF0sDQoJCSAgWzMsIDIsIDQsIDgsIDE2LCAyMywgNDIsIDU4LCA3MSwgODQsIDk5LCAxMDEsIDExMywgMTI3LCAxMzUsIDE0OSwgMTYwXSwNCgkJICBbMSwgNSwgNSwgMiwgOSwgMTQsIDI4LCAzMywgNDEsIDU2LCA3MywgODksIDEwMCwgMTA1LCAxMTcsIDEyNCwgMTM3LCAxNDFdDQpdLA0KInRhcmdldHMiOiBbMjIwLCAyMzYsIDIwMF0NCn0="
	func_code:=`
vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> numMap;
    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];
        if (numMap.find(complement) != numMap.end()) {
            return {numMap[complement], i};
        }
        numMap[nums[i]] = i;
    }
    return {};
}
`
	language_id:=111

	// Decode Base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err)
		return
	}

	// Convert decoded bytes to string (JSON)
	originalJSON := string(decodedBytes)
	fmt.Println("Decoded JSON:", originalJSON)


	var testCase TestCase
	err = json.Unmarshal([]byte(originalJSON), &testCase)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Marshal the `input` part back into JSON
	b, err := json.Marshal(testCase.Input)
	if err != nil {
		log.Fatalf("Error marshalling input to JSON: %v", err)
	}

	// Marshal the `targets` part back into JSON
	a, err := json.Marshal(testCase.Targets)
	if err != nil {
		log.Fatalf("Error marshalling targets to JSON: %v", err)
	}
	
    // jsonInput1 := string(Input1)
	// jsonInput2 := string(Input2)
	
fmt.Println("jsonInput1",a)
fmt.Println("jsonInput2",b)
jsonInput1 := `[[2,7,11,15],[3,2,4]]`
    jsonInput2 := `[9,6]`

	language := mapLanguage[language_id]
	version := mapVersion[language_id]
	name :="index"+mapExtensions[language_id]

	code := ""

switch language_id {
	case 111:
		code = cpp.CodeCplusCplus(jsonInput1, jsonInput2, func_code)
	case 121:
		code = java.JavaCode(jsonInput1, jsonInput2, func_code)
	case 141:
		code = python.PythonCode(jsonInput1, jsonInput2, func_code)
	case 151:
		code = javascript.JavascriptCode(jsonInput1, jsonInput2, func_code)
	case 161:
		code = golang.GoCode(jsonInput1, jsonInput2, func_code)
	case 173:
		code = rust.RustCode(jsonInput1, jsonInput2, func_code)
	case 183:
		code = swift.SwiftCode(jsonInput1, jsonInput2, func_code)
	case 193:
		code = kotlin.KotlinCode(jsonInput1, jsonInput2, func_code)
	case 177:
		code = ruby.RubyCode(jsonInput1, jsonInput2, func_code)
	case 187:
		code = typescript.TypescriptCode(jsonInput1, jsonInput2, func_code)
	case 191:
		code = dotnet.CsharpCode(jsonInput1, jsonInput2, func_code)
	case 197:
		code = dart.DartCode(jsonInput1, jsonInput2, func_code)
}

	// Initialize args (if needed)
	args := []string{}

	// Create the JSON request payload
	requestPayload := ExecuteRequest{
		Language:           language,
		Version:            version,
		Files:              []File{{Name: name, Content: code}},
		Stdin:              "",
		Args:               args,
		CompileTimeout:     10000,
		RunTimeout:         3000,
		CompileMemoryLimit: -1,
		RunMemoryLimit:     -1,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON payload
	fmt.Println("Generated JSON:")
	fmt.Println(string(jsonData))

	// Send the POST request
	resp, err := http.Post("http://localhost:2000/api/v2/execute", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	var responseBody bytes.Buffer
	responseBody.ReadFrom(resp.Body)
	fmt.Println("Response from server:")
	fmt.Println(responseBody.String())
}


