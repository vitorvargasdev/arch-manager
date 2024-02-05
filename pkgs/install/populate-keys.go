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
		return "", err
	}

	return "", nil
}

func initializePacmanKeyring() (string, error) {
	cmd := exec.Command("pacman-key", "--init")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not initialize pacman-keyring: ", err)
		return "", err
	}

	return "", nil
}

func populatePacmanKeyring() (string, error) {
	cmd := exec.Command("pacman-key", "--populate")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not populate pacman-keyring: ", err)
		return "", err
	}

	return "", nil
}

func populateKeys() (string, error) {
	if _, err := installArchLinuxKeyring(); err != nil {
		fmt.Println("Error installing archlinux-keyring: ", err)
		return "", err
	}

	if _, err := initializePacmanKeyring(); err != nil {
		fmt.Println("Error initializing pacman-keyring: ", err)
		return "", err
	}

	if _, err := populatePacmanKeyring(); err != nil {
		fmt.Println("Error populating pacman-keyring: ", err)
		return "", err
	}

	return "done", nil
}
