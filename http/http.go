package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. 응답에 사용할 구조체 정의 (JSON 태그 포함)
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"Author"`
}

// 2. 핸들러 함수: 특정 경로로 들어왔을 때 실행될 로직
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// r: 요청 정보 (브라우저가 보낸 것)
	// w: 응답 도구 (우리가 브라우저에 보낼 것)
	fmt.Fprint(w, "Hello, Go Web Server!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "홍길동", Email: "gildong@example.com"}

	// 구조체를 JSON 형식으로 변환하여 응답
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {

	mybooks := []Book{
		{ID: 1, Title: "A", Author: "a"},
		{ID: 2, Title: "B", Author: "b"},
		{ID: 3, Title: "C", Author: "c"},
		{ID: 4, Title: "D", Author: "d"},
		{ID: 5, Title: "E", Author: "e"},
	}

	// 3. 경로와 핸들러 함수 연결 (라우팅)
	http.HandleFunc("/", helloHandler)    // http://localhost:8080/
	http.HandleFunc("/user", userHandler) // http://localhost:8080/user
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mybooks) // mybooks 슬라이스를 통째로 JSON으로 변환
	})

	fmt.Println("서버를 8080 포트에서 시작합니다...")

	// 4. 서버 시작 및 에러 확인
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("서버 시작 실패:", err)
	}
}
