package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	Courseprice int     `json:"price"`
	Author      *Author `json:"Author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database

var courses []Course

// Middleware, helper - file

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("My First Api in golang")
	r := mux.NewRouter()

	// Seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", Courseprice: 299, Author: &Author{Fullname: "Mayank", Website: "mydev.com"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Mern Stack", Courseprice: 199, Author: &Author{Fullname: "Mayank", Website: "mydev.com"}})

	// Routing

	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getAllcourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":9000", r))

}

// Controller - file

// serve home request

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my First Api </h1>"))
}

func getAllcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("content-Type", "aplication/json")
	json.NewEncoder(w).Encode(courses) //Basically for getting or displaying
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-Type", "application/json")

	// Grab id from request send by the user

	params := mux.Vars(r)

	// Loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	json.NewEncoder(w).Encode("No Course find with given id")
	return //No need but for asurity purpose

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// What if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// What if : {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) //Here we are storing all the data from request into the created variable -> Course

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// Generated unique id convert it into string
	// append the created course var inside the function to Courses fake db

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return //Again not required but for surity purpose

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Firstly grab id from the request

	params := mux.Vars(r)

	// Loop -> match the id -> remove it from the struct Courses -> add the id from the reqest

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return //Again no need but for assurity
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop -> find the id -> Remove it

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
