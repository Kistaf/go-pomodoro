package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	timerTime = 0
	mode = "work"
)

func handleArguments() (int, string) {
	hourPtr := flag.Int("hours", 0, "an Int")
	minPtr := flag.Int("mins", 0, "an Int")
	secPtr := flag.Int("secs", 0, "an Int")
	presetPtr := flag.String("preset", "", "work | break")
	flag.Parse()

	// presets
	handlePresetModes(*presetPtr)

	// customized time
	if *hourPtr != 0 || *minPtr != 0 || *secPtr != 0 {
		timerTime = *secPtr + (*minPtr * 60) + (*hourPtr * 60 * 60)
	}

	if timerTime == 0 {
		fmt.Println("Please specify a countdown length using at least one flag:\n\t-hours={ex: 1}\n\t-mins={ex: 45}\n\t-secs={ex: 30}\n\t-preset={ex: work | break}")
		os.Exit(1)
	}
	return timerTime, mode
}

func handlePresetModes(presetPtr string) {
	if presetPtr != "" {
		switch presetPtr {
		case "work":
			timerTime = 25 * 60
		case "break":
			timerTime = 5 * 60
			mode = "break"
		}
	}
}
