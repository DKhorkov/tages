package errors

type FileNotFoundError struct {
	Message string
}

func (e FileNotFoundError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "File not found."
}
