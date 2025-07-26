# 🚀 logger2

An AWS deployment with 2 microservices that talk to eachother, for the sake of learning terraform.

## ⚙️ How it works

The repository consists of two (micro)services:

- a database service which exposes an endpoint that can be used to read/write
- a logger service, constantly writing to the database via the endpoint

## 🌐 What's cool about this

The deployment to AWS is done via Terraform files, meaning no AWS UI whatsoever.

We can re-deploy or modify the entire infrastructure by just modifying our file and running one command.

## 🛠️ Tools used:

- Go
- Docker
- Terraform
- AWS
