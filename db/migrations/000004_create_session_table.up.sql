CREATE TABLE session (
    session_id BIGINT PRIMARY KEY,
    auth_id BIGINT NOT NULL,
    customer_id BIGINT NOT NULL,
    token VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (auth_id) REFERENCES auth(auth_id),
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id)
);
