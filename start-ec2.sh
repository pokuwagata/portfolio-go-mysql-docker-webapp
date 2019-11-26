#!/bin/bash

# for EC2 instance
sudo chown 999:999 mysql/logs/error.log && docker-compose -f docker-compose.prod.yml up --build