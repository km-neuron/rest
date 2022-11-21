package client

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/rest/model"
)

func Student(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("template", "student.html")

	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var client = &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/student", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var student []model.Student
	err = json.Unmarshal(body, &student)
	if err != nil {
		panic(err)
	}

	var data = map[string]interface{}{
		"title": "Student Bootcamp",
		"data":  student,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
