package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"authorName"`
	Website    string `json:"website"`
}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

var courses []Course

func main() {
	fmt.Println("Server started running....")

	//seeding data
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 199, Author: &Author{AuthorName: "Sam Ram", Website: "sram.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "GoLang", CoursePrice: 999, Author: &Author{AuthorName: "Abhi", Website: "abhi.go.dev"}})
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>Welcome to courses api</h2>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses....")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course...")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, c := range courses {
		if c.CourseId == params["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("No course is found with course id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating the course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("There is no data in body")
	}
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, val := range courses {
		if val.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CoursePrice = val.CoursePrice
			course.Author = val.Author
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete the course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, c := range courses {
		if c.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode("Course is deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No course found!!")

}
