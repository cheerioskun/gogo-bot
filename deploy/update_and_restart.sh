#!/bin/bash

# Stop further execution if any error occurs
set -e
cd ~/projects/gogo-bot

#update repo
git pull
# Stop and remove latest container & image
sudo docker stop gogo-bot-container
sudo docker rm gogo-bot-container
sudo docker rmi gogo-bot-image
sudo docker ps
# Create new image and container
sudo docker build -t gogo-bot-image .
sudo docker run -d --name gogo-bot-container -e TZ=Asia/Kolkata gogo-bot-image
sudo docker ps