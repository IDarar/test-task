# test-task

Docker image can be built with ```make dockerbuild```
All app's dependencies are in docker compose. And also to run app you need to create table in postgres. 
Kafka cluster is from cloud provier. It's configs are in this repository.
To store messages from kafka to ClickHouse you need to add ClickHouse's configs files from repository to your ClickHouse instance and create needed tables.

App is running on my vps and rpc calls are available at subjless.space:8080
