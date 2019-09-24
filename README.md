# LIFE CAFE BACKEND

## How to run

### Requirement
- Install golang
- Install docker

### Config environment
Follow command
- ```cat .env.example > .env```
- ```cat .env_migrator.yaml.example > .env_migrator.yaml```

### Setup db and adminer
Follow command
- ```docker-compose up -f docker-compose-local.yaml -d```

### Run server 
Follow command
- ```make dev```

### Run test
Flolow command
-```make test```

