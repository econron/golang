package student

type Student struct {
	studentID string
	name string
	age string
}

func NewStudent(name string, age string) *Student {
	return &Student{
		name: name,
		age: age,
	}
}