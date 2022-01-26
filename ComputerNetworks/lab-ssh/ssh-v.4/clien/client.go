package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	username := "alex"
	hostname := "0.0.0.0"
	port := "2200"

	client, session, err := connectToHost(username, hostname+":"+port)
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	pass := "123"
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
