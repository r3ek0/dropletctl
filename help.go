package main

import "fmt"

func printDeleteDropletHelp() {
	helpmsg := `
dropletctl - digital ocean cli

 Delete droplet:
  dropletctl delete droplet <name_of_droplet>

`
	fmt.Println(helpmsg)
}
func printCreateHelp() {
	helpmsg := `
dropletctl - digital ocean cli

 create droplet:
  dropletctl create droplet -help
  dropletctl create droplets -c 3 -r fra1 -n testsrv -k testkey
`
	fmt.Println(helpmsg)
}

func printCreateDropletHelp() {
	helpmsg := `
dropletctl - digital ocean cli

 create droplet:
  dropletctl create droplet -help
  dropletctl create droplets -c 3 -r fra1 -n testsrv -k testkey
`
	fmt.Println(helpmsg)
}

func printListHelp() {
	helpmsg := `
dropletctl - digital ocean cli

 List stuff:
  dropletctl list droplets
  dropletctl list keys

`
	fmt.Println(helpmsg)
}

func printBasicHelp() {
	helpmsg := `
dropletctl - digital ocean cli

 Usage:
  dropletctl <mode> <item> [<flags>]

 Modes:
  create	create an item
  delete	delete an item
  list		list an item

 Items:
  keys, sshkeys
  droplet, droplets
`
	fmt.Println(helpmsg)
}
