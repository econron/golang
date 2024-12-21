package mock

import (
	"cleanarch/domain/repo"
	"cleanarch/domain/student"
)

type SrepoMock struct {

}

func NewSrepo() repo.Srepo {
	return &SrepoMock{}
}

func (r *SrepoMock) Save(s *student.Student) error {
	return nil
}

/*
型を定義する
インターフェース型を返すけど定義した型を返すようにする
定義した型にインターフェースで定義したメソッドを配置する
*/