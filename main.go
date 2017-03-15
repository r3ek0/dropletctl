package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type DropletParams struct {
	Count      int
	Region     string
	Name       string
	Size       string
	Image      string
	SSHKeyName string
}

type SSHKeyParams struct {
	SSHKey     string
	SSHKeyFile string
	SSHKeyName string
}

func main() {

	if len(os.Args) < 2 {
		printBasicHelp()
		os.Exit(1)
	}

	createDropletFlags := flag.NewFlagSet("Flags for create droplet", flag.ContinueOnError)
	var dropletparams DropletParams
	createDropletFlags.StringVar(&dropletparams.Name, "name", "testlet", "Name of the droplets")
	createDropletFlags.StringVar(&dropletparams.Name, "n", "testlet", "Name of the droplets (shorthand)")

	createDropletFlags.StringVar(&dropletparams.Region, "region", "fra1", "Region in which the droplet should live.")
	createDropletFlags.StringVar(&dropletparams.Region, "r", "fra1", "Region in which the droplet should live.")

	createDropletFlags.StringVar(&dropletparams.Image, "image", "ubuntu-16-04-x64", "Name of the image to use.")
	createDropletFlags.StringVar(&dropletparams.Image, "i", "ubuntu-16-04-x64", "Name of the image to use.")

	createDropletFlags.StringVar(&dropletparams.SSHKeyName, "key", "none", "Name of the sshkey to use.")
	createDropletFlags.StringVar(&dropletparams.SSHKeyName, "k", "none", "Name of the sshkey to use.")

	createDropletFlags.StringVar(&dropletparams.Size, "size", "512mb", "Size of the droplet (512mb, 1gb, 2gb...)")
	createDropletFlags.StringVar(&dropletparams.Size, "s", "512mb", "Size of the droplet (512mb, 1gb, 2gb...)")

	createDropletFlags.IntVar(&dropletparams.Count, "count", 1, "Number of droplets")
	createDropletFlags.IntVar(&dropletparams.Count, "c", 1, "Number of droplets")

	var sshkeyparams SSHKeyParams
	createSSHKeyFlags := flag.NewFlagSet("sshkey flags", flag.ContinueOnError)
	createSSHKeyFlags.StringVar(&sshkeyparams.SSHKeyName, "name", "none", "Name for the sshkey")
	createSSHKeyFlags.StringVar(&sshkeyparams.SSHKeyName, "n", "none", "Name for the sshkey")

	createSSHKeyFlags.StringVar(&sshkeyparams.SSHKey, "key", "none", "The public sshkey")
	createSSHKeyFlags.StringVar(&sshkeyparams.SSHKey, "k", "none", "The public sshkey")

	createSSHKeyFlags.StringVar(&sshkeyparams.SSHKeyFile, "file", "none", "Read key from file.")
	createSSHKeyFlags.StringVar(&sshkeyparams.SSHKeyFile, "f", "none", "Read key from file.")

	switch os.Args[1] {
	case "create":
		if len(os.Args) < 3 {
			printCreateHelp()
			os.Exit(1)
		}

		if stringInSlice(os.Args[2], []string{"droplet", "droplets"}) {
			fmt.Println("Create droplets")
			if len(os.Args) < 4 {
				printCreateDropletHelp()
				os.Exit(1)
			}
			createDropletFlags.Parse(os.Args[3:])
			session := NewDigiSession()
			_, err := session.CreateDroplets(dropletparams)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if stringInSlice(os.Args[2], []string{"sshkey"}) {
			fmt.Println("create a new sshkey entry.")
			createSSHKeyFlags.Parse(os.Args[3:])
			if sshkeyparams.SSHKeyFile == "none" {
				fmt.Println("You must specify a file (-f) which contains the public key.")
				os.Exit(1)
			}
			keyfileblob, err := ioutil.ReadFile(sshkeyparams.SSHKeyFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			sshkeyparams.SSHKey = string(keyfileblob)
			session := NewDigiSession()
			newkey, err := session.CreateSSHKey(sshkeyparams)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("newkey:\n%+v\n", newkey)
		}
		os.Exit(0)

	case "delete":
		if len(os.Args) < 3 {
			printBasicHelp()
			os.Exit(1)
		}
		if stringInSlice(os.Args[2], []string{"sshkey", "sshkeys"}) {
			if len(os.Args) < 4 {
				printDeleteHelp()
				os.Exit(1)
			}
			name := os.Args[3]
			session := NewDigiSession()
			key, err := session.DeleteSSHKey(name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Delete key:\n%+v\n", key)
			os.Exit(0)
		}
		if stringInSlice(os.Args[2], []string{"droplet", "droplets"}) {
			if len(os.Args) < 4 {
				printDeleteHelp()
				os.Exit(1)
			}
			name := os.Args[3]
			session := NewDigiSession()
			err := session.DeleteDroplet(name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		os.Exit(0)

	case "list":
		if len(os.Args) < 3 {
			printListHelp()
			os.Exit(1)
		}
		if stringInSlice(os.Args[2], []string{"keys", "sshkeys"}) {
			session := NewDigiSession()
			session.ListSSHKeys()
		} else if stringInSlice(os.Args[2], []string{"droplet", "droplets"}) {
			session := NewDigiSession()
			session.ListDroplets()
		}
		os.Exit(0)

	default:
		printBasicHelp()
		os.Exit(1)

	}
}
