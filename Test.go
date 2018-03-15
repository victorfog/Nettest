package main

import (
	"bufio"
	"fmt"
	"os"
	//"code.google.com/x/go.crypto/ssh"
	"golang.org/x/crypto/ssh"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter IP: ")
	ip, _ := reader.ReadString ('\n')
	fmt.Print("IP cisco: ", ip)
	
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	user, _ := reader1.ReadString('\n')
	fmt.Print("user", user)
	
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	passwd, _ := reader2.ReadString ('\n')
	fmt.Print("passwd", passwd)
	
	sshConfig := &ssh.ClientConfig {
		User: "&user",
		Auth: []ssh.AuthMethod {
			ssh.Password("&passwd")
		},
	
		
	}
	connection, err := ssh.Dial("tcp", "ip", sshConfig) //Establishing new SSH connection
	if err != nil {
		return nil, fmtErrorf("Failed to dial: %s", err)
	}
	
	session, err := connection.NewSession () //creating a new session
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}
	models := ssh.TerminalModels{
		ssh.ECHO: 0,
		ssh.TTY_OP_ISPEED: 14400, //input
		ssh.TTY_OP_OSPEED: 14400, //output
	}
	if err := session.RequestPty("xterm" , 80, 40, modes); err!= nil {
		session.Close()
		return nil, fmt.Errorf ( "request fot pseudo terminal failed: %s", err )
	}
	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("Unable to setup stdin for session: %v", err)
	}
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Unable to setup stdout for session: %v", err)
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		return fmt.Errorf("Unable to setup stderr for session: %v", err)
	}
	go io.Copy(os.Stderr, stderr)
	
	err = session.Run("ls -l $LC_USR_DIR")
	if err := session.Setenv("LC_USR_DIR", "/usr"); err != nil {
		return err
	}
}
