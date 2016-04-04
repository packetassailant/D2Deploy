# D2Deploy

## Objective
```
A mass Digital Ocean droplet deployment utility
```

## Usage
```
$ ./D2Deploy -help
Usage of ./D2Deploy:
  -config string
    	Provide an optional YAML configuration file
  -deploy int
    	The number of DO droplets to deploy
  -destroy string
    	The droplet slug to destroy (use existing for names)
  -destroyall
    	Destroy all droplet slugs (teardown all)
  -existing
    	Retrieves currently deployed DO droplets
  -images
    	Get all available deployable droplet images
  -limit
    	The DO droplet limit for this account
  -name string
    	The name to prepend each droplet
  -regions
    	The list of deployable DO geo-regions
  -size string
    	Set the size of the droplet (default: 512MB) (default "512mb")
  -sizes
    	Get all available slug sizes and rates
  -slug string
    	The name of a deployable image slug (use images for slug types)
  -ssh string
    	The ssh key fingerprint for droplet auth
  -token string
    	The digital ocean oauth api token

```

## Installation
```
Installation
---------------------------------------------------
1. Install GO (tested on 1.5.2)
2 .Git clone this repo (git clone https://github.com/packetassailant/D2Deploy.git)
3. cd into the repo directory and type "go get" (this will install dependencies)
4. Type "go build" (you will now have a D2Deploy binary)
```

## Command Line Flags  

### Sample Run - limit
```
This will provide the number of droplet instances associated with the
DO account.

$ ./D2Deploy -token DigOceanAPIKey -limit
The VPS deployment limit is: 25

```

### Sample Run - regions
```
This will provide the global geographic regions available for mass droplet
deployment. (Note this will be variant depending on regional network/system saturation)

$ ./D2Deploy -token DigOceanAPIKey -regions
The available VPS deployment regions are: nyc1,sfo1,nyc2,ams2,sgp1,lon1,nyc3,ams3,fra1,tor1

```

### Sample Run - deploy
```
This will perform the actual mass deployment of specific droplet instances. The
"ssh" option is required to do public key auth with every deployed system. Further,
the ssh fingerprint can be obtained via the 'ssh-keygen -lf ~/.ssh/id_rsa.pub | cut -d " " -f 2'
command.

The following command deploys 10 droplets with the name "packetresearch" appended,
while automatically installing the ssh public key at time of build.

$ ./D2Deploy -token DigOceanAPIKey -deploy 10 -name "packetresearch" -ssh "aa:bb:cc:dd:ee:ff:gg:hh:ii:jj:kk:ll:mm:nn:oo:pp"
Successfully created all droplets

```

### Sample Run - existing
```
This will provide details about the the droplets that are currently deployed.
Information includes the droplet ID, the droplet name, the droplet IP, and the region.

$ ./D2Deploy -token DigOceanAPIKey -existing
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

### Sample Run - destroy
```
This will destroy a SINGLE droplet by ID. Warning: this literally removes the droplet
and all system data will cease to exist.

$ ./D2Deploy -token DigOceanAPIKey -destroy 12563675
Successfully destroyed droplet: 12563675

```

### Sample Run - destroyall
```
This will destroy ALL droplets. Warning: this literally removes the droplet
and all system data will cease to exist.

Note this is useful for mass cleanup when destroying an entire network of droplets
is warranted.

$ ./D2Deploy -token DigOceanAPIKey -destroyall
Successfully destroyed all droplets

```

### Sample Run - images
```
This will get all images available include both distribution and application images.

$ ./D2Deploy -token DigOceanAPIKey -images
The following images are available for distribution:
Name: 991.2.0 (beta) Distribution: CoreOS Slug: coreos-beta
Name: 5.10 x64 Distribution: CentOS Slug: centos-5-8-x64
Name: 5.10 x32 Distribution: CentOS Slug: centos-5-8-x32
Name: 10.1 Distribution: FreeBSD Slug: freebsd-10-1-x64
Name: 22 x64 Distribution: Fedora Slug: fedora-22-x64
Name: 10.2 Distribution: FreeBSD Slug: freebsd-10-2-x64
Name: 23 x64 Distribution: Fedora Slug: fedora-23-x64
Name: 6.7 x32 Distribution: CentOS Slug: centos-6-5-x32
Name: 6.7 x64 Distribution: CentOS Slug: centos-6-5-x64
Name: 8.3 x64 Distribution: Debian Slug: debian-8-x64
Name: 8.3 x32 Distribution: Debian Slug: debian-8-x32
Name: 15.10 x64 Distribution: Ubuntu Slug: ubuntu-15-10-x64
Name: 15.10 x32 Distribution: Ubuntu Slug: ubuntu-15-10-x32
Name: 12.04.5 x32 Distribution: Ubuntu Slug: ubuntu-12-04-x32
Name: 12.04.5 x64 Distribution: Ubuntu Slug: ubuntu-12-04-x64
Name: 14.04.4 x32 Distribution: Ubuntu Slug: ubuntu-14-04-x32
Name: 14.04.4 x64 Distribution: Ubuntu Slug: ubuntu-14-04-x64
Name: 7.2 x64 Distribution: CentOS Slug: centos-7-0-x64
Name: cloudbench-ycsb-on-1404.030216-1 Distribution: Ubuntu Slug: Unavailable
Name: cloudbench-hibench-on-1404.030216-1 Distribution: Ubuntu Slug: Unavailable
Name: cloudbench-nullworkload-on-1404.030316-6 Distribution: Ubuntu Slug: Unavailable
Name: 899.13.0 (stable) Distribution: CoreOS Slug: coreos-stable
Name: 1000.0.0 (alpha) Distribution: CoreOS Slug: coreos-alpha
Name: 7.10 x32 Distribution: Debian Slug: debian-7-0-x32
Name: 7.10 x64 Distribution: Debian Slug: debian-7-0-x64
Name: MEAN on 14.04 Distribution: Ubuntu Slug: mean
Name: Elixir on 14.04 Distribution: Ubuntu Slug: elixir
Name: Drone on 14.04 Distribution: Ubuntu Slug: drone
Name: PHPMyAdmin on 14.04 Distribution: Ubuntu Slug: phpmyadmin
Name: Discourse on 14.04 Distribution: Ubuntu Slug: discourse
Name: Ghost 0.7.8 on 14.04 Distribution: Ubuntu Slug: ghost
Name: ELK Logging Stack on 14.04 Distribution: Ubuntu Slug: elk
Name: Mumble Server (murmur) on 14.04 Distribution: Ubuntu Slug: mumble
Name: Ruby on Rails on 14.04 (Postgres, Nginx, Unicorn) Distribution: Ubuntu Slug: ruby-on-rails
Name: Django on 14.04 Distribution: Ubuntu Slug: django
Name: Drupal 8.0.5 on 14.04 Distribution: Ubuntu Slug: drupal
Name: ownCloud 9.0.0 on 14.04 Distribution: Ubuntu Slug: owncloud
Name: Docker 1.10.3 on 14.04 Distribution: Ubuntu Slug: docker
Name: Redis 3.0.7 on 14.04 Distribution: Ubuntu Slug: redis
Name: node v4.4.0 on 14.04 Distribution: Ubuntu Slug: node
Name: Cassandra on 14.04 Distribution: Ubuntu Slug: cassandra
Name: MongoDB 3.2.4 on 14.04 Distribution: Ubuntu Slug: mongodb
Name: Joomla! 3.5.0 on 14.04 Distribution: Ubuntu Slug: joomla
Name: GitLab 8.6.0 CE on 14.04 Distribution: Ubuntu Slug: gitlab
Name: MediaWiki on 14.04 Distribution: Ubuntu Slug: mediawiki
Name: WordPress on 14.04 Distribution: Ubuntu Slug: wordpress
Name: LEMP on 14.04 Distribution: Ubuntu Slug: lemp
Name: Dokku v0.5.3 on 14.04 Distribution: Ubuntu Slug: dokku
Name: LAMP on 14.04 Distribution: Ubuntu Slug: lamp
Name: Magento 2.0.4 on 14.04 Distribution: Ubuntu Slug: magento
Name: Redmine on 14.04 Distribution: Ubuntu Slug: redmine
Name: Primary_Image Distribution: Ubuntu Slug: Unavailable

```

### Sample Run - sizes
```
This will get all available deployable droplet system requirements including.:
1. Slug name (used for manually setting size via the "size" flag)
2. Memory allocation
3. Number of VCPUs
4. Disk size in GBs
5. Allowable network Transfer
6. Hourly rate per droplet

$ ./D2Deploy -token DigOceanAPIKey -sizes
The following sizes are available (Cost is hourly):
Slug: 512mb Memory: 512 VCPU: 1 Disk: 20 Transfer: 1 Cost: 0.00744
Slug: 1gb Memory: 1024 VCPU: 1 Disk: 30 Transfer: 2 Cost: 0.01488
Slug: 2gb Memory: 2048 VCPU: 2 Disk: 40 Transfer: 3 Cost: 0.02976
Slug: 4gb Memory: 4096 VCPU: 2 Disk: 60 Transfer: 4 Cost: 0.05952
Slug: 8gb Memory: 8192 VCPU: 4 Disk: 80 Transfer: 5 Cost: 0.11905
Slug: 16gb Memory: 16384 VCPU: 8 Disk: 160 Transfer: 6 Cost: 0.2381
Slug: 32gb Memory: 32768 VCPU: 12 Disk: 320 Transfer: 7 Cost: 0.47619
Slug: 48gb Memory: 49152 VCPU: 16 Disk: 480 Transfer: 8 Cost: 0.71429
Slug: 64gb Memory: 65536 VCPU: 20 Disk: 640 Transfer: 9 Cost: 0.95238

```


## Optional - Configuration File  

### d2d.yml.sample
```
Copy the file to something like d2d.yml and edit the following key/values.
Note that examples are provided to help determine sane values. The command line
options will be useful in determining values to populate in this file such as
slug_name and slug_size.

# The DO OAuth token
token: <DO OAuth API Key>

# The number of droplets to deploy
# Example: deploy_num: 5
deploy_num: <number of deployments>

# The name to be prepend to each droplet
# Example: droplet_name: packetresearch
droplet_name: test

# The name of the droplet slug
# Example: slug_name: docker
slug_name: docker

# The size of the droplet slug (512mb is the default)
# Example: slug_size: 512mb
slug_size: 512mb

# The ssh public key fingerprint
# Example: ssh_id: "aa:bb:cc:dd:ee:ff:gg:hh:ii:jj:kk:ll:mm:nn:oo:pp"
ssh_id: "aa:bb:cc:dd:ee:ff:gg:hh:ii:jj:kk:ll:mm:nn:oo:pp"

# The location of the yaml configuration file containing droplet build scripts
# Example: userdata_file: "configs/user-data.yml"
userdata_file: "configs/user-data.yml"

```

### user-data.yml.sample
```
This is an optional file that can be used to populate build scripts that should
automatically be included during droplet creation. The following is a very basic
example but demonstrates the ability. Consider anything that can be performed with
Shell scripting to be applicable in this file.

#!/bin/bash

apt-get -y update
apt-get -y install sslscan

```

### Sample Run - config
```
Configure at least the d2d.yml file and pass it using the "config" flag.

$ ./D2Deploy -config "configs/d2d.yml"
Successfully created all droplets
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
