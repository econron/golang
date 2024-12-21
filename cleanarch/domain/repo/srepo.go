package repo

import (
	"cleanarch/domain/student"
)

type Srepo interface {
	Save(s *student.Student) error
}