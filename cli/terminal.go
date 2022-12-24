// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/anigkus/kush/sys"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type Terminal struct {
	Address    string
	Username   string
	Password   string
	PublicKey  string
	Port       int32
	Client     *ssh.Client
	LastResult string
}

func New(address string, username string, password string, publicKey string, port int32) *Terminal {
	terminal := new(Terminal)
	terminal.Address = address
	terminal.Username = username
	terminal.Password = password
	terminal.PublicKey = publicKey
	terminal.Port = port
	return terminal
}

// connect establish ssh connection
func (t *Terminal) connect() error {
	auth := make([]ssh.AuthMethod, 1)
	if t.Password != "" || len(t.Password) > 0 {
		auth[0] = ssh.Password(t.Password)
	}
	if t.PublicKey != "" || len(t.PublicKey) > 0 {
		pemBytes, err := ioutil.ReadFile(t.PublicKey)
		if err != nil {
			return err
		}
		signer, err := ssh.ParsePrivateKey(pemBytes)
		if err != nil {
			return err
		}
		auth[0] = ssh.PublicKeys(signer)
	}
	// config := ssh.ClientConfig{
	// 	Config: ssh.Config{},
	// 	User:   t.Username,
	// 	Auth:   auth,
	// 	HostKeyCallback: func(hostname string, remote net.Addr, publicKey ssh.PublicKey) error {
	// 		return nil
	// 	},
	// 	BannerCallback: func(message string) error {
	// 		return nil
	// 	},
	// 	ClientVersion:     "",
	// 	HostKeyAlgorithms: []string{},
	// 	Timeout:           30 * time.Second,
	// }

	config := ssh.ClientConfig{
		User: t.Username,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, publicKey ssh.PublicKey) error {
			return nil
		},
		Timeout: sys.TERMINAL_TIMEOUT,
	}

	addr := fmt.Sprintf("%s:%d", t.Address, t.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	t.Client = sshClient
	return nil
}

// RunTerminal reader terminal parameters
func (t *Terminal) RunTerminal(stdout, stderr io.Writer) error {
	if t.Client == nil {
		if err := t.connect(); err != nil {
			return err
		}
	}
	session, err := t.Client.NewSession()
	if err != nil {
		return err
	}
	// close2
	defer func() {
		if err := session.Close(); err != nil {
			// Required: It must be like this,otherwise there will be an extra % character at the end of the terminal
			// Combined with the close1 call
			fmt.Println("")
		}
	}()

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState)

	session.Stdout = stdout
	session.Stderr = stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := term.GetSize(fd)
	if err != nil {
		return err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		return err
	}
	// start remote shell
	if err := session.Shell(); err != nil {
		return err
	} // open1 else {}
	if err := session.Wait(); err != nil {
		return err
	} else {
		// close 1
		// Combined with the close2 call
		fmt.Print("Connection to " + t.Address + " closed.")
	}
	return nil
}
