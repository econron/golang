package lesson

type Lesson struct {
	slot *Slot
	studentID string
}

func NewLesson(slot *Slot, studentID string) *Lesson {
	return &Lesson {
		slot: slot,
		studentID: studentID,
	}
}