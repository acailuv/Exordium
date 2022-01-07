CREATE TABLE IF NOT EXISTS users(
   id VARCHAR(255) PRIMARY KEY,
   created_at TIMESTAMP,
   created_by VARCHAR(255),
   updated_at TIMESTAMP,
   updated_by VARCHAR(255),
   balance INTEGER
);