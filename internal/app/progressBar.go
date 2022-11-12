package app

import (
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func (s *Service) ProgressBarFilesAndFolders() {
	time.Sleep(500 * time.Millisecond)
	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("  [green][1/2]Creating Folders...[reset] "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Millisecond)
	}

	bar = progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("  [green][2/2]Writing files...[reset] "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Millisecond)
	}

	bar.Close()

	time.Sleep(300 * time.Millisecond)
}
func (s *Service) ProgressBarGoModules() {
	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("  [cyan][1/2]Executing go mod init...[reset] "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[cyan]=[reset]",
			SaucerHead:    "[cyan]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	for i := 0; i < 800; i++ {
		bar.Add(2)
		time.Sleep(1 * time.Millisecond)
	}
	bar = progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("  [cyan][2/2]Executing go mod tidy...[reset] "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[cyan]=[reset]",
			SaucerHead:    "[cyan]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	for i := 0; i < 800; i++ {
		bar.Add(2)
		time.Sleep(1 * time.Millisecond)
	}
	bar.Close()
}
