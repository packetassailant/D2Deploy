package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/packetassailant/D2Deploy/models"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type DoDropletMarshaller struct{}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func (dm *DoDropletMarshaller) GetClientHandle(token string) *godo.Client {
	tokenSource := &TokenSource{
		AccessToken: token,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)
	return client
}

func (dm *DoDropletMarshaller) GetDropLimit(c *godo.Client) (int, error) {
	acctDetails, _, err := c.Account.Get()
	if err != nil {
		fmt.Printf("Could not retrieve account details: %s\n\n", err)
		return 0, err
	}
	dropletLimit := acctDetails.DropletLimit
	return dropletLimit, nil
}

func (dm *DoDropletMarshaller) GetExistingDroplets(c *godo.Client) ([]models.DropletStruct, error) {
	dropletLO := godo.ListOptions{}
	eDroplets, _, err := c.Droplets.List(&dropletLO)
	if len(eDroplets) <= 0 {
		return nil, fmt.Errorf("There are NO droplets currently deployed\n")
	}
	d := &models.DropletStruct{}
	var dropletList []models.DropletStruct

	for _, v := range eDroplets {
		ip, _ := v.PublicIPv4()
		d.ID = v.ID
		d.Name = v.Name
		d.IPAddress = ip
		d.Region = v.Region.Name
		dropletList = append(dropletList, *d)
	}

	if err != nil {
		return nil, fmt.Errorf("Could not retrieve existing droplets: %s\n\n", err)
	}
	return dropletList, nil
}

func (dm *DoDropletMarshaller) GetRegions(c *godo.Client) ([]string, error) {
	regionLO := godo.ListOptions{}
	regionList, _, err := c.Regions.List(&regionLO)
	if err != nil {
		fmt.Printf("Could not obtain regions: %s\n\n", err)
		return nil, err
	}

	var regionsAll []string
	for r := range regionList {
		regionsAll = append(regionsAll, regionList[r].Slug)
	}
	return regionsAll, nil
}

func (dm *DoDropletMarshaller) DeployDroplet(s *models.DODeployStruct) (string, error) {
	c := s.Client
	allowableDroplets := s.DropletLimit - s.CurrentDeployNum
	if allowableDroplets == 0 {
		return "", fmt.Errorf("Deployment not allowed: Droplets maximum capacity is " + strconv.Itoa(s.DropletLimit))
	}

	depIdx := 0
	regIdx := 0
	for depIdx < s.NewDeployNum {
		// Inexpensive technique to create circular regions list
		if depIdx >= len(s.RegionsAll) {
			s.RegionsAll = append(s.RegionsAll, s.RegionsAll...)
		}
		createRequest := &godo.DropletCreateRequest{
			Name:   s.DropletName + "-" + strconv.Itoa(depIdx),
			Region: s.RegionsAll[regIdx],
			Size:   "512mb",
			// ssh-keygen -lf ~/.ssh/id_rsa.pub | cut -d " " -f 2
			SSHKeys: []godo.DropletCreateSSHKey{
				godo.DropletCreateSSHKey{
					Fingerprint: s.Sshfprint,
				},
			},
			Image: godo.DropletCreateImage{
				// Deploy a 14.04 build w/ Docker pre-installed
				Slug: "docker",
			},
			// The section after #cloud-config can be edited with cloud config scripting
			UserData: `
			#cloud-config
			runcmd:
				- apt-get install -y git
			`,
		}
		_, res, err := c.Droplets.Create(createRequest)

		// Cycle through 422 status and error substrings
		if res.StatusCode == 422 && strings.Contains(err.Error(), "invalid key identifiers") {
			return "", fmt.Errorf("Error configuring SSH key(s): %v", err)
		}
		if res.StatusCode == 422 && strings.Contains(err.Error(), "Region is not available") {
			fmt.Printf("Region unavailable, trying the next region: %s\n\n", err)
			regIdx++
			continue
		}
		if err != nil {
			return "", fmt.Errorf("Error creating droplet: %s", err)
		}
		regIdx++
		depIdx++
	}
	return fmt.Sprint("Successfully created all droplets\n"), nil
}

func (dm *DoDropletMarshaller) DestroyDroplet(c *godo.Client, id string) (string, error) {
	intID, _ := strconv.Atoi(id)
	res, err := c.Droplets.Delete(intID)
	if err != nil || res.StatusCode != 204 {
		return "", fmt.Errorf("Failed to delete droplet: %v", intID)
	}
	return fmt.Sprintf("Successfully destroyed droplet: %v\n", intID), nil
}

func (dm *DoDropletMarshaller) DestroyDropletAll(c *godo.Client, d []models.DropletStruct) (string, error) {
	count := 0
	for count < len(d) {
		res, err := c.Droplets.Delete(d[count].ID)
		if err != nil || res.StatusCode != 204 {
			return "", fmt.Errorf("Failed to delete droplet: %v", d[count].ID)
		}
		count++
	}
	return fmt.Sprint("Successfully destroyed all droplets\n"), nil
}
