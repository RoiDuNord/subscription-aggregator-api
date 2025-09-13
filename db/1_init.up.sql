CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    service_name VARCHAR(50) NOT NULL,
    price INT NOT NULL,
    start_date DATE
);

CREATE INDEX idx_subscription_price 
ON subscriptions(price);