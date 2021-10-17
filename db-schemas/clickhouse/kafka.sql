--Kafka engine that will consume messages from Kafka topic
CREATE TABLE IF NOT EXISTS taskdb.kafka (
  id Int64,
  name String,
  created_at DateTime
) ENGINE = Kafka() SETTINGS kafka_broker_list = 'glider-01.srvs.cloudkafka.com:9094',
kafka_topic_list = 'hoq0cmws-user_creation',
kafka_group_name = 'sample_group',
kafka_format = 'JSONEachRow';

--Resulting table with messages from Kafka
CREATE TABLE taskdb.kafka_data_source (
  id Int64,
  name String,
  created_at DateTime
) ENGINE = MergeTree()
ORDER BY
  id;

--View that will read Kafka engine table and store data in kafka_data_source in it's initial state
  CREATE MATERIALIZED VIEW taskdb.data_view TO taskdb.kafka_data_source AS
SELECT
  *
FROM
  taskdb.kafka;