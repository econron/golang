package tutor

type Tutor struct {
	tutorID string
	applicationInfo string
	bankAccountInfo string
	ourUniqueLevel string
	name string
	age string
}

func NewTutor(tutorID string, applicationInfo string, bankAccountInfo string, ourUniqueLevel string, name string, age string) *Tutor {
	return &Tutor{
		tutorID: tutorID,
		applicationInfo: applicationInfo,
		bankAccountInfo: bankAccountInfo,
		ourUniqueLevel: ourUniqueLevel,
		name: name,
		age: age,
	}
}