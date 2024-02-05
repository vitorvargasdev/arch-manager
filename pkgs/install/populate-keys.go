package install

import (
	"fmt"
	"os"
	"os/exec"
)

func installArchLinuxKeyring() (string, error) {
	cmd := exec.Command("pacman", "-Sy", "--noconfirm", "archlinux-keyring")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not install archlinux-keyring: ", err)
	}

	return "", nil
}

func initializePacmanKeyring() (string, error) {
	cmd := exec.Command("pacman-key", "--init")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not initialize pacman-keyring: ", err)
	}

	return "", nil
}

func populatePacmanKeyring() (string, error) {
	cmd := exec.Command("pacman-key", "--populate")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not populate pacman-keyring: ", err)
	}

	return "", nil
}

func populateKeys() {
	installArchLinuxKeyring()
	initializePacmanKeyring()
	populatePacmanKeyring()
}
