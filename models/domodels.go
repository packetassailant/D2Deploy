package models

import "github.com/digitalocean/godo"

type DODeployStruct struct {
	Client           *godo.Client
	CurrentDeployNum int
	NewDeployNum     int
	DropletLimit     int
	DropletName      string
	Sshfprint        string
	RegionsAll       []string
}

type DropletStruct struct {
	ID        int
	Name      string
	IPAddress string
	Region    string
}
