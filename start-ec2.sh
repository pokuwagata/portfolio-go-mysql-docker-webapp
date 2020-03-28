#!/bin/bash

# for EC2 instance
sudo chown 999:999 mysql/logs/error.log && sudo docker-compose -f docker-compose.prod.yml up -d --build
