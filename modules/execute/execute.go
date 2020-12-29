package execute

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

//ExternalCommand Execute external command as subprocess.
func ExternalCommand(cmd ...string) {
	log.Printf("Executing : %s ...\n", cmd)
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		log.Panicln(err)
	}
	if err := c.Wait(); err != nil {
		log.Panicln(err)
	}
}

//CommandOutput represents the output of a command execution (stdout, stderr).
type CommandOutput struct {
	Stdout string
	Stderr string
}

//ExternalCommandWithRedirect Execute external command as subprocess and return stdout and stderr
//plus possible error.
func ExternalCommandWithRedirect(cmd ...string) (result CommandOutput, err error) {
	log.Printf("Executing : %s ...\n", cmd)

	c := exec.Command(cmd[0], cmd[1:]...)

	var stdoutBuffer, stderrBuffer bytes.Buffer
	c.Stdout = &stdoutBuffer
	c.Stderr = &stderrBuffer
	if err := c.Start(); err != nil {
		log.Panicln(err)
	}
	if err := c.Wait(); err != nil {
		log.Panicln(err)
	}
	return CommandOutput{Stdout: stdoutBuffer.String(), Stderr: stderrBuffer.String()}, nil
}
