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
	"go.wdy.de/nago/application/role"
	"go.wdy.de/nago/application/session"
	"go.wdy.de/nago/pkg/std"
	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/hero/solid"
	. "go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/web/vuejs"
	"time"
)

//go:embed vote-for-blog.jpg
var logo application.StaticBytes

func main() {
	application.Configure(func(cfg *application.Configurator) {
		cfg.SetApplicationID("de.worldiety.tutorial")
		cfg.Serve(vuejs.Dist())
		option.MustZero(cfg.StandardSystems())
		roleManagement := std.Must(cfg.RoleManagement())
		userManagement := std.Must(cfg.UserManagement())
		votingLogo := cfg.Resource(logo)

		questionRepo := application.SloppyRepository[question.Question, question.ID](cfg)
		votingRepo := application.SloppyRepository[voting.Voting, session.ID](cfg)

		cfg.SetDecorator(cfg.NewScaffold().
			Logo(Image().URI(votingLogo).Frame(Frame{Height: L96})).
			Login(true).
			MenuEntry().Title("Start").Icon(icons.Newspaper).Forward(".").OneOfRole().
			MenuEntry().Title("Aktuelle Abstimmung").Icon(icons.UserGroup).Forward("aktuelleAbstimmung").OneOfRole().
			//MenuEntry().Title("Chat").Icon(icons.ChatBubbleLeft).Forward("chat").OneOfRole().
			MenuEntry().Title("Übersicht").Icon(icons.ArchiveBox).Forward("overview").OneOfRole().
			MenuEntry().Title("Alte Fragen").Icon(icons.QuestionMarkCircle).Forward("oldquestions").Private().
			Decorator())
		std.Must(std.Must(cfg.UserManagement()).UseCases.EnableBootstrapAdmin(time.Now().Add(time.Hour), "%6UbRsCuM8N$auy"))
		std.Must(roleManagement.UseCases.Upsert(userManagement.UseCases.SysUser(), role.Role{
			ID:          "asphaltier.role",
			Name:        "Asphaltier",
			Description: "Per ist ein Asphaltier",
			Permissions: nil,
		}))

		cfg.RootViewWithDecoration(".", func(wnd core.Window) core.View {
			return VStack(
				Text("Willkommen bei Ratiety").Font(Font{Size: L40}).Frame(Frame{Height: L96}).Padding(Padding{Top: L16}),
				Text("Sie möchten auch mal ihre Meinung äußern, weil man Ihnen sonst nie zuhört? Oder aber ist es so, dass sie reden jeden Tag mit "+
					"unmengen an Leuten und sie kommen gar nicht mehr klar, weil sie so vielen Menschen ihre Meinung mitteilen ? "+
					"Egal was sie denken, teilen sie es mit anderen. Alle unsere Meinungen sollten gehört werden und alle Menschen "+
					"sollten ihre Meinung äußern dürfen. Weil ich das denke erfand ich Ratiety, eine Plattform auf der jeder eine"+
					" Stimme hat. Nehmen sie teil an der wöchentlichen Umfrage und sehen sie wer genauso denkt wie sie. Hier "+
					"können sie anhand von alltäglichen Fragen, die sich viele stellen Ihre Meinung kundtun. Also was lesen sie "+
					"hier noch so blöd rum? Machen sie sich auf, stimmen sie ab und das Allerwichtigste: Haben sie möglichst viel Freude daran. "+
					"Ratiety wünscht Ihnen viel Spaß bei diesem überaus lustigen Zeitvertreib. ").Font(Font{Size: L16}).
					Frame(Frame{Width: L480}).AccessibilityLabel("Easter Egg: Banane"))

		})

		cfg.RootViewWithDecoration("aktuelleAbstimmung", func(wnd core.Window) core.View {
			return HStack(
				PageVoting(wnd, votingRepo, questionRepo),
				VLine(),
			).Gap(L16).Alignment(Top).Frame(Frame{}.FullWidth())
		})

		//cfg.RootViewWithDecoration("chat", func(wnd core.Window) core.View {
		//	return VStack(Text("Dies ist der Chat. Probieren sie gerne den Rest der menschlichen Bevölkerung auf dieser Plattform von " +
		//		"Ihrer, sicherlich guten, Meinung zu überzeugen. Übrigens sind jegliche Arten von Beleidigungen vollkommen unangebracht und " +
		//		"müssen mit dem möglichen Ausschluss aus diesem Chat bestraft werden. Trotzdessen wünschen wir Ihnen weiterhin noch viel Spaß. " +
		//		"(Der Chat muss noch programmiert werden und existiert deswegen noch nicht.)").Font(Font{Size: L20}).Frame(Frame{Width: L480}))
		//})
		// TODO chat repository erstellen genau wie voting repository, so hat man einen leichten chat
		cfg.RootViewWithDecoration("overview", func(wnd core.Window) core.View {
			return PageVotingOverview(votingRepo, questionRepo)
		})

		cfg.RootViewWithDecoration("oldquestions", func(wnd core.Window) core.View {
			return uiquest.PageQuestions(wnd, questionRepo)
		})

	}).Run()
}
