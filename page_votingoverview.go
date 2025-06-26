package main

import (
	"TheRatietyProject/question"
	"TheRatietyProject/voting"
	"go.wdy.de/nago/pkg/xslices"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"strconv"
)

func PageVotingOverview(votingRepo voting.Repository, questionRepo question.Repository) core.View {
	votings, err := xslices.Collect2(votingRepo.All())
	if err != nil {
		return alert.BannerError(err)
	}

	questions, err := xslices.Collect2(questionRepo.All())
	if err != nil {
		return alert.BannerError(err)
	}
	matchVwithQ := make(map[question.ID]int)

	for _, q := range questions {
		counter := 0

		for _, v := range votings {
			if v.Question == q.ID && v.Voted {
				counter++
			}
		}

		matchVwithQ[q.ID] = counter
	}

	return ui.VStack(
		ui.ForEach(questions, func(q question.Question) core.View {
			counter := matchVwithQ[q.ID]
			return ui.VStack(
				ui.H2(q.Text),
				ui.Table(
					ui.TableColumn(ui.Text("Abstimmungen Insgesamt")),
					ui.TableColumn(ui.Text(q.Answer0)),
					ui.TableColumn(ui.Text(q.Answer1)),
					ui.TableColumn(ui.Text(q.Answer2)),
				).Rows(
					ui.TableRow(
						ui.TableCell(
							ui.Text(strconv.Itoa(counter)),
						),
						ui.TableCell(
							ui.VStack(
								ui.ForEach(votings, func(v voting.Voting) core.View {
									if v.Question == q.ID && v.Answer0 == true {
										return ui.Text(v.Name)
									}

									return nil
								})...,
							),
						),
						ui.TableCell(
							ui.VStack(
								ui.ForEach(votings, func(v voting.Voting) core.View {
									if v.Question == q.ID && v.Answer1 == true {
										return ui.Text(v.Name)
									}

									return nil
								})...,
							),
						),
						ui.TableCell(
							ui.VStack(
								ui.ForEach(votings, func(v voting.Voting) core.View {
									if v.Question == q.ID && v.Answer2 == true {
										return ui.Text(v.Name)
									}

									return nil
								})...,
							),
						),
					)),
			).FullWidth()
		})...,
	).Gap(ui.L32)
}
