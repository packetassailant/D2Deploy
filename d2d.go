package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/digitalocean/godo"
	"github.com/packetassailant/D2Deploy/libs"
	"github.com/packetassailant/D2Deploy/models"
)

var (
	fConfig     = flag.String("config", "", "Provide an optional YAML configuration file")
	fToken      = flag.String("token", "", "The digital ocean oauth api token")
	fLimit      = flag.Bool("limit", false, "The DO droplet limit for this account")
	fRegions    = flag.Bool("regions", false, "The list of deployable DO geo-regions")
	fExisting   = flag.Bool("existing", false, "Retrieves currently deployed DO droplets")
	fImages     = flag.Bool("images", false, "Get all available deployable droplet images")
	fSizes      = flag.Bool("sizes", false, "Get all available slug sizes and rates")
	fSize       = flag.String("size", "512mb", "Set the size of the droplet (default: 512MB)")
	fDeployNum  = flag.Int("deploy", 0, "The number of DO droplets to deploy")
	fName       = flag.String("name", "", "The name to prepend each droplet")
	fSlug       = flag.String("slug", "", "The name of a deployable image slug (use images for slug types)")
	fSSHId      = flag.String("ssh", "", "The ssh key fingerprint for droplet auth")
	fDestroy    = flag.String("destroy", "", "The droplet slug to destroy (use existing for names)")
	fDestroyAll = flag.Bool("destroyall", false, "Destroy all droplet slugs (teardown all)")
)

var (
	deployNum   int
	dropletName string
	sshID       string
	slug        string
	size        string
	userData    string
)

func main() {
	flag.Parse()

	deployNum = *fDeployNum
	dropletName = *fName
	sshID = *fSSHId
	slug = *fSlug
	size = *fSize
	userData = ""

	dm := libs.DoDropletMarshaller{}
	t := libs.TokenSource{}
	t.Token()
	var client *godo.Client

	if *fToken != "" && *fConfig != "" {
		log.Fatal("The \"token\" flag and \"config\" flag are mutually exclusive.")
	}

	if *fToken == "" && *fConfig == "" {
		log.Fatal("Either the \"token\" flag or \"config\" flag are required.")
	}

	if *fConfig != "" {
		filename, _ := filepath.Abs(*fConfig)
		ymlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Unable to open the config file: %v", err)
		}

		cs := &models.DODeploymentStruct{}
		err = yaml.Unmarshal(ymlFile, cs)
		if err != nil {
			log.Fatalf("Unable to parse the config file: %v", err)
		}
		client = dm.GetClientHandle(cs.Token)
		deployNum = cs.DeployNum
		dropletName = cs.Name
		slug = cs.Slug
		sshID = cs.SSHId
		size = cs.Size

		uf, _ := filepath.Abs(cs.UserData)
		ufYmlFile, err := ioutil.ReadFile(uf)
		if err != nil {
			log.Fatalf("Unable to open the config file: %v", err)
		}
		userData = string(ufYmlFile)
	} else {
		client = dm.GetClientHandle(*fToken)
	}

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

	if *fImages {
		images, err := dm.GetAllImages(client)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Println("The following images are available for distribution:")
			for i := 0; i < len(images); i++ {
				fmt.Printf("Name: %v Distribution: %v Slug: %v\n", images[i].Name, images[i].Distribution, images[i].Slug)
			}
		}
	}

	if *fSizes {
		sizes, err := dm.GetDropletSizes(client)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Println("The following sizes are available (Cost is hourly):")
			for i := 0; i < len(sizes); i++ {
				fmt.Printf("Slug: %v Memory: %v VCPU: %v Disk: %v Transfer: %v Cost: %v\n",
					sizes[i].Slug, sizes[i].Memory, sizes[i].VCPU, sizes[i].Disk, sizes[i].Transfer, sizes[i].Cost)
			}
		}
	}

	if deployNum > 0 {
		if dropletName == "" {
			log.Fatal("The \"name\" flag and value are required.\n")
		}
		if sshID == "" {
			log.Fatal("the \"ssh\" flag and value are required")
		}
		if slug == "" {
			log.Fatal("the \"slug\" flag and value are required")
		}
		dds := &models.DODeployStruct{}
		existing, _ := dm.GetExistingDroplets(client)
		limit, _ := dm.GetDropLimit(client)
		regions, _ := dm.GetRegions(client)
		dds.Client = client
		dds.CurrentDeployNum = len(existing)
		dds.NewDeployNum = deployNum
		dds.DropletLimit = limit
		dds.DropletName = dropletName
		dds.Sshfprint = sshID
		dds.Slug = slug
		dds.Size = size
		dds.UserData = userData
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
