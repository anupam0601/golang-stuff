package awsftp

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Awsftp : function to be exported
func Awsftp(localpath string) {
	client, err := ssh.Dial("tcp", "host:22", &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("$yncron1234"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	fmt.Println(err)
	fmt.Println(client)

	// SSH session
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	fmt.Println("Successfully connected to ssh server.")
	defer session.Close()

	// Open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created SFTP conn....")
	defer sftp.Close()

	// srcPath := "/home/anupam/files/test_demo_anupam.txt"
	// dstPath := "/input/test_demo_anupam.txt"
	// // fileName := "test_anupam_demo.txt"

	//  Local file path to test   and   Folder on a remote machine
	// var localFilePath = localpath
	var remoteDir = "/input/"
	srcFile, err := os.Open(localpath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localpath)
	dstFile, err := sftp.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	fmt.Println("copy file to remote server finished!")

	//STEP 4. After the connection is created,
	//we can run commands on the remote machine using Run command using the session that we just created.
	// var b bytes.Buffer
	// session.Stdout = &b
	// if err := session.Run("/usr/bin/ls"); err != nil {
	// 	panic("Failed to run: " + err.Error())
	// }
	// fmt.Println(b.String())
}
