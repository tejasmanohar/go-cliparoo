
package cliparoo

import (
	"errors"
	"os/exec"
	"runtime"
)

func Copy(input string) (err error) {
	switch runtime.GOOS {
	case "freebsd": execute("xsel -i -b", input)
	case "linux": execute("xclip -selection clipboard", input)
	case "darwin": execute("pbcopy", input)
	default: err = errors.New("unsupported platform")
	}
	return
}

func execute(command string, input string) (stdout []byte, err error) {
	cmd := exec.Command("sh", "-c", "printf " + input + "| " + command)
	stdout, err = cmd.Output()
	return
}
