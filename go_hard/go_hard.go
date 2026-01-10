package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Searcher interface {
	Search(query string) ([]Book, error)
}

type Library struct {
	Name string
	Data []Book
}

func (l Library) Search(query string) ([]Book, error) {
	time.Sleep(500 * time.Millisecond) // 지연

	var found []Book
	for _, b := range l.Data {
		if b.Title == query {
			found = append(found, b)
		}
	}
	return found, nil
}

var lib1 = Library{Name: "A도서관", Data: []Book{{ID: 1, Title: "Go", Author: "Rob"}}}
var lib2 = Library{Name: "B도서관", Data: []Book{{ID: 2, Title: "Go", Author: "Ken"}}}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "검색어를 입력하세요", http.StatusBadRequest)
		return
	}

	// 1. 결과 전송용 채널 만들기 (결과값인 []Book을 담는 통로)
	// 각각의 고루틴이 찾은 책 리스트를 던져줄 것입니다.
	resultChan := make(chan []Book)

	// 2. 2개의 Library에서 각각 고루틴으로 Search 호출
	go func() {
		books, _ := lib1.Search(query)
		resultChan <- books // 결과 전송
	}()

	go func() {
		books, _ := lib2.Search(query)
		resultChan <- books // 결과 전송
	}()

	// 3. 채널에서 결과 2개를 수집하여 하나로 합치기
	var allResults []Book
	for i := 0; i < 2; i++ {
		books := <-resultChan // 채널에서 데이터가 나올 때까지 대기(Blocking)
		allResults = append(allResults, books...)
	}

	// 4. JSON으로 응답하기
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResults)
}

func main() {
	// 경로 연결
	http.HandleFunc("/search", searchHandler)

	fmt.Println("서버 시작: http://localhost:8080/search?q=Go")

	// 서버 실행
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("서버 종료: %v\n", err)
	}
}
