SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

CREATE TABLE launchpads (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    full_name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE destinations (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE bookings (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    gender character varying NOT NULL,
    birthday date NOT NULL,
    launchpad_id uuid NOT NULL,
    destination_id uuid NOT NULL,
    launch_date date NOT NULL,
    deleted boolean DEFAULT false, 
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE launchpad_schedule (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    launchpad_id uuid NOT NULL,
    day_of_week text CHECK (day_of_week IN ('Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday')) NOT NULL,
    destination_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

ALTER TABLE ONLY launchpads
    ADD CONSTRAINT launchpads_pkey PRIMARY KEY (id);

ALTER TABLE ONLY destinations
    ADD CONSTRAINT destinations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);

ALTER TABLE ONLY launchpad_schedule
    ADD CONSTRAINT launchpad_schedule_pkey PRIMARY KEY (id);

INSERT INTO launchpads(id, full_name, created_at, updated_at) VALUES
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Vandenberg Space Force Base Space Launch Complex 3W', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Cape Canaveral Space Force Station Space Launch Complex 40', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'SpaceX South Texas Launch Site', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Kwajalein Atoll Omelek Island', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Vandenberg Space Force Base Space Launch Complex 4E', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Kennedy Space Center Historic Launch Complex 39A', NOW(), NOW());

INSERT INTO destinations(id, name, created_at, updated_at) VALUES
    ('466fc378-14eb-4ed9-8bec-d29abe54c5a9', 'Moon', NOW(), NOW()),
    ('f47eef79-675f-46da-86f9-ee598185d204', 'Mars', NOW(), NOW()),
    ('fbd40165-03c7-47a5-be72-c79f81ebbf67', 'Pluto', NOW(), NOW()),
    ('13b91e0c-cdb4-4108-9c48-5a49d8ded732', 'Asteroid Belt', NOW(), NOW()),
    ('998f4a82-5a1c-4542-8497-e3fa24618d79', 'Europa', NOW(), NOW()),
    ('12549fca-d086-4e9f-b14e-dcb3b0d09c63', 'Titan', NOW(), NOW()),
    ('3840d5ce-b939-4af7-9dd8-ac12c09d1493', 'Ganymede', NOW(), NOW());

INSERT INTO bookings(first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date, created_at, updated_at) VALUES (
    'Brian', 'Blessed', 'Male', '1936-10-09', 'd95c83bb-be3f-4bdb-93fe-77015d95f759', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', '2021-12-01', NOW(), NOW()
);

INSERT INTO launchpad_schedule(launchpad_id, day_of_week, destination_id, created_at, updated_at) VALUES
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Sunday', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', NOW(), NOW()),
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Monday', 'f47eef79-675f-46da-86f9-ee598185d204', NOW(), NOW()),
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Tuesday', 'fbd40165-03c7-47a5-be72-c79f81ebbf67', NOW(), NOW()),
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Wednesday', '13b91e0c-cdb4-4108-9c48-5a49d8ded732', NOW(), NOW()),
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Thursday', '998f4a82-5a1c-4542-8497-e3fa24618d79', NOW(), NOW()),
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Friday', '12549fca-d086-4e9f-b14e-dcb3b0d09c63', NOW(), NOW()),
    ('d95c83bb-be3f-4bdb-93fe-77015d95f759', 'Saturday', '3840d5ce-b939-4af7-9dd8-ac12c09d1493', NOW(), NOW()),

    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Monday', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Tuesday', 'f47eef79-675f-46da-86f9-ee598185d204', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Wednesday', 'fbd40165-03c7-47a5-be72-c79f81ebbf67', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Thursday', '13b91e0c-cdb4-4108-9c48-5a49d8ded732', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Friday', '998f4a82-5a1c-4542-8497-e3fa24618d79', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Saturday', '12549fca-d086-4e9f-b14e-dcb3b0d09c63', NOW(), NOW()),
    ('b542c0cf-7fe3-4bb1-a63f-7cbdf8359975', 'Sunday', '3840d5ce-b939-4af7-9dd8-ac12c09d1493', NOW(), NOW()),

    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Tuesday', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Wednesday', 'f47eef79-675f-46da-86f9-ee598185d204', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Thursday', 'fbd40165-03c7-47a5-be72-c79f81ebbf67', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Friday', '13b91e0c-cdb4-4108-9c48-5a49d8ded732', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Saturday', '998f4a82-5a1c-4542-8497-e3fa24618d79', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Sunday', '12549fca-d086-4e9f-b14e-dcb3b0d09c63', NOW(), NOW()),
    ('b09e0b80-51ca-44ac-820a-d5b95b209cad', 'Monday', '3840d5ce-b939-4af7-9dd8-ac12c09d1493', NOW(), NOW()),

    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Wednesday', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Thursday', 'f47eef79-675f-46da-86f9-ee598185d204', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Friday', 'fbd40165-03c7-47a5-be72-c79f81ebbf67', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Saturday', '13b91e0c-cdb4-4108-9c48-5a49d8ded732', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Sunday', '998f4a82-5a1c-4542-8497-e3fa24618d79', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Monday', '12549fca-d086-4e9f-b14e-dcb3b0d09c63', NOW(), NOW()),
    ('e169113a-ae89-4c39-9a28-3cbc1c96e5e0', 'Tuesday', '3840d5ce-b939-4af7-9dd8-ac12c09d1493', NOW(), NOW()),

    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Thursday', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Friday', 'f47eef79-675f-46da-86f9-ee598185d204', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Saturday', 'fbd40165-03c7-47a5-be72-c79f81ebbf67', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Sunday', '13b91e0c-cdb4-4108-9c48-5a49d8ded732', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Monday', '998f4a82-5a1c-4542-8497-e3fa24618d79', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Tuesday', '12549fca-d086-4e9f-b14e-dcb3b0d09c63', NOW(), NOW()),
    ('9f8cb517-ca3b-4810-baef-80b48b8cf5e6', 'Wednesday', '3840d5ce-b939-4af7-9dd8-ac12c09d1493', NOW(), NOW()),

    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Friday', '466fc378-14eb-4ed9-8bec-d29abe54c5a9', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Saturday', 'f47eef79-675f-46da-86f9-ee598185d204', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Sunday', 'fbd40165-03c7-47a5-be72-c79f81ebbf67', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Monday', '13b91e0c-cdb4-4108-9c48-5a49d8ded732', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Tuesday', '998f4a82-5a1c-4542-8497-e3fa24618d79', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Wednesday', '12549fca-d086-4e9f-b14e-dcb3b0d09c63', NOW(), NOW()),
    ('4079f070-3e58-4e61-8af7-05c8de8e1fbf', 'Thursday', '3840d5ce-b939-4af7-9dd8-ac12c09d1493', NOW(), NOW());