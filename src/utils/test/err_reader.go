package utils_test

type ErrReader struct{ Error error }

func (e *ErrReader) Read([]byte) (int, error) {
	return 0, e.Error
}
