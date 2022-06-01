package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Reader() string { // Function for collecting user input easier as a string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func command_generator(line string, protocol string, destination string) string {
	if protocol == "2" || protocol == "SSH" || protocol == "ssh" {
		return "cd " + destination + " && git clone git@github.com:" + line
	} else {
		return "cd " + destination + " && git clone https://github.com/" + line
	}
}
func main() {

	fmt.Println("Please enter the location of the slug file (exact location from root)")
	source := Reader() // Receive user input for the location of the slug file

	slugfile, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the destination folder for the Git repositories (exact location from root)")
	destination := Reader()

	fmt.Println("Which protocol would you like to use?")
	fmt.Println("(1): HTTP")
	fmt.Println("(2): SSH")
	protocol := Reader()

	slugs := strings.Split(string(slugfile), "\n")
	for _, line := range slugs {
		command_string := command_generator(line, protocol, destination)
		cmd := exec.Command(`bash`, `-c`, command_string)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Run()
	}

}
