package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	createDropletFlags := flag.NewFlagSet("Flags for create droplet", flag.ExitOnError)
	var dropletparams DropletParams
	createDropletFlags.StringVar(&dropletparams.Name, "-name", "testlet", "Name of the droplets")
	createDropletFlags.StringVar(&dropletparams.Name, "n", "testlet", "Name of the droplets (shorthand)")

	createDropletFlags.StringVar(&dropletparams.Region, "-region", "fra1", "Region in which the droplet should live.")
	createDropletFlags.StringVar(&dropletparams.Region, "r", "fra1", "Region in which the droplet should live.")

	createDropletFlags.StringVar(&dropletparams.Image, "-image", "ubuntu-16-04-x64", "Name of the image to use.")
	createDropletFlags.StringVar(&dropletparams.Image, "i", "ubuntu-16-04-x64", "Name of the image to use.")

	createDropletFlags.StringVar(&dropletparams.SSHKeyFP, "-key", "none", "FP of the sshkey to use.")
	createDropletFlags.StringVar(&dropletparams.SSHKeyFP, "k", "none", "FP of the sshkey to use.")

	createDropletFlags.StringVar(&dropletparams.Size, "-size", "512mb", "Size of the droplet (512mb, 1gb, 2gb...)")
	createDropletFlags.StringVar(&dropletparams.Size, "s", "512mb", "Size of the droplet (512mb, 1gb, 2gb...)")

	createDropletFlags.IntVar(&dropletparams.Count, "-count", 1, "Number of droplets")
	createDropletFlags.IntVar(&dropletparams.Count, "c", 1, "Number of droplets")

	switch os.Args[1] {
	case "create":
		fmt.Println("Create droplets")
		createDropletFlags.Parse(os.Args[3:])
		session := NewDigiSession()
		_, err := session.CreateDroplets(dropletparams)
		if err != nil {
			panic(err)
		}
		os.Exit(0)

	case "delete":
		if os.Args[2] == "droplet" {
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
		if os.Args[2] == "keys" {
			session := NewDigiSession()
			session.ListSSHKeys()
		} else if os.Args[2] == "droplets" {
			session := NewDigiSession()
			session.listDroplets()
		}
		os.Exit(0)

	default:
		os.Exit(0)

	}
}
