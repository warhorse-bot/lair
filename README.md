# Lair Docker-Compose

This is a dockerized version of the Lair Framework. This version is a follow on from https://github.com/ryhanson/lair-docker and has been designed for a Docker-Compose deployment. There have been a lot a changes to how this compose works from lair-docker. The biggest change is everything is version locked and all docker images are part of WarHorse. You will not need to build images locally this is also what breaks things. What this means is that a year from now this Docker-Compose should still work as we are not useing upstream docker images and relaying on building code. Lastly this image uses Traefik which is an amazing load blancer. This should allow for lets encrypt certs among ALOT of other things.

Ngrok is SUPPORTED. Even easyter then ever to get a public URL.

This was inspired and built off the initial work done by b00stfr3ak here: https://github.com/b00stfr3ak/dockerfiles/tree/master/lair; && ryhanson https://github.com/ryhanson/

## Installation

First make sure you have docker installed. Here are OS X Beta and Linux instructions.

To install docker-compose in Linux (in case you have problem with version 1.5.2 of the repos) use the following guide. You may also need to create a symbolic link
```
$ ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

then use the following commands to get this repo:

```
$ git clone https://github.com/war-horse/lair.git
$ cd lair
```
Set lair password inside the docker-compose.yaml look for the line below and change it to what ever you want the password to be:

```
    environment:
      LAIR_PASSWORD: "password"
```

Start the docker compose:

```
$ docker-compose up -d
```

If everything worked you should see something like below

```
Creating network "dockercompose_default" with the default driver
Creating dockercompose_lairdb_1  ... done
Creating dockercompose_traefik_1 ... done
Creating dockercompose_lair_1    ... done
Creating dockercompose_lairapi_1 ... done
```

To access Lair go to the following URL

```
https://lair.localhost
```

The login should be

```
username: admin@localhost
password: password

```


## Usage
Before running any docker-compose commands, make sure you are in the docker-compose directory.

When you are ready to start up Lair again, run:

```
$ docker-compose start
```

To stop all Lair services, run:

```
$ docker-compose stop
```

## Ngrok

Ngrok is supported out of box. Ngrok is also enabled by default. You can get your public Ngrok URL by visting the following site.

```
http://ngrok.localhost
```

If you do not need OR want ngrok you can easly disable this by commenting out these lines in the docker-compose.yaml

```
  # ngrok:
  #     hostname: ngrok
  #     container_name: lair_ngrok
  #     image: wernight/ngrok
  #     restart: always
  #     labels:
  #       - "traefik.enable=true"
  #       - "traefik.port=4040"
  #       - "traefik.backend=ngrok"
  #       - "traefik.frontend.rule=Host:ngrok.localhost"
  #     entrypoint: ngrok http traefik:80 -host-header=lair.localhost
  #     links:
  #       - traefik
  #     networks:
  #       - default
  
```

## MongoDB database
If you previously used lair-docker from ryhanson https://github.com/ryhanson/ then you will already have a database and it will import this and you should not loose any data. You can even remove all containers and you will not loose any data. If you need to backup this database it will be saved at the location below

```
$ $HOME/.lair/mongodb/data/db
```


## Maintainers
Ralph May
- [@ralphte1](https://twitter.com/ralphte1)
- ralph@thedarkcloud.net

