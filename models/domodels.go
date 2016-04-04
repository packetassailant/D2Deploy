package models

import "github.com/digitalocean/godo"

type DODeployStruct struct {
	Client           *godo.Client
	CurrentDeployNum int
	NewDeployNum     int
	DropletLimit     int
	DropletName      string
	Sshfprint        string
	Slug             string
	Size             string
	UserData         string
	RegionsAll       []string
}

type DropletStruct struct {
	ID        int
	Name      string
	IPAddress string
	Region    string
}

type ImageStruct struct {
	Name         string
	Distribution string
	Slug         string
}

type SizeStruct struct {
	Slug     string
	Memory   int
	VCPU     int
	Disk     int
	Transfer float64
	Cost     float64
}
