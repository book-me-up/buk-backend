CREATE DATABASE buk_db;

\c buk_db;

CREATE TABLE customer (
    customer_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL
);

CREATE TABLE partner (
    partner_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL
);

CREATE TABLE appointment (
    appointment_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    partner_id uuid NOT NULL,
    customer_id uuid NOT NULL,
    begin_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    duration TIMESTAMP NOT NULL,
    description TEXT NULL,
    CONSTRAINT fk_partner_id FOREIGN KEY(partner_id) REFERENCES partner(partner_id),
    CONSTRAINT fk_customer_id FOREIGN KEY(customer_id) REFERENCES customer(customer_id)
);

INSERT INTO partner (first_name, last_name, email, phone)
VALUES ("Vanessa", "Nutricionista", "vanessa@gmail.com", "Sei n");
