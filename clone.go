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

func command_generator(line string, protocol string, destination string) string { // Function for generating the command for each repo
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
	fmt.Println("(1): HTTP") // HTTP will be used if anything other than 2 or SSH is entered, it is the default
	fmt.Println("(2): SSH")
	protocol := Reader()

	slugs := strings.Split(string(slugfile), "\n") // Makes slices of each line
	for _, line := range slugs {
		command_string := command_generator(line, protocol, destination) // Use the function to generate the command
		cmd := exec.Command(`bash`, `-c`, command_string) // Define the command

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Run() // Execute the command
	}

}
