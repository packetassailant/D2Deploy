# D2Deploy

## Objective
```
A mass Digital Ocean droplet deployment utility
```

## Usage
```
$ ./D2Deploy -help
Usage of ./D2Deploy:
  -deploy int
    	The number of DO droplets to deploy
  -destroy string
    	The droplet slug to destroy (use existing for names)
  -destroyall
    	Destroy all droplet slugs (teardown all)
  -existing
    	Retrieves currently deployed DO droplets
  -limit
    	The DO droplet limit for this account
  -name string
    	The name to prepend each droplet
  -regions
    	The list of deployable DO geo-regions
  -ssh string
    	The ssh key fingerprint for droplet auth
```

## Installation
```
Installation
---------------------------------------------------
1. Install GO (tested on 1.5.2)
2 .Git clone this repo (git clone https://github.com/packetassailant/D2Deploy.git)
3. Edit the d2d.go file variable "token" to represent the actual DO API key
4. cd into the repo and type go build (you will now have a D2Deploy binary)
```

## Sample Run - limit
```
This will provide the number of droplet instances associated with the
DO account.

$ ./D2Deploy -limit
The VPS deployment limit is: 25

```

## Sample Run - regions
```
This will provide the global geographic regions available for mass droplet
deployment. (Note this will be variant depending on regional network/system saturation)

$ ./D2Deploy -regions
The available VPS deployment regions are: nyc1,sfo1,nyc2,ams2,sgp1,lon1,nyc3,ams3,fra1,tor1

```

## Sample Run - deploy
```
This will perform the actual mass deployment of specific droplet instances. The
"ssh" option is required to do public key auth with every deployed system. Further,
the ssh fingerprint can be obtained via the 'ssh-keygen -lf ~/.ssh/id_rsa.pub | cut -d " " -f 2'
command.

The following command deploys 10 droplets with the name "packetresearch" appended,
while automatically installing the ssh public key at time of build.

$ ./D2Deploy -deploy 10 -name "packetresearch" -ssh "aa:bb:cc:dd:ee:ff:gg:hh:ii:jj:kk:ll:mm:nn:oo:pp"
Successfully created all droplets

```

## Sample Run - existing
```
This will provide details about the the droplets that are currently deployed.
Information includes the droplet ID, the droplet name, the droplet IP, and the region.

$ ./D2Deploy -existing
The following droplets are currently deployed:
Id: 12563675 Name: packetresearch-0 IP: 198.199.120.27 Region: New York 1
Id: 12563676 Name: packetresearch-1 IP: 159.203.235.30 Region: San Francisco 1
Id: 12563677 Name: packetresearch-2 IP: 162.243.65.179 Region: New York 2
Id: 12563678 Name: packetresearch-3 IP: 146.185.131.62 Region: Amsterdam 2
Id: 12563679 Name: packetresearch-4 IP: 188.166.247.115 Region: Singapore 1
Id: 12563680 Name: packetresearch-5 IP: 188.166.168.111 Region: London 1
Id: 12563681 Name: packetresearch-6 IP: 104.236.208.178 Region: New York 3
Id: 12563682 Name: packetresearch-7 IP: 178.62.208.130 Region: Amsterdam 3
Id: 12563683 Name: packetresearch-8 IP: 46.101.243.210 Region: Frankfurt 1
Id: 12563684 Name: packetresearch-9 IP: 159.203.11.232 Region: Toronto 1

```

## Sample Run - destroy
```
This will destroy a SINGLE droplet by ID. Warning: this literally removes the droplet
and all system data will cease to exist.

$ ./D2Deploy -destroy 12563675
Successfully destroyed droplet: 12563675

```

## Sample Run - destroyall
```
This will destroy ALL droplets. Warning: this literally removes the droplet
and all system data will cease to exist.

Note this is useful for mass cleanup when destroying an entire network of droplets
is warranted.

$ ./D2Deploy -destroyall
Successfully destroyed all droplets

```

## Developing
```
Alpha code under active development
```

## Contact
```
# Author: Chris Patten
# Contact (Email): cpatten[t.a.]packetresearch[t.o.d]com
# Contact (Twitter): packetassailant
```
