package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sb3/level5/modules"
	"testing"
	"time"
)

func TestGetDepartment(t *testing.T) {
	srv := httptest.NewServer(handler(getDepartmentHandler))
	defer srv.Close()

	response, err := http.Get(fmt.Sprintf("%s/department/1", srv.URL))
	if err != nil {
		t.Fatalf("gagal routering : %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("response status tidak sama 200 OK : %v", response.StatusCode)
	}
	var data modules.Department
	b, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("gagal baca body : %v", err)
	}
	defer response.Body.Close()

	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("kesalahan pada convert : %v", err)
	}
	fmt.Println(data)
	if data.ID != 1 {
		t.Fatalf("tidak sama nilainya : %v", data.ID)
	}
}

func TestPostDepartment(t *testing.T) {
	srv := httptest.NewServer(handler(postDepartmentHandler))
	defer srv.Close()
	rand.Seed(time.Now().UnixNano())

	randName := "dummy" + RandStringBytes(3)
	values := modules.Department{Name: randName}
	jsonValue, _ := json.Marshal(values)
	response, err := http.Post(fmt.Sprintf("%s/department", srv.URL), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("gagal routering : %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("response status tidak sama 200 OK : %v", response.StatusCode)
	}

	var data modules.Department
	b, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("gagal baca body : %v", err)
	}
	defer response.Body.Close()

	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("kesalahan pada convert : %v", err)
	}

	if data.Name != randName {
		t.Fatalf("tidak sama nilainya : %v", data.ID)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
