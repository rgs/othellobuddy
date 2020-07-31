package board

type BoardError struct {
	Message string
	Where Coordinates
}

func (e *BoardError) Error() string {
	return e.Message
}
