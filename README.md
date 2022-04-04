# tcms
Telegram client automation system.

Package provides gRPC interface to store automations and launch automations.
## Preparation
System support GoLang 1.17 or higher.

Automation data stores in MongoDB

System recieves telegram events from Kafka. To produce events to Kafka use [tcms telegram bridge](https://github.com/BlenderistDev/telegram-bridge).
## Installation
````
git clone https://github.com/BlenderistDev/tcms
````
## Launch
````
make build
./bin/tcms
````
