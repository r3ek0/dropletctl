# dropletctl
Commandline tool to manage digital ocean droplets
It uses https://github.com/digitalocean/godo 

Install
-------

Install go

```
export GOPATH=<path_to_go_dev>
go get github.com/r3ek0/dropletctl...
go build github.com/r3ek0/dropletctl
```

Run
---
```
export DIGIOCEAN_TOKEN=<your-api-token>
./dropletctl help
```

List SSH keys
-------------
```
./dropletctl list keys
```

Create droplets
---------------

```
./dropletctl create droplets -help
./dropletctl create droplets -n testsrv -c 3 -s 512mb -k mysshkeyname -r nyc1
```

List droplets
--------------

```
./dropletctl list droplets
```

Delete droplets
---------------

```
./dropletctl delete droplet testsrv-01
```

