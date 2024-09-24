package common

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path"
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

func ReadNonemptyLine(msg, msgWhenEmpty string) string {
	fmt.Print(msg)
	for {
		reply := strings.TrimSpace(Readline(""))
		if reply == "" {
			fmt.Println(msgWhenEmpty)
			continue
		}
		return reply
	}
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

func GetFrontierVersion() (string, error) {
	res, err := http.Get("https://raw.githubusercontent.com/Cosmoteer-Modding-Tools/frontier/refs/heads/main/version.txt")
	if err != nil {
		return "", err
	}

	version, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	} else if string(version) == "404: Not Found" {
		return "", err
	}

	return string(version), nil
}

func CheckForFrontierUpdate() (bool, [2]Version, error) {
	dir, err := GetHomeDir()
	if err != nil {
		return false, [2]Version{}, err
	}

	var remoteVersion Version

	remoteVersionText, err := GetFrontierVersion()
	if err != nil {
		return false, [2]Version{}, err
	}

	remoteVersion, err = NewVersionFromVersionString(remoteVersionText)
	if err != nil {
		return false, [2]Version{}, err
	}

	vpath := path.Join(path.Clean(dir), ".frontierversion")
	content, err := ReadFile(vpath)
	if err != nil {
		if ErrorIsFile404(err) {
			if err := WriteFile(vpath, remoteVersion.Fmt()); err != nil {
				return false, [2]Version{}, err
			}
		} else {
			return false, [2]Version{}, err
		}
	}

	localVersion, err := NewVersionFromVersionString(strings.TrimSpace(content))
	if err != nil {
		return false, [2]Version{}, err
	}

	return localVersion.Compare(remoteVersion) == -1, [2]Version{localVersion, remoteVersion}, nil
}
