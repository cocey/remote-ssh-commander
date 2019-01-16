package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Version constant
const Version string = "1.0.0-beta"

// help message string
var helpMsg = `Remote SSH Commander = run commands on remote devices via SSH
usage: rsshc [options]
options:
`

// print help
func printHelp() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}

// Print Version
func printVersion() {
	fmt.Printf("aurimg version %v\n", Version)
}

var hostKey ssh.PublicKey

type Device struct {
	Name     string
	Ip       string
	Username string
	Password string
	Commands []string
}

func run4Device(d Device) {
	config := &ssh.ClientConfig{
		User: d.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(d.Password),
		},
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Println("connecting to " + d.Ip)
	client, err := ssh.Dial("tcp", d.Ip+":22", config)

	if err != nil {
		//panic("Failed to dial: " + err.Error())
		fmt.Println("Failed to dial: " + err.Error())
	} else {
		fmt.Println("connected")
		fmt.Println("running commands")
		for _, c := range d.Commands {
			session, err := client.NewSession()
			if err != nil {
				//log.Fatal("Failed to create session: ", err)
				fmt.Println("Failed to create session: ", err)
			} else {
				//defer session.Close()
				var b bytes.Buffer
				session.Stdout = &b

				command := strings.Replace(c, "{%name%}", d.Name, -1)
				command = strings.Replace(command, "{%ip%}", d.Ip, -1)

				if err := session.Run(command); err != nil {
					fmt.Println("Failed to run: " + err.Error())
				} else {
					fmt.Println(b.String())
				}
				session.Close()
			}

		}
	}

}
func main() {
	var versionFlag = flag.Bool("v", false, "output version information and exit.")
	var helpFlag = flag.Bool("h", false, "display this help dialog")
	var commandsFile = flag.String("c", "commands.json", "commands file path in json format")

	flag.Parse()

	if *helpFlag == true {
		printHelp()
		os.Exit(0)
	}

	if *versionFlag == true {
		printVersion()
		os.Exit(0)
	}
	if *commandsFile == "" {
		printHelp()
		os.Exit(0)
	}

	file, e := ioutil.ReadFile(*commandsFile)
	if e != nil {
		fmt.Printf("Commands file not found: %v\n", e)
		printHelp()
		os.Exit(1)
	}
	//fmt.Printf("reading commands %s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var devices []Device
	json.Unmarshal(file, &devices)
	//fmt.Printf("Results: %v\n", devices)
	for _, d := range devices {
		fmt.Println("Running for : ", d.Name)
		//fmt.Println(d)
		run4Device(d)
	}

}
