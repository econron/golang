package handler

import (
	"fmt"
	"net/http"
	"cleanarch/domain/student"
	"cleanarch/repo/mock"
)

// POST /student/register
func StudentRegisterHandler(req http.Request) http.Response {
	// TODO validate values
	// TODO リクエストからname, ageを取り出す
	name := "John Doe"
	age := "20"
	// 構造体に詰め
	s := student.NewStudent(name,age)
	// dbに入れる
	m := mock.NewSrepo()
	err := m.Save(s)

	if err != nil {
		return http.Response{
			StatusCode: 500,
			Status: fmt.Sprintf("#%v", err),
		}
	}

	return http.Response{
		StatusCode: 200,
	}
	
}