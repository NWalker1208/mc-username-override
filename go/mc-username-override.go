package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/v4/process"
)

const (
	closeExistingInstance = true
	reportFileName        = "report.txt"
)

func main() {
	process := findMinecraftProcess()
	if process == nil {
		log.Fatal("Could not find a running instance of Minecraft Java Edition. Make sure Minecraft is running, then try again.")
	}

	cmdline, err := process.CmdlineSlice()
	if err != nil {
		log.Fatal(err)
	}

	username := inputUsername()
	modifiedCmdline := setUsername(cmdline, username)
	if modifiedCmdline == nil {
		log.Println("Could not figure out how to change the username. This script likely needs to be updated.")
		reportPath, err := filepath.Abs(reportFileName)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(reportPath, []byte(strings.Join(cmdline, "\n")), 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal("Please send the developer a copy of the report located here: " + reportPath)
	}

	restartProcess(process, modifiedCmdline)
	log.Println("Done! Minecraft may take a moment to finish opening.")
}
