docker rmi localdemo-web && \
mv config/example.env config/.env && \
cd deployments/docker/localdemo && \
docker-compose -f docker-compose.yaml up -d && \
cd ../../../ && \
mv config/.env config/example.env