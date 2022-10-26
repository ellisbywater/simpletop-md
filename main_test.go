package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_makeUI(t *testing.T) {
	var testCfg config
	edit, preview := testCfg.makeUI()
	test.Type(edit, "Hello World")
	if preview.String() != "Hello World" {
		t.Error("RichText did not update")
	}
}

func Test_RunApp(t *testing.T) {
	var testCfg config
	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test")
	edit, preview := testCfg.makeUI()
	testCfg.createMenuItems(testWin)
	testWin.SetContent(container.NewHSplit(edit, preview))
	testApp.Run()
	test.Type(edit, "Hello World")
	if preview.String() != "Hello World" {
		t.Error("App Failed")
	}
}
