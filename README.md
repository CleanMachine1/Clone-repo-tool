# Clone-repo-tool

Clone-repo-tool is a simple tool to help import Git repos easier when setting up an environment.
First you create a `slugfile` such as [the one in this repo](https://github.com/CleanMachine1/Clone-repo-tool/blob/3a1b5bc8679854421b699e01dee531239e6c966b/slugfile.txt) using the same syntax.

The script supports SSH and HTTPS and allows you to save the downloaded repositories to your chosen destination.

## Installation

Install a [Go compiler](https://go.dev/doc/install)

Then clone the repository:

`git clone https://github.com/CleanMachine1/Clone-repo-tool`

Then finally, cd into the directory

`cd Clone-repo-tool`

## Usage

`go run clone.go`
