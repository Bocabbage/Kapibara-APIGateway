mv config/example.env config/.env && \
cd deployments/docker/localdemo && \
docker-compose -f docker-compose.yaml up -d && \
mv config/.env config/example.env