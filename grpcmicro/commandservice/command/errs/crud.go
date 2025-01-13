package errs

type CRUDError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返す
func (e *CRUDError) Error() string {
	return e.message
}

// コンストラクタ
func NewCRUDError(message string) *CRUDError {
	return &CRUDError{
		message: message,
	}
}