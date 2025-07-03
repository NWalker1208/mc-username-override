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

func findMinecraftProcess() *process.Process {
	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			continue
		}
		if strings.ToLower(name) == "javaw.exe" {
			cmdline, err := process.CmdlineSlice()
			if err != nil {
				continue
			}
			if containsMainClass(cmdline) {
				return process
			}
		}
	}
	return nil
}

func containsMainClass(cmdline []string) bool {
	mainClasses := []string{
		"net.minecraft.client.main.Main",
		"cpw.mods.bootstraplauncher.BootstrapLauncher",
		"net.fabricmc.loader.impl.launch.knot.KnotClient",
	}
	for _, mainClass := range mainClasses {
		for _, arg := range cmdline {
			if arg == mainClass {
				return true
			}
		}
	}
	return false
}

func inputUsername() string {
	fmt.Print("Enter your fake username: ")
	var username string
	fmt.Scanln(&username)
	return username
}

func setUsername(cmdline []string, username string) []string {
	for i, arg := range cmdline {
		if arg == "--username" {
			cmdline[i+1] = username
			return cmdline
		}
	}
	return nil
}
