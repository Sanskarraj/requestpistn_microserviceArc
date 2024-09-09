package rust

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func formatToVecString(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		var buffer bytes.Buffer
		buffer.WriteString("vec![")
		for i, item := range v {
			buffer.WriteString(formatToVecString(item))
			if i < len(v)-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString("]")
		return buffer.String()
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("\"%v\"", v)
	default:
		return ""
	}
}

// func main() {
// 	jsonData1 := `[
// 		[2, 7, 11, 15],
// 		[3, 2, 4],
// 		[1, 5, 5],
// 		[1, 2, 3, 4, 5],
// 		[5, 3, 5, 7]
// 	]`

// 	jsonData2 := `
// 	[
// 	 [
//         "abc",
//         "def",
//         "ghi"
//     ],
// 	[
//         "abczxc",
//         "defghij",
//         "def",
//         "pqrstuv"
//     ]
// 		]
// 	`
// function_code:=`fn two_sum(nums: &[i32], target: i32) -> Vec<i32> {
//     let mut num_map = HashMap::new();
//     for (i, &num) in nums.iter().enumerate() {
//         let complement = target - num;
//         if let Some(&index) = num_map.get(&complement) {
//             return vec![index as i32, i as i32];
//         }
//         num_map.insert(num, i);
//     }
//     vec![]
// }`


// // fmt.Println(testcaseCode(jsonData1,jsonData2))

// code:=rustCode(jsonData1,jsonData2,function_code)

// 	fmt.Println(code)

// }



func processJSON(jsonData string, in int )string {
	var data interface{}
	json.Unmarshal([]byte(jsonData), &data)
	result := formatToVecString(data)

	name:="nums"
	if in==1{
		name= "targets"
	}

	s:= "let "+name+" = "+result+";"
	// fmt.Printf(s)
return s;

}


func testcaseCode(json1 string,json2 string)string{
	q:= processJSON(json1,0)+processJSON(json2,1)
	return q;
}


func RustCode(json1 string, json2 string, function_code string) string {
	header:=
`use std::collections::*;     
use std::cmp::*;            
use std::fmt::*;            
use std::str::*;            
use std::iter::*;
`
timeComplexity:=`
fn calculate_complexity(nums: &[i32], target: i32) -> (f64, usize) {
    let start_time = Instant::now();
    
    let result = two_sum(nums, target);
    
    let duration = start_time.elapsed();
    let time_complexity = duration.as_secs_f64() * 1000.0; // convert to milliseconds
    
    // Space complexity (in bytes)
    let space_complexity = std::mem::size_of_val(nums) + nums.len() * 32; // Approximate size of HashMap
    
    (time_complexity, space_complexity)
}
`

mainClass:=`
fn main() 
{
`
mainCode:=`

let mut results = Vec::new();
    let mut complexities = Vec::new();

    for (i, num) in nums.iter().enumerate() {
        let target = targets[i];
        let complexity = calculate_complexity(num, target);
        complexities.push(complexity);
        
        let result = two_sum(num, target);
        results.push(result);
        
        println!("Test case {}:", i + 1);
        print!("Input: num = [");
        for (j, &n) in num.iter().enumerate() {
            print!("{}", n);
            if j < num.len() - 1 {
                print!(",");
            }
        }
        println!("], target = {}", target);
        
        print!("Output: [");
        for (j, &r) in results[i].iter().enumerate() {
            print!("{}", r);
            if j < results[i].len() - 1 {
                print!(",");
            }
        }
        println!("]");
        
        println!("Time complexity: {:.6} ms", complexities[i].0);
        println!("Space complexity: {} bytes\n", complexities[i].1);
    }
}`


code := header  + function_code + timeComplexity + mainClass + testcaseCode(json1,json2) + mainCode ;
return code;


}