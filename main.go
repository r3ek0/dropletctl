package main

import (
	"flag"
	"fmt"
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

func main() {

	if len(os.Args) < 2 {
		printBasicHelp()
		os.Exit(1)
	}

	createDropletFlags := flag.NewFlagSet("Flags for create droplet", flag.ExitOnError)
	var dropletparams DropletParams
	createDropletFlags.StringVar(&dropletparams.Name, "-name", "testlet", "Name of the droplets")
	createDropletFlags.StringVar(&dropletparams.Name, "n", "testlet", "Name of the droplets (shorthand)")

	createDropletFlags.StringVar(&dropletparams.Region, "-region", "fra1", "Region in which the droplet should live.")
	createDropletFlags.StringVar(&dropletparams.Region, "r", "fra1", "Region in which the droplet should live.")

	createDropletFlags.StringVar(&dropletparams.Image, "-image", "ubuntu-16-04-x64", "Name of the image to use.")
	createDropletFlags.StringVar(&dropletparams.Image, "i", "ubuntu-16-04-x64", "Name of the image to use.")

	createDropletFlags.StringVar(&dropletparams.SSHKeyName, "-key", "none", "Name of the sshkey to use.")
	createDropletFlags.StringVar(&dropletparams.SSHKeyName, "k", "none", "Name of the sshkey to use.")

	createDropletFlags.StringVar(&dropletparams.Size, "-size", "512mb", "Size of the droplet (512mb, 1gb, 2gb...)")
	createDropletFlags.StringVar(&dropletparams.Size, "s", "512mb", "Size of the droplet (512mb, 1gb, 2gb...)")

	createDropletFlags.IntVar(&dropletparams.Count, "-count", 1, "Number of droplets")
	createDropletFlags.IntVar(&dropletparams.Count, "c", 1, "Number of droplets")

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
		os.Exit(0)

	case "delete":
		if len(os.Args) < 3 {
			printDeleteDropletHelp()
			os.Exit(1)
		}

		if stringInSlice(os.Args[2], []string{"droplet", "droplets"}) {
			if len(os.Args) < 4 {
				printDeleteDropletHelp()
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
