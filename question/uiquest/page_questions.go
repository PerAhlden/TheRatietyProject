package uiquest

import (
	"TheRatietyProject/question"
	"fmt"
	"go.wdy.de/nago/pkg/xslices"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

func PageQuestions(repo question.Repository) core.View {
	questions, err := xslices.Collect2(repo.All())
	if err != nil {
		return alert.BannerError(err)
	}

	return ui.VStack(
		ui.Table(
			ui.TableColumn(ui.Text("Frage")),
			ui.TableColumn(ui.Text("Antwort 1")),
			ui.TableColumn(ui.Text("Antwort 2")),
			ui.TableColumn(ui.Text("Antwort 3")),
		).Rows(
			ui.ForEach(questions, func(v question.Question) ui.TTableRow {

				return ui.TableRow(
					ui.TableCell(ui.Text(v.Text)),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer0))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer1))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer2))),
				)
			})...,
		),
	).FullWidth()
}
