package courses

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-martini/martini"
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
