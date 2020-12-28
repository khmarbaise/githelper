package execute

import (
	"log"
	"os"
	"os/exec"
)

//RunExternalCommand Execute external command as subprocess.
func RunExternalCommand(cmd ...string) {
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
