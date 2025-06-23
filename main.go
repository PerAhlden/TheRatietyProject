package main

import (
	"go.wdy.de/nago/application"
	"go.wdy.de/nago/presentation/core"
	. "go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/web/vuejs"
)

// the main function of the program, which is like the java public static void main.
func main() {
	// we use the applications package to bootstrap our configuration
	application.Configure(func(cfg *application.Configurator) {
		cfg.SetApplicationID("de.worldiety.tutorial_01")
		cfg.Serve(vuejs.Dist())

		cfg.RootView(".", func(wnd core.Window) core.View {
			return VStack(
				Text("hello per").BackgroundColor("#19a6a1"),
				Text("Willkommen im Team ").BackgroundColor("#19a6a1"),
			).
				Frame(Frame{MinWidth: "2000px", MinHeight: "600px"}).Border(Border{
				LeftWidth: "10px",
				TopWidth:  "10px",
				LeftColor: "#ad1a37",
				TopColor:  "#1aad41",
			})

		})
	}).
		// don't forget to call the run method, which starts the entire thing and blocks until finished
		Run()
}
