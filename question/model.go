package question

import (
	"fmt"
	"github.com/worldiety/option"
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
	Active  bool // If a question is active, this question will be the current one to vote for
}

func (q Question) Identity() ID {
	return q.ID
}

type Repository data.Repository[Question, ID]

func ActivateQuestion(repo Repository, id ID) error {
	for quest, err := range repo.All() {
		if err != nil {
			return fmt.Errorf("cannot range over questions: %w", err)
		}

		if quest.ID == id {
			quest.Active = true
		} else {
			quest.Active = false
		}

		if err := repo.Save(quest); err != nil {
			return fmt.Errorf("cannot save question: %w", err)
		}
	}

	return nil
}

func FindActiveQuestion(repo Repository) (option.Opt[Question], error) {
	for quest, err := range repo.All() {
		// TODO error handling
		_ = err
		if quest.Active {
			return option.Some(quest), nil
		}
	}

	// TODO return none for not found

	return option.Opt[Question]{}, nil
}
