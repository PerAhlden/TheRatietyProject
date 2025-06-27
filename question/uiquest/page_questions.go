package uiquest

import (
	"TheRatietyProject/question"
	"fmt"
	"go.wdy.de/nago/pkg/data"
	"go.wdy.de/nago/pkg/xslices"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"go.wdy.de/nago/presentation/ui/form"
)

func PageQuestions(wnd core.Window, repo question.Repository) core.View {
	if !wnd.Subject().Valid() {
		return alert.Banner("Anmeldung erforderlich", "Um fortzufahren, müssen Sie sich am System anmelden.")
	}

	questions, err := xslices.Collect2(repo.All())
	if err != nil {
		return alert.BannerError(err)
	}

	createQuestDialogPresented := core.AutoState[bool](wnd)

	return ui.VStack(
		createQuestDialog(wnd, repo, createQuestDialogPresented),
		ui.HStack(
			ui.PrimaryButton(func() {
				createQuestDialogPresented.Set(true)
			}).Title("Neue Frage"),
		).FullWidth().Alignment(ui.Trailing),

		ui.Space(ui.L32),

		ui.Table(
			ui.TableColumn(ui.Text("Fragen")),
			ui.TableColumn(ui.Text("Antwort 1")),
			ui.TableColumn(ui.Text("Antwort 2")),
			ui.TableColumn(ui.Text("Antwort 3")),
			ui.TableColumn(ui.Text("Aktiv")),
			ui.TableColumn(ui.Text("Aktionen")),
		).Rows(
			ui.ForEach(questions, func(v question.Question) ui.TTableRow {

				return ui.TableRow(
					ui.TableCell(ui.Text(v.Text)),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer0))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer1))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Answer2))),
					ui.TableCell(ui.Text(fmt.Sprintf("%v", v.Active))),
					ui.TableCell(ui.SecondaryButton(func() {
						if err := question.ActivateQuestion(repo, v.ID); err != nil {
							alert.ShowBannerError(wnd, err)
							return
						}

						createQuestDialogPresented.Invalidate()
					}).Title("Frage aktivieren").Enabled(wnd.Subject().HasRole("asphaltier.role"))),
				)
			})...,
		),
	).FullWidth()
}

type CreateQuestionModel struct {
	Quest   string `label:"Frage" supportingText:"Hier deine Frage"`
	Answer0 string `label:"Antwort 1"`
	Answer1 string `label:"Antwort 2"`
	Answer2 string `label:"Antwort 3"`
}

func createQuestDialog(wnd core.Window, repo question.Repository, presented *core.State[bool]) core.View {
	if !presented.Get() {
		return nil
	}

	questionState := core.AutoState[CreateQuestionModel](wnd)

	return alert.Dialog(
		"Neue Frage erstellen",
		form.Auto(form.AutoOptions{Window: wnd}, questionState),
		presented,
		alert.Cancel(nil),
		alert.Save(func() (close bool) {
			model := questionState.Get()
			err := repo.Save(question.Question{
				ID:      data.RandIdent[question.ID](),
				Text:    model.Quest,
				Answer0: model.Answer0,
				Answer1: model.Answer1,
				Answer2: model.Answer2,
			})

			if err != nil {
				alert.ShowBannerError(wnd, err)
				return false
			}

			return true
		}),
	)
}
