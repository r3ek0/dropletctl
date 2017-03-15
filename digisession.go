package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/digitalocean/godo"
)

type DigiSession struct {
	Client godo.Client
}

func NewDigiSession() *DigiSession {
	oauthClient := NewOAuthClient()
	client := godo.NewClient(oauthClient)
	session := &DigiSession{
		Client: *client,
	}
	return session
}

func (ds *DigiSession) DeleteDroplet(rm_droplet string) error {
	droplet, err := ds.GetDropletByName(rm_droplet)
	if err != nil {
		return err
	}
	if droplet.ID == 0 {
		return errors.New("Droplet not found.")
	}

	fmt.Printf("delete droplet: %8d - %s %s ", droplet.ID, droplet.Name, droplet.SizeSlug)
	_, err = ds.Client.Droplets.Delete(context.TODO(), droplet.ID)
	if err != nil {
		fmt.Printf(" [error]\n")
		return err
	}
	fmt.Printf(" [done]\n")
	return nil
}

func (ds *DigiSession) GetDropletByName(name string) (*godo.Droplet, error) {
	list, err := ds.DropletList()
	if err != nil {
		return nil, err
	}

	for _, droplet := range list {
		if droplet.Name == name {
			return &droplet, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No droplet with name '%s' found.", name))
}

func (ds *DigiSession) GetSSHKeyByName(name string) (*godo.Key, error) {

	sshkeys, err := ds.SSHKeyList()
	if err != nil {
		return nil, err
	}

	for _, key := range sshkeys {
		if key.Name == name {
			return &key, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No key with name '%s' found.", name))
}

func (ds *DigiSession) ListSSHKeys() error {
	sshkeys, err := ds.SSHKeyList()
	if err != nil {
		return err
	}

	for _, key := range sshkeys {
		fmt.Printf("Key: %12d Name: %16s FP: %s\n", key.ID, key.Name, key.Fingerprint)
	}
	return nil

}

func (ds *DigiSession) CreateSSHKey(sshkeyparams SSHKeyParams) (*godo.Key, error) {
	ctx := context.TODO()
	req := &godo.KeyCreateRequest{
		Name:      sshkeyparams.SSHKeyName,
		PublicKey: sshkeyparams.SSHKey,
	}
	key, _, err := ds.Client.Keys.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (ds *DigiSession) DeleteSSHKey(name string) (*godo.Key, error) {
	keylist, err := ds.SSHKeyList()
	if err != nil {
		return nil, err
	}
	for _, key := range keylist {
		if key.Name == name {
			_, err := ds.Client.Keys.DeleteByFingerprint(context.TODO(), key.Fingerprint)
			if err != nil {
				return nil, err
			}
			return &key, nil
		}
	}
	return nil, errors.New("No such key.")
}

func (ds *DigiSession) SSHKeyList() ([]godo.Key, error) {
	opts := &godo.ListOptions{}
	list := []godo.Key{}
	ctx := context.TODO()

	for {
		sshkeys, resp, err := ds.Client.Keys.List(ctx, opts)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range sshkeys {
			list = append(list, d)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opts.Page = page + 1
	}

	return list, nil
}

func (ds *DigiSession) DropletList() ([]godo.Droplet, error) {
	ctx := context.TODO()
	// create a list to hold our droplets
	list := []godo.Droplet{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}

	for {
		droplets, resp, err := ds.Client.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			list = append(list, d)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}

func (ds *DigiSession) ListDroplets() error {
	list, err := ds.DropletList()
	if err != nil {
		return err
	}
	for _, droplet := range list {
		fmt.Printf("(%6s) %-15s %4s %5s %2dcpu %10s %s\n",
			droplet.Status,
			droplet.Name,
			droplet.Region.Slug,
			droplet.SizeSlug,
			droplet.Vcpus,
			droplet.Image.Slug,
			droplet.Networks.V4[0].IPAddress,
		)
	}
	return nil
}
func (ds *DigiSession) CreateDroplets(params DropletParams) ([]godo.Droplet, error) {
	for c := 0; c < params.Count; c++ {
		dropletName := fmt.Sprintf("%s-%02d", params.Name, c+1)
		_, found := ds.GetDropletByName(dropletName)
		if found == nil {
			return nil, errors.New("A droplet with name '" + dropletName + "' already exists.\nAbort.")
		}
	}
	var droplets []godo.Droplet
	var sshkeys []godo.DropletCreateSSHKey

	key, err := ds.GetSSHKeyByName(params.SSHKeyName)
	if err != nil {
		return nil, err
	}
	sshkey := &godo.DropletCreateSSHKey{
		ID:          key.ID,
		Fingerprint: key.Fingerprint,
	}
	sshkeys = append(sshkeys, *sshkey)

	fmt.Printf("Params: %+v\n", params)
	for c := 0; c < params.Count; c++ {
		dropletName := fmt.Sprintf("%s-%02d", params.Name, c+1)
		createRequest := &godo.DropletCreateRequest{
			Name:   dropletName,
			Region: params.Region,
			Size:   params.Size,
			Image: godo.DropletCreateImage{
				Slug: params.Image,
			},
			SSHKeys: sshkeys,
		}

		newDroplet, _, err := ds.Client.Droplets.Create(context.TODO(), createRequest)
		if err != nil {
			return nil, err
		}
		droplets = append(droplets, *newDroplet)

	}
	for _, droplet := range droplets {
		fmt.Printf("(%6s) %-15s %4s %5s %dcpu %15s\n",
			droplet.Status,
			droplet.Name,
			droplet.Region.Slug,
			droplet.SizeSlug,
			droplet.Vcpus,
			droplet.Image.Slug,
		)

	}
	return droplets, nil
}
