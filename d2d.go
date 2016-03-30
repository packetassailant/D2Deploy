package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/packetassailant/D2Deploy/models"
	"github.com/packetassailant/D2Deploy/utils"
)

const token = "<enter digital ocean api key in between quotes>"

var (
	fLimit      = flag.Bool("limit", false, "The DO droplet limit for this account")
	fRegions    = flag.Bool("regions", false, "The list of deployable DO geo-regions")
	fExisting   = flag.Bool("existing", false, "Retrieves currently deployed DO droplets")
	fDeployNum  = flag.Int("deploy", 0, "The number of DO droplets to deploy")
	fName       = flag.String("name", "", "The name to prepend each droplet")
	fSSHSig     = flag.String("ssh", "", "The ssh key fingerprint for droplet auth")
	fDestroy    = flag.String("destroy", "", "The droplet slug to destroy (use existing for names)")
	fDestroyAll = flag.Bool("destroyall", false, "Destroy all droplet slugs (teardown all)")
)

func main() {
	flag.Parse()

	dm := utils.DoDropletMarshaller{}
	t := utils.TokenSource{}
	t.Token()
	client := dm.GetClientHandle(token)

	if *fLimit {
		limit, _ := dm.GetDropLimit(client)
		fmt.Printf("The VPS deployment limit is: %v\n", limit)
	}

	if *fRegions {
		regions, _ := dm.GetRegions(client)
		fmt.Printf("The available VPS deployment regions are: %v\n", strings.Join(regions, ","))
	}

	if *fExisting {
		existing, err := dm.GetExistingDroplets(client)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Println("The following droplets are currently deployed:")
			for i := 0; i < len(existing); i++ {
				fmt.Printf("Id: %v Name: %v IP: %v Region: %v\n", existing[i].ID, existing[i].Name, existing[i].IPAddress, existing[i].Region)
			}
		}
	}

	if *fDeployNum > 0 {
		if *fName == "" {
			log.Fatal("The \"name\" flag and value are required.\n")
		}
		if *fSSHSig == "" {
			log.Fatal("the \"ssh\" flag and value are required")
		}
		dds := &models.DODeployStruct{}
		existing, _ := dm.GetExistingDroplets(client)
		limit, _ := dm.GetDropLimit(client)
		regions, _ := dm.GetRegions(client)
		dds.Client = client
		dds.CurrentDeployNum = len(existing)
		dds.NewDeployNum = *fDeployNum
		dds.DropletLimit = limit
		dds.DropletName = *fName
		if *fSSHSig != "" {
			dds.Sshfprint = *fSSHSig
		}
		dds.RegionsAll = regions
		str, err := dm.DeployDroplet(dds)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(str)
	}

	if *fDestroy != "" {
		str, err := dm.DestroyDroplet(client, *fDestroy)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(str)
	}

	if *fDestroyAll {
		existing, _ := dm.GetExistingDroplets(client)
		str, err := dm.DestroyDropletAll(client, existing)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(str)
	}
}
