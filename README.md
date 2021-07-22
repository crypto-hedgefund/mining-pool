# mining-pool

Ethereum mining pool written in Go

#### Starting Redis server from Docker

`sudo docker run --name=dev-redis --publish=6379:6379 --hostname=redis --restart=on-failure --detach redis:latest`

#### Stop Redis docker server

`sudo docker ps`
`docker stop <container id>`

####

docker ps -a (Show all containers)

docker image ls (Show all installed images)

docker rmi -f (Remove image)
