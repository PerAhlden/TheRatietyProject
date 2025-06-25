package question

import (
	"go.wdy.de/nago/pkg/data"
)

type ID string
type Question struct {
	ID      ID
	Text    string
	Answer0 string
	Answer1 string
	Answer2 string
	Answer3 string
	Answer4 string
}

func (q Question) Identity() ID {
	return q.ID
}

type Repository data.Repository[Question, ID]
