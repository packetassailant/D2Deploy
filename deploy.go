package d2deploy

import (
	"fmt"
	"strconv"
	"strings"

	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"
)

type DoDropletMarshaller struct{}

type dropletsStruct struct {
	droplet []dropletStruct
}

type dropletStruct struct {
	ID        int    `json:"godo.Droplet"`
	Name      string `json:"godo.Droplet"`
	IPAddress string `json:"godo.NetworkV4"`
}

func (dm *DoDropletMarshaller) GetClientHandle(token string) *godo.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	client := godo.NewClient(t.Client())
	return client
}

func (dm *DoDropletMarshaller) GetDropLimit(c *godo.Client) (int, error) {
	acctDetails, _, err := c.Account.Get()
	if err != nil {
		fmt.Printf("Could not retrieve account details: %s\n\n", err)
		return 0, err
	}
	dropletLimit := acctDetails.Account.DropletLimit
	return dropletLimit, nil
}

func (dm *DoDropletMarshaller) GetExistingDroplets(c *godo.Client) ([]godo.Droplet, error) {
	dropletLO := godo.ListOptions{}
	eDroplets, _, err := c.Droplets.List(&dropletLO)
	if err != nil {
		fmt.Printf("Could not retrieve existing droplets: %s\n\n", err)
		return nil, err
	}
	return eDroplets, nil
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

// currentDeployNum = the number of droplets currently deployed
// newDeployNum = the number of new droplets to create and deploy
// dropletLimit = the number of droplets allowable on the specified DO account
func (dm *DoDropletMarshaller) DeployDroplet(c *godo.Client, currentDeployNum, newDeployNum, dropletLimit int, dropletName, sshfprint string, regionsAll []string) {

	allowableDroplets := dropletLimit - currentDeployNum
	if allowableDroplets == 0 {
		fmt.Println("Deployment not allowed: Droplets maximum capacity is " + strconv.Itoa(dropletLimit))
		return
	}

	depIdx := 0
	regIdx := 0
	for depIdx < newDeployNum {
		// Inexpensive technique to create circular regions list
		if depIdx >= len(regionsAll) {
			regionsAll = append(regionsAll, regionsAll...)
		}
		createRequest := &godo.DropletCreateRequest{
			Name:   dropletName + "-" + strconv.Itoa(depIdx),
			Region: regionsAll[regIdx],
			Size:   "512mb",
			// ssh-keygen -lf ~/.ssh/id_rsa.pub | cut -d " " -f 2
			SSHKeys: []godo.DropletCreateSSHKey{
				godo.DropletCreateSSHKey{
					Fingerprint: sshfprint,
				},
			},
			Image: godo.DropletCreateImage{
				// Deploy a 14.04 build w/ Docker pre-installed
				Slug: "docker",
			},
			UserData: `
			#cloud-config
			runcmd:
				- apt-get install -y git
			`,
		}
		result, res, err := c.Droplets.Create(createRequest)

		// Cycle through 422 status and error substrings
		if res.StatusCode == 422 && strings.Contains(err.Error(), "invalid key identifiers") {
			fmt.Printf("Error configuring SSH key(s): %s\n\n", err)
			return
		}
		if res.StatusCode == 422 && strings.Contains(err.Error(), "Region is not available") {
			fmt.Printf("Region unavailable, trying the next region: %s\n\n", err)
			regIdx++
			continue
		}
		if err != nil {
			fmt.Printf("Error creating droplet: %s\n\n", err)
			return
		} else {
			fmt.Println(result)
			fmt.Println(res)
			regIdx++
			depIdx++
		}
	}
}

func (dm *DoDropletMarshaller) DestroyDroplet(c *godo.Client, d godo.Droplet) {
	res, err := c.Droplets.Delete(d.ID)
	if err != nil || res.StatusCode != 204 {
		fmt.Printf("Failed to delete droplet: %s\n\n", err)
		return
	}
}

func (dm *DoDropletMarshaller) DestroyDropletAll(c *godo.Client, d []godo.Droplet) {
	count := 0
	for count < len(d) {
		res, err := c.Droplets.Delete(d[count].ID)
		if err != nil || res.StatusCode != 204 {
			fmt.Printf("Failed to delete droplet: %s\n\n", err)
			return
		}
		count++
	}
}
