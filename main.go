package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

type Data struct {
	Data []Response
}
type Response struct {
	ID             int       `json:"id"`
	CourseName     string    `json:"course_name"`
	Description    string    `json:"description"`
	CourseProfName string    `json:"course_prof_name"`
	CreatedDt      time.Time `json:"created_dt"`
}

func GetCourseById(params martini.Params, r render.Render) {
	id := params["id"]
	fmt.Println(string(id))
	url := "http://localhost:8088/courselist"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data Data
	json.Unmarshal(body, &data)
	fmt.Println("data: ", data)
	for i, row := range data.Data {
		fmt.Println(i, row.CourseName, row.CourseProfName)
	}

	r.HTML(200, "courses", data.Data)
}

func GetAllCourse() string {
	return fmt.Sprintf("ALL course list")
}

type StudentData struct {
	Data []StudentResponse
}
type StudentResponse struct {
	StId      int       `json:"st_id"`
	StName    string    `json:"st_name"`
	StEmail   string    `json:"st_email"`
	StPhone   string    `json:"st_phone"`
	CreatedDt time.Time `json:"created_dt"`
}

func GetStudentList(params martini.Params, r render.Render) {
	id := params["id"]
	fmt.Println(string(id))
	url := "http://localhost:8088/studentlist"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data StudentData
	json.Unmarshal(body, &data)
	fmt.Println("data: ", data)
	for i, row := range data.Data {
		fmt.Println(i, row.StId, row.StName)
	}

	r.HTML(200, "students", data.Data)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	// static := martini.StaticOptions{Prefix: "static"}

	// m.Use(martini.Static("static", static))

	m.Group("/courses", func(r martini.Router) {
		r.Get("", GetAllCourse)
		r.Get("/:id", GetCourseById)
	})

	m.Group("/studentlist", func(r martini.Router) {
		r.Get("", GetStudentList)
	})
	m.Run()
}
