package terminal

import tui "github.com/marcusolsson/tui-go"

func newTUI() {
	box := tui.NewVBox(
		tui.NewLabel("tui-go"),
	)

	ui, err := tui.New(box)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
