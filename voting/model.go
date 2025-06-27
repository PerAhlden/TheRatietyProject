package voting

import (
	"go.wdy.de/nago/application/session"
	"go.wdy.de/nago/pkg/data"
	"theRatietyProject/question"
)

type Voting struct {
	ID       session.ID
	Question question.ID
	Voted    bool
	Answer0  bool
	Answer1  bool
	Answer2  bool
	Answer3  bool
	Name     string
}

func (v Voting) Identity() session.ID {
	return v.ID
}

type Repository data.Repository[Voting, session.ID]
