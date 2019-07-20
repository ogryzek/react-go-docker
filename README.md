# React, Go, Heroku

Use a single Dockerfile to spin up a ReactJS client, and a Go server.  
  
## Usage
  
```
# create the image
docker build -t golang-heroku .

# run a container
docker run --detach --name full-stack -p 3000:8080 -d golang-heroku

# remove container
docker container stop full-stack
docker container rm full-stack

# delete the image
docker rmi golang-heroku
```
