package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Platform string   `json:"platform"`
	Price    int      `json:"price"`
	Duration int      `json:"duration"`
	Password string   `json:"-"` // - will not send the column
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json part")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	var myCourses = []course{
		{"ReactJS", "Udemy", 599, 222, "223hdhd", []string{"FE", "SPA"}},
		{"ExpressJS", "Youtube", 0, 190, "122jfj", []string{"BE", "FAST", "NODEJS"}},
		{"MERN", "Coursera", 999, 900, "vcvg48db", nil},
		{"GoLang", "Udemy", 499, 100, "0bb44", []string{"CLOUD Lang", "FAST"}},
	}

	byteCode, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byteCode))
}

func DecodeJson() {
	dataFromWeb := []byte(`
	{
		"coursename": "GoLang",
		"platform": "Udemy",
		"price": 499,
		"duration": 100,
		"tags": ["CLOUD Lang","FAST"]
    }
	`)

	isValid := json.Valid(dataFromWeb)

	if isValid {
		fmt.Println("data is valid json")
		var myCourses course
		json.Unmarshal(dataFromWeb, &myCourses)
		fmt.Println(myCourses)
		fmt.Printf("%#v\n", myCourses)
	}
	//2nd way
	var myOnlineCourses map[string]interface{} //will give u as alias name only
	json.Unmarshal(dataFromWeb, &myOnlineCourses)
	fmt.Printf("%#v\n", myOnlineCourses)

	for k, v := range myOnlineCourses {
		fmt.Printf("%v : %v and type of data is %T\n", k, v, v)
	}
}
