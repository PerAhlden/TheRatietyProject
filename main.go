// Copyright (c) 2025 worldiety GmbH
// This file is part of the NAGO Low-Code Platform.
// Licensed under the terms specified in the LICENSE file.
// SPDX-License-Identifier: Custom-License
package main

import (
	"TheRatietyProject/question"
	"TheRatietyProject/question/uiquest"
	"TheRatietyProject/voting"
	_ "embed"
	"github.com/worldiety/option"
	"go.wdy.de/nago/application"
	"go.wdy.de/nago/application/session"
	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/hero/solid"
	. "go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/web/vuejs"
)

//go:embed vote-for-blog.jpg
var logo application.StaticBytes

func main() {
	application.Configure(func(cfg *application.Configurator) {
		cfg.SetApplicationID("de.worldiety.tutorial")
		cfg.Serve(vuejs.Dist())

		votingLogo := cfg.Resource(logo)

		questionRepo := application.SloppyRepository[question.Question, question.ID](cfg)
		votingRepo := application.SloppyRepository[voting.Voting, session.ID](cfg)

		option.MustZero(questionRepo.Save(question.Question{
			ID:      "1",
			Text:    "Ist meine erste Frage gut?",
			Answer0: "Sehr weise gewählt",
			Answer1: "Ist einigermaßen ok",
			Answer2: "Verdammt schlecht",
		}))

		cfg.SetDecorator(cfg.NewScaffold().
			Logo(Image().URI(votingLogo).Frame(Frame{Height: L96})).
			Login(false).
			MenuEntry().Title("Start").Icon(icons.Newspaper).Forward(".").OneOfRole().
			MenuEntry().Title("Aktuelle Abstimmung").Icon(icons.UserGroup).Forward("aktuelleAbstimmung").OneOfRole().
			MenuEntry().Title("Chat").Icon(icons.ChatBubbleLeft).Forward("chat").OneOfRole().
			MenuEntry().Title("Übersicht").Icon(icons.ArchiveBox).Forward("overview").OneOfRole().
			MenuEntry().Title("Alte Fragen").Icon(icons.QuestionMarkCircle).Forward("fragen").OneOfRole().
			Decorator())

		cfg.RootViewWithDecoration(".", func(wnd core.Window) core.View {
			return VStack(
				Text("Dies ist Ratiety, die Plattform der ewigen Bewertung. Hier können sie wöchentlich alles und jeden in einer bestimmten" +
					" Kategorie bewerten, aber seihen sie gewarnt: Auch sie werden nicht verschont bleiben! Scheuen sie sich nicht davor teilzunehmen und krönen " +
					"sie sich zum Sieger, der ganz alleine am höchsten Punkt der Rangliste verweilen darf."))
		})

		cfg.RootViewWithDecoration("aktuelleAbstimmung", func(wnd core.Window) core.View {
			return HStack(
				PageVoting(wnd, votingRepo),
				VLine(),
			).Gap(L16).Alignment(Top).Frame(Frame{}.FullWidth())
		})

		cfg.RootViewWithDecoration("chat", func(wnd core.Window) core.View {
			return VStack(Text("Dies ist der Chat. Probieren sie gerne den Rest der menschlichen Bevölkerung auf dieser Plattform von " +
				"Ihrer, sicherlich guten, Meinung zu überzeugen. Übrigens sind jegliche Arten von Beleidigungen vollkommen unangebracht und " +
				"müssen mit dem möglichen Ausschluss aus diesem Chat bestraft werden. Trotzdessen wünschen wir Ihnen weiterhin noch viel Spaß."))
		})

		cfg.RootViewWithDecoration("overview", func(wnd core.Window) core.View {
			return PageVotingOverview(votingRepo)
		})

		cfg.RootViewWithDecoration("fragen", func(wnd core.Window) core.View {
			return uiquest.PageQuestions(questionRepo)
		})

	}).Run()
}
