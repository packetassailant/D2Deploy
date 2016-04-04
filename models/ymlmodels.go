package models

type DODeploymentStruct struct {
	Token     string `yaml:"token"`
	DeployNum int    `yaml:"deploy_num"`
	Name      string `yaml:"droplet_name"`
	Slug      string `yaml:"slug_name"`
	Size      string `yaml:"slug_size"`
	SSHId     string `yaml:"ssh_id"`
	UserData  string `yaml:"userdata_file"`
}
