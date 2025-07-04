package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/shirou/gopsutil/v4/process"
)

const (
	closeExistingInstance = true
	reportFileName        = "report.txt"
	creationFlags         = syscall.CREATE_NEW_PROCESS_GROUP | 0x00000008 /* DETACHED_PROCESS */
)

func main() {
	process := findMinecraftProcess()
	if process == nil {
		fmt.Println("Could not find a running instance of Minecraft Java Edition. Make sure Minecraft is running, then try again.")
		os.Exit(1)
	}

	cmdline, err := process.Cmdline()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	username := inputUsername()
	modifiedCmdline := setUsername(cmdline, username)
	if modifiedCmdline == nil {
		fmt.Println("Could not figure out how to change the username. This script likely needs to be updated.")
		reportPath, err := filepath.Abs(reportFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = os.WriteFile(reportPath, []byte(cmdline), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Please send the developer a copy of the report located here: " + reportPath)
		os.Exit(1)
	}

	restartProcess(process, *modifiedCmdline)
	fmt.Println("Done! Minecraft may take a moment to finish opening.")
}

func findMinecraftProcess() *process.Process {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			continue
		}
		if strings.ToLower(name) == "javaw.exe" {
			cmdline, err := process.Cmdline()
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

func containsMainClass(cmdline string) bool {
	mainClasses := []string{
		"net.minecraft.client.main.Main",
		"cpw.mods.bootstraplauncher.BootstrapLauncher",
		"net.fabricmc.loader.impl.launch.knot.KnotClient",
	}
	for _, mainClass := range mainClasses {
		if strings.Contains(cmdline, mainClass) {
			return true
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

func setUsername(cmdline string, username string) *string {
	usernameArgument := "--username"
	argIndex := strings.Index(cmdline, " "+usernameArgument+" ") + len(usernameArgument) + 2
	if argIndex == -1 {
		return nil
	}

	prevUsernameLen := strings.Index(cmdline[argIndex:], " ")
	if prevUsernameLen == -1 {
		prevUsernameLen = len(cmdline[argIndex:])
	}

	cmdline = cmdline[:argIndex] + username + cmdline[argIndex+prevUsernameLen:]
	return &cmdline
}

func restartProcess(p *process.Process, cmdline string) {
	cwd, err := p.Cwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	env, err := p.Environ()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if closeExistingInstance {
		// 	fmt.Println("Closing Minecraft...")
		// 	closeWindows(p)
		// 	err = p.Wait()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		fmt.Println("Restarting Minecraft...")
	} else {
		fmt.Println("Opening Minecraft...")
	}

	// There seems to be a bug in CommandLineToArgv that truncates individual arguments to 8192 characters.
	// As a workaround, we use the cmdline string directly.
	cmdNameLen := strings.Index(cmdline, " ")
	if cmdNameLen == -1 {
		cmdNameLen = len(cmdline)
	}
	cmd := exec.Command(cmdline[:cmdNameLen])
	cmd.Dir = cwd
	cmd.Env = env
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CmdLine:       cmdline,
		CreationFlags: creationFlags,
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// TODO: Briefly wait for program to exit before assuming it has started successfully
}
