package main

import "github.com/jroimartin/gocui"

func keybindings(g *gocui.Gui) error {
	g.InputEsc = true

	// global
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// projects
	if err := g.SetKeybinding("projects", gocui.KeyEsc, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyCtrlA, gocui.ModNone, inputView); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyEnter, gocui.ModNone, inputView); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyArrowRight, gocui.ModNone, selectItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyTab, gocui.ModNone, selectItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyCtrlR, gocui.ModNone, deleteItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyDelete, gocui.ModNone, deleteItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("projects", gocui.KeyCtrlD, gocui.ModNone, addDescription); err != nil {
		return err
	}

	// tasks
	if err := g.SetKeybinding("tasks", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyCtrlA, gocui.ModNone, inputView); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyEnter, gocui.ModNone, inputView); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyArrowRight, gocui.ModNone, selectItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyTab, gocui.ModNone, selectItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyCtrlR, gocui.ModNone, deleteItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyDelete, gocui.ModNone, deleteItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyArrowLeft, gocui.ModNone, goBack); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyEsc, gocui.ModNone, goBack); err != nil {
		return err
	}
	if err := g.SetKeybinding("tasks", gocui.KeyCtrlD, gocui.ModNone, addDescription); err != nil {
		return err
	}

	// entries
	if err := g.SetKeybinding("entries", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("entries", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("entries", gocui.KeyArrowLeft, gocui.ModNone, goBack); err != nil {
		return err
	}
	if err := g.SetKeybinding("entries", gocui.KeyEsc, gocui.ModNone, goBack); err != nil {
		return err
	}
	if err := g.SetKeybinding("entries", gocui.KeyCtrlA, gocui.ModNone, newEntry); err != nil {
		return err
	}
	if err := g.SetKeybinding("entries", gocui.KeyEnter, gocui.ModNone, newEntry); err != nil {
		return err
	}
	if err := g.SetKeybinding("entries", gocui.KeyCtrlR, gocui.ModNone, deleteItem); err != nil {
		return err
	}

	// output
	if err := g.SetKeybinding("output", gocui.KeyCtrlS, gocui.ModNone, save); err != nil {
		return err
	}

	return nil
}
