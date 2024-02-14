package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sb3/level4/modules"
	"strconv"
	"strings"
)

func main() {
	http.ListenAndServe(":8080", handler())
}

func handler() http.Handler {
	hd := http.NewServeMux()
	hd.HandleFunc("/department/{id}", DepartmentHandler)
	hd.HandleFunc("/department", DepartmentHandler)
	return hd
}

func DepartmentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDepartmentHandler(w, r)
	case http.MethodPost:
		postDepartmentHandler(w, r)
	case http.MethodPut:
		putDepartmentHandler(w, r)
	case http.MethodDelete:
		delDepartmentHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/department/")
	if id != "" {
		vid, _ := strconv.ParseInt(id, 10, 64)
		s, err := modules.GetDepartmentByID(vid)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		}
		data, err := json.Marshal(s)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("kesalahan convert json : %v", err)))
		}
		w.Write([]byte(string(data)))
	}
}

func postDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	var data modules.Department
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan pada body : %v", err)))
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan pada convert : %v", err)))
	}

	s, err := modules.PostDepartment(data.Name)
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}
	dp := modules.Department{ID: s.ID, Name: s.Name}
	rt, err := json.Marshal(dp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan pada convert : %v", err)))
	}
	w.Write([]byte(string(rt)))
}

func putDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	var data modules.Department
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan pada body : %v", err)))
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan pada convert : %v", err)))
	}
	s, err := modules.PutDepartment(data.ID, data.Name)
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}
	dp := modules.Department{ID: s.ID, Name: s.Name}
	rt, err := json.Marshal(dp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan pada convert : %v", err)))
	}
	w.Write([]byte(string(rt)))
}

func delDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/department/")
	if id != "" {
		vid, _ := strconv.ParseInt(id, 10, 64)
		err := modules.DelDepartment(vid)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		}
	}
	w.Write([]byte(fmt.Sprint("hapus data berhasil")))
}
