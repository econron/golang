package errs

type DomainError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返す
func (e *DomainError) Error() string {
	return e.message
}

// コンストラクタ
func NewDomainError(message string) *DomainError {
	return &DomainError{
		message: message,
	}
}