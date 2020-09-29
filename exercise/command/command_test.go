package command

import (
	"os/exec"
	"testing"
)

func Test_Command(t *testing.T) {
	const command = "docker ps | grep mysql:wine | awk '{print $11}' | sed -r 's/0.0.0.0:([0-9]+)->3306\\/tcp/\\1/g'"
	cmd := exec.Command("/bin/bash", "-c", command)
	docker, err := cmd.Output()
	if err != nil {
		t.Error(err)
		return
	}
	//
	t.Log(docker)
	return
}

/*func Test_Command(t *testing.T) {
	const command = "ping -c 10 127.0.0.1"
	cmd := exec.Command("/bin/bash","-c")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Error("cmd StdoutPipe error:", err)
		return
	}
	cmd.Start()

	var end_line string
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		end_line = line
	}
	//
	t.Log(cmd.Process.Pid)
	t.Log(end_line)
	return
}*/