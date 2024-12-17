package main

import (
	"os"
	"accustomedtogo/warning"
)

func main() {
	var warn warning.Warning // warningインターフェース型の変数warnを宣言

	warn = &warning.ConsoleWarning{}
	warn.Show("Hello World to console")
	warn = &warning.DesktopWarning{}
	warn.Show("Hello World to desktop")
	warn = &warning.SlackWarning{
		URL: os.Getenv("SLACK_URL"),
		Channel: "#general",
	}
	warn.Show("Hello World to slack")

}

