package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sb3/level5/modules"
	"strconv"
	"strings"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	app := fiber.New()
	jwtAuth := JWTAuthMiddleware("sb3-hanif")
	basicAuth := BasicAuthMiddleware()

	app.Post("/login", basicAuth, adaptor.HTTPHandler(handler(genToken)))
	app.Get("/department/:id", basicAuth, adaptor.HTTPHandler(handler(getDepartmentHandler)))
	app.Post("/department", jwtAuth, adaptor.HTTPHandler(handler(postDepartmentHandler)))
	app.Put("/department", jwtAuth, adaptor.HTTPHandler(handler(putDepartmentHandler)))
	app.Delete("/department/:id", jwtAuth, adaptor.HTTPHandler(handler(delDepartmentHandler)))

	app.Listen(":8080")
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func BasicAuthMiddleware() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"hanif": "hanif123",
		},
	})
}

func JWTAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	})
}

func genToken(w http.ResponseWriter, r *http.Request) {
	claims := jwt.MapClaims{
		"name": "testing",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("sb3-hanif"))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("kesalahan generate : %v", err)))
	}
	w.Write([]byte(fmt.Sprintf("token : %v", t)))
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
	w.Write([]byte(fmt.Sprintln("hapus data berhasil")))
}
