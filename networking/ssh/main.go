package main

import (
	"bytes"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	client, err := ssh.Dial("tcp", "192.168.0.104:22", &ssh.ClientConfig{
		User: "anupamdebnath",
		Auth: []ssh.AuthMethod{
			ssh.Password("ANUPAM"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	fmt.Println(err)
	fmt.Println(client)

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	//STEP 4. After the connection is created,
	//we can run commands on the remote machine using Run command using the session that we just created.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
