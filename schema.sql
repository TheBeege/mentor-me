--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: main; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA main;


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = main, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: mentor_topic; Type: TABLE; Schema: main; Owner: -; Tablespace: 
--

CREATE TABLE mentor_topic (
    user_id integer NOT NULL,
    topic_id integer NOT NULL,
    level integer,
    description text
);


--
-- Name: topic; Type: TABLE; Schema: main; Owner: -; Tablespace: 
--

CREATE TABLE topic (
    id integer NOT NULL,
    name text NOT NULL
);


--
-- Name: topic_id_seq; Type: SEQUENCE; Schema: main; Owner: -
--

CREATE SEQUENCE topic_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: topic_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: -
--

ALTER SEQUENCE topic_id_seq OWNED BY topic.id;


--
-- Name: user; Type: TABLE; Schema: main; Owner: -; Tablespace: 
--

CREATE TABLE "user" (
    id integer NOT NULL,
    username text,
    display_name text,
    email text,
    created timestamp without time zone DEFAULT now(),
    last_activity timestamp without time zone,
    password text,
    description text,
    icon_url text
);


--
-- Name: user_user_id_seq; Type: SEQUENCE; Schema: main; Owner: -
--

CREATE SEQUENCE user_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: -
--

ALTER SEQUENCE user_user_id_seq OWNED BY "user".id;


--
-- Name: id; Type: DEFAULT; Schema: main; Owner: -
--

ALTER TABLE ONLY topic ALTER COLUMN id SET DEFAULT nextval('topic_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: main; Owner: -
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_user_id_seq'::regclass);


--
-- Name: mentor_topic_id; Type: CONSTRAINT; Schema: main; Owner: -; Tablespace: 
--

ALTER TABLE ONLY mentor_topic
    ADD CONSTRAINT mentor_topic_id PRIMARY KEY (user_id, topic_id);


--
-- Name: topic_id; Type: CONSTRAINT; Schema: main; Owner: -; Tablespace: 
--

ALTER TABLE ONLY topic
    ADD CONSTRAINT topic_id PRIMARY KEY (id);


--
-- Name: topic_name; Type: CONSTRAINT; Schema: main; Owner: -; Tablespace: 
--

ALTER TABLE ONLY topic
    ADD CONSTRAINT topic_name UNIQUE (name);


--
-- Name: user_pk; Type: CONSTRAINT; Schema: main; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: user_unq_email; Type: CONSTRAINT; Schema: main; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_unq_email UNIQUE (email);


--
-- Name: user_unq_username; Type: CONSTRAINT; Schema: main; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_unq_username UNIQUE (username);


--
-- Name: mentor_topic_topic_id; Type: FK CONSTRAINT; Schema: main; Owner: -
--

ALTER TABLE ONLY mentor_topic
    ADD CONSTRAINT mentor_topic_topic_id FOREIGN KEY (topic_id) REFERENCES topic(id);


--
-- Name: mentor_topic_user_id; Type: FK CONSTRAINT; Schema: main; Owner: -
--

ALTER TABLE ONLY mentor_topic
    ADD CONSTRAINT mentor_topic_user_id FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- PostgreSQL database dump complete
--

