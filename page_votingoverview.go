package main

import (
	"TheRatietyProject/voting"
	"fmt"
	"go.wdy.de/nago/pkg/xslices"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

func PageVotingOverview(repo voting.Repository) core.View {
	votings, err := xslices.Collect2(repo.All())
	if err != nil {
		return alert.BannerError(err)
	}

	return ui.VStack(
		ui.Table(
			ui.TableColumn(ui.Text("Name")),
			ui.TableColumn(ui.Text("Antwort 1")),
			ui.TableColumn(ui.Text("Antwort 2")),
			ui.TableColumn(ui.Text("Antwort 3")),
		).Rows(
			ui.ForEach(votings, func(v voting.Voting) ui.TTableRow {

				return ui.TableRow(
					ui.TableCell(ui.Text(v.Name)),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer0))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer1))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer2))),
				)
			})...,
		),
	).FullWidth()
}
