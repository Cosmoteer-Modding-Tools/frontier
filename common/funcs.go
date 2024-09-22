package common

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func WriteFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func ErrorIsFile404(err error) bool {
	return strings.HasSuffix(strings.ToLower(err.Error()), ": no such file or directory") || strings.HasSuffix(strings.ToLower(err.Error()), ": The system cannot find the file specified.")
}

func RunCommand(command string, args []string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func Readline(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}

func DoesItemExist(item string) bool {
	check := `printf "if [ -f "` + item + `" ] || [ -d "` + item + `" ]; then\necho "true"\nelse\necho "false"\nfi" | sh`
	if strings.ToLower(runtime.GOOS) == "windows" {
		check = "Test-Path " + item
	}
	so, se, err := RunCommand(check, []string{})
	if se != "" || err != nil {
		return false
	}
	return strings.TrimSpace(strings.ToLower(so)) == "true"
}
