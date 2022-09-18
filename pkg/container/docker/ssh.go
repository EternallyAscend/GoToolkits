package docker

import "fmt"

// StartSshServerInDockerCommand Connect with `ssh -p port user@ip`
func StartSshServerInDockerCommand(newPwd, user string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("apt-get install -y openssh-server openssh-clients"))
	cmds = append(cmds, fmt.Sprintf("echo \"%s\" | passwd --stdin %s", newPwd, user))
	cmds = append(cmds, fmt.Sprintf("echo \"PermitRootLogin yes\" >> /etc/ssh/sshd_config"))
	cmds = append(cmds, fmt.Sprintf("/etc/init.d/sshd restart"))
	return cmds
}
