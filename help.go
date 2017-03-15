package main

import "fmt"

func printDeleteHelp() {
	helpmsg := `
dropletctl - digital ocean cli

  dropletctl delete droplet <name_of_droplet>
  dropletctl delete sshkey <name_of_sshkey>

`
	fmt.Println(helpmsg)
}
func printCreateHelp() {
	helpmsg := `
dropletctl - digital ocean cli

 create droplet:
  dropletctl create droplets -c 3 -r fra1 -n testsrv -k testkey

 create sshkey:
  dropletctl create sshkey -f sshkey.pub -n mykey

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

func printCreateSSHKeyHelp() {
	helpmsg := `
dropletctl - digiocean cli

 create sshkey:
  dropletctl create sshkey -help
  dropletctl create sshkey -f sshkey.pub -n mykey

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
