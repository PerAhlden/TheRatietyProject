package main

import (
	"TheRatietyProject/question"
	"TheRatietyProject/voting"
	"fmt"
	"go.wdy.de/nago/application/session"
	"go.wdy.de/nago/presentation/core"
	. "go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"strings"
)

// TODO ohne Namen keine Abstimmung
func PageVoting(wnd core.Window, votingRepo voting.Repository, questRepo question.Repository) core.View {
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

	var counter = core.AutoState[int](wnd)
	return VStack(
		Text(quest.Text).Font(Title),
		Text(fmt.Sprintf("%d", counter.Get())),
		TextField("Dein Name", nameState.Get()).InputValue(nameState).Disabled(vote.Voted),
		// TODO Disabled ergänzen mit oder (if) nutzername leer
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
