package errs

type IntenalError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返す
func (e *IntenalError) Error() string {
	return e.message
}

// コンストラクタ
func NewIntenalError(message string) *IntenalError {
	return &IntenalError{
		message: message,
	}
}