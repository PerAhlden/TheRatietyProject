package uivoting

import (
	"fmt"
	"go.wdy.de/nago/application/session"
	"go.wdy.de/nago/pkg/xslices"
	"go.wdy.de/nago/presentation/core"
	. "go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"strings"
	"theRatietyProject/question"
	"theRatietyProject/voting"
)

func PageVoting(wnd core.Window, votingRepo voting.Repository, questRepo question.Repository) core.View {
	votings, err := xslices.Collect2(votingRepo.All())
	if err != nil {
		return alert.BannerError(err)
	}

	questions, err := xslices.Collect2(questRepo.All())
	if err != nil {
		return alert.BannerError(err)

	}

	invalidate := core.AutoState[int](wnd)

	optActiveQuest, err := question.FindActiveQuestion(questRepo)
	// TODO error handling
	if optActiveQuest.IsNone() {
		return Text("Keine aktive Frage zum abstimmen")
	}

	quest := optActiveQuest.Unwrap()
	id := session.ID(strings.Join([]string{string(wnd.Session().ID()), string(quest.ID)}, ""))

	optVoting, err := votingRepo.FindByID(id)
	if err != nil {
		return alert.BannerError(err)
	}

	var defaultEntry voting.Voting
	vote := optVoting.UnwrapOr(defaultEntry)
	vote.ID = id

	nameState := core.AutoState[string](wnd).Init(func() string {
		return vote.Name
	})

	//var counter = core.AutoState[int](wnd)
	giveCount := make(map[question.ID]int)

	for _, q := range questions {
		if q.ID == quest.ID {
			counter := 0

			for _, v := range votings {
				if v.Question == q.ID && v.Voted {
					counter++
				}
			}

			giveCount[q.ID] = counter
		}
	}

	return VStack(
		Text(quest.Text).Font(Title),
		Text(fmt.Sprintf("%d", giveCount[quest.ID])),
		TextField("Dein Name", nameState.Get()).InputValue(nameState).Disabled(vote.Voted),

		SecondaryButton(func() {
			vote.Answer0 = true
			vote.Voted = true
			vote.Question = quest.ID
			vote.Name = nameState.Get()
			if err := votingRepo.Save(vote); err != nil {
				alert.ShowBannerError(wnd, err)
			}
			invalidate.Invalidate()
		}).Title(quest.Answer0).Enabled(!vote.Voted && nameState.Get() != ""),

		If(quest.Answer1 != "", SecondaryButton(func() {
			vote.Answer1 = true
			vote.Voted = true
			vote.Question = quest.ID
			vote.Name = nameState.Get()
			if err := votingRepo.Save(vote); err != nil {
				alert.ShowBannerError(wnd, err)
			}
			invalidate.Invalidate()
		}).Title(quest.Answer1).Enabled(!vote.Voted && nameState.Get() != "")),

		If(quest.Answer2 != "", SecondaryButton(func() {
			vote.Answer2 = true
			vote.Voted = true
			vote.Question = quest.ID
			vote.Name = nameState.Get()
			if err := votingRepo.Save(vote); err != nil {
				alert.ShowBannerError(wnd, err)
			}
			invalidate.Invalidate()
		}).Title(quest.Answer2).Enabled(!vote.Voted && nameState.Get() != "")),
	).Gap(L16).
		Padding(Padding{}.All(L16))
}
