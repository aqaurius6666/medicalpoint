NETWORK_NAME=medicalchain-dev
docker-compose --project-name=medicalchain-dev -f deploy/dev/docker-compose.yaml up -d
docker exec -it medicalchain-dev_mainservice_1 sh
docker logs medicalchain-dev_mainservice_1 -f