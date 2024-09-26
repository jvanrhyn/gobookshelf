package customererrors

type (
	RecordNotFoundError struct {
		Message string
	}
)

func (e RecordNotFoundError) Error() string {
	return e.Message
}
