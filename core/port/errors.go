package port

import "fmt"

type ErrEntityNoExist struct {
	ID string
}

func (e ErrEntityNoExist) Error() string {
	return fmt.Sprintf("entity with id %s does not exist", e.ID)
}
