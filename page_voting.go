package main

import (
	"TheRatietyProject/voting"
	"fmt"
	"go.wdy.de/nago/presentation/core"
	. "go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

func PageVoting(wnd core.Window, votingRepo voting.Repository) core.View {
	invalidate := core.AutoState[int](wnd)

	optVoting, err := votingRepo.FindByID(wnd.Session().ID())
	if err != nil {
		return alert.BannerError(err)
	}

	var defaultEntry voting.Voting
	defaultEntry.ID = wnd.Session().ID()
	vote := optVoting.UnwrapOr(defaultEntry)

	nameState := core.AutoState[string](wnd).Init(func() string {
		return vote.Name
	})

	var counter = core.AutoState[int](wnd)
	return VStack(
		Text("The actual Voting ").Font(Title),
		Text(fmt.Sprintf("%d", counter.Get())),
		TextField("Dein Name", nameState.Get()).InputValue(nameState).Disabled(vote.Voted),

		SecondaryButton(func() {
			vote.Answer0 = true
			vote.Voted = true
			vote.Name = nameState.Get()
			if err := votingRepo.Save(vote); err != nil {
				alert.ShowBannerError(wnd, err)
			}

			invalidate.Invalidate()
		}).Title("First Candidate").Enabled(!vote.Voted),
		SecondaryButton(func() {
			vote.Answer1 = true
			vote.Voted = true
			vote.Name = nameState.Get()
			if err := votingRepo.Save(vote); err != nil {
				alert.ShowBannerError(wnd, err)
			}
			invalidate.Invalidate()
		}).Title("Second Candidate").Enabled(!vote.Voted),
	).Gap(L16).
		Padding(Padding{}.All(L16))
}
