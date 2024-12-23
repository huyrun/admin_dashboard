--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Homebrew)
-- Dumped by pg_dump version 17.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
-- SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA IF NOT EXISTS public;


ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: activities; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.activities (
                                   id bigint NOT NULL,
                                   action character varying(255) NOT NULL,
                                   points smallint NOT NULL,
                                   user_id bytea NOT NULL,
                                   entity_id bigint,
                                   created_at timestamp with time zone
);


ALTER TABLE public.activities OWNER TO wish;

--
-- Name: activities_id_seq; Type: SEQUENCE; Schema: public; Owner: wish
--

CREATE SEQUENCE public.activities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.activities_id_seq OWNER TO wish;

--
-- Name: activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wish
--

ALTER SEQUENCE public.activities_id_seq OWNED BY public.activities.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.categories (
                                   id bigint NOT NULL,
                                   name character varying(64) NOT NULL
);


ALTER TABLE public.categories OWNER TO wish;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: wish
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO wish;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wish
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: entities; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.entities (
                                 id bigint NOT NULL
);


ALTER TABLE public.entities OWNER TO wish;

--
-- Name: entities_id_seq; Type: SEQUENCE; Schema: public; Owner: wish
--

CREATE SEQUENCE public.entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.entities_id_seq OWNER TO wish;

--
-- Name: entities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wish
--

ALTER SEQUENCE public.entities_id_seq OWNED BY public.entities.id;


--
-- Name: entity_comments; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.entity_comments (
                                        created_at timestamp with time zone,
                                        updated_at timestamp with time zone,
                                        entity_id bigint NOT NULL,
                                        user_id bytea NOT NULL,
                                        comment_no bigint NOT NULL,
                                        comment character varying(2048) NOT NULL
);


ALTER TABLE public.entity_comments OWNER TO wish;

--
-- Name: entity_comments_comment_no_seq; Type: SEQUENCE; Schema: public; Owner: wish
--

CREATE SEQUENCE public.entity_comments_comment_no_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.entity_comments_comment_no_seq OWNER TO wish;

--
-- Name: entity_comments_comment_no_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wish
--

ALTER SEQUENCE public.entity_comments_comment_no_seq OWNED BY public.entity_comments.comment_no;


--
-- Name: entity_tags; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.entity_tags (
                                    entity_id bigint NOT NULL,
                                    tag_id text NOT NULL
);


ALTER TABLE public.entity_tags OWNER TO wish;

--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.goose_db_version (
                                         id integer NOT NULL,
                                         version_id bigint NOT NULL,
                                         is_applied boolean NOT NULL,
                                         tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goose_db_version OWNER TO wish;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: wish
--

CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.goose_db_version_id_seq OWNER TO wish;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wish
--

ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;


--
-- Name: liked_entities; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.liked_entities (
                                       entity_id bigint NOT NULL,
                                       user_id bytea NOT NULL,
                                       amount smallint
);


ALTER TABLE public.liked_entities OWNER TO wish;

--
-- Name: saved_entities; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.saved_entities (
                                       entity_id bigint NOT NULL,
                                       user_id bytea NOT NULL,
                                       type text DEFAULT 'wish'::text NOT NULL
);


ALTER TABLE public.saved_entities OWNER TO wish;

--
-- Name: tags; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.tags (
                             tag_name text NOT NULL
);


ALTER TABLE public.tags OWNER TO wish;

--
-- Name: user_relationships; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.user_relationships (
                                           first_user_id bytea NOT NULL,
                                           second_user_id bytea NOT NULL,
                                           first_to_second_status text,
                                           second_to_first_status text,
                                           are_friends boolean GENERATED ALWAYS AS (((first_to_second_status = 'subscribed'::text) AND (second_to_first_status = 'subscribed'::text))) STORED
);


ALTER TABLE public.user_relationships OWNER TO wish;

--
-- Name: users; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.users (
                              created_at timestamp with time zone,
                              updated_at timestamp with time zone,
                              id bytea NOT NULL,
                              username character varying(30) NOT NULL,
                              first_name character varying(255),
                              last_name character varying(255),
                              email character varying(100) NOT NULL,
                              role character varying(10) DEFAULT 'user'::character varying NOT NULL,
                              password_hash text,
                              age smallint,
                              dob date,
                              sex bigint DEFAULT 0,
                              country character varying(2),
                              city character varying(50),
                              points smallint DEFAULT 0,
                              avatar_url character varying(1000),
                              google_sub character varying(24),
                              fb_id character varying(24),
                              status smallint DEFAULT 0
);


ALTER TABLE public.users OWNER TO wish;

--
-- Name: wish_stories; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.wish_stories (
                                     created_at timestamp with time zone,
                                     updated_at timestamp with time zone,
                                     entity_id bigint NOT NULL,
                                     body character varying(10000) NOT NULL,
                                     image character varying(256),
                                     status character varying(12) DEFAULT 'draft'::character varying
);


ALTER TABLE public.wish_stories OWNER TO wish;

--
-- Name: wish_stories_entity_id_seq; Type: SEQUENCE; Schema: public; Owner: wish
--

CREATE SEQUENCE public.wish_stories_entity_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wish_stories_entity_id_seq OWNER TO wish;

--
-- Name: wish_stories_entity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: wish
--

ALTER SEQUENCE public.wish_stories_entity_id_seq OWNED BY public.wish_stories.entity_id;


--
-- Name: wishes; Type: TABLE; Schema: public; Owner: wish
--

CREATE TABLE public.wishes (
                               created_at timestamp with time zone,
                               updated_at timestamp with time zone,
                               entity_id bigint NOT NULL,
                               type character varying(255) NOT NULL,
                               title character varying(255) NOT NULL,
                               description character varying(2048) DEFAULT NULL::character varying,
                               story character varying(10000) DEFAULT NULL::character varying,
                               price bigint,
                               currency character varying(10) DEFAULT NULL::character varying,
                               category_id bigint,
                               visible_by bigint DEFAULT 1,
                               image character varying(256),
                               user_id bytea NOT NULL,
                               status character varying(12) DEFAULT 'new'::character varying
);


ALTER TABLE public.wishes OWNER TO wish;

--
-- Name: activities id; Type: DEFAULT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.activities ALTER COLUMN id SET DEFAULT nextval('public.activities_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: entities id; Type: DEFAULT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entities ALTER COLUMN id SET DEFAULT nextval('public.entities_id_seq'::regclass);


--
-- Name: entity_comments comment_no; Type: DEFAULT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_comments ALTER COLUMN comment_no SET DEFAULT nextval('public.entity_comments_comment_no_seq'::regclass);


--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);


--
-- Name: wish_stories entity_id; Type: DEFAULT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wish_stories ALTER COLUMN entity_id SET DEFAULT nextval('public.wish_stories_entity_id_seq'::regclass);


--
-- Name: activities activities_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: entities entities_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entities
    ADD CONSTRAINT entities_pkey PRIMARY KEY (id);


--
-- Name: entity_comments entity_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_comments
    ADD CONSTRAINT entity_comments_pkey PRIMARY KEY (entity_id, comment_no);


--
-- Name: entity_tags entity_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_tags
    ADD CONSTRAINT entity_tags_pkey PRIMARY KEY (entity_id, tag_id);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: liked_entities liked_entities_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.liked_entities
    ADD CONSTRAINT liked_entities_pkey PRIMARY KEY (entity_id, user_id);


--
-- Name: saved_entities saved_entities_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.saved_entities
    ADD CONSTRAINT saved_entities_pkey PRIMARY KEY (entity_id, user_id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (tag_name);


--
-- Name: categories uni_categories_name; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT uni_categories_name UNIQUE (name);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: users uni_users_username; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_username UNIQUE (username);


--
-- Name: user_relationships user_relationships_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.user_relationships
    ADD CONSTRAINT user_relationships_pkey PRIMARY KEY (first_user_id, second_user_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wish_stories wish_stories_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wish_stories
    ADD CONSTRAINT wish_stories_pkey PRIMARY KEY (entity_id);


--
-- Name: wishes wishes_pkey; Type: CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wishes
    ADD CONSTRAINT wishes_pkey PRIMARY KEY (entity_id, user_id);


--
-- Name: idx_user_relationships_first_to_second_status; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_user_relationships_first_to_second_status ON public.user_relationships USING btree (first_to_second_status);


--
-- Name: idx_user_relationships_second_to_first_status; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_user_relationships_second_to_first_status ON public.user_relationships USING btree (second_to_first_status);


--
-- Name: idx_users_username; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_users_username ON public.users USING btree (username);


--
-- Name: idx_wishes_created_at; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_wishes_created_at ON public.wishes USING btree (created_at);


--
-- Name: idx_wishes_status; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_wishes_status ON public.wishes USING btree (status);


--
-- Name: idx_wishes_title; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_wishes_title ON public.wishes USING btree (title);


--
-- Name: idx_wishes_visible_by; Type: INDEX; Schema: public; Owner: wish
--

CREATE INDEX idx_wishes_visible_by ON public.wishes USING btree (visible_by);


--
-- Name: activities fk_activities_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.activities
    ADD CONSTRAINT fk_activities_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: activities fk_activities_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.activities
    ADD CONSTRAINT fk_activities_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: entity_comments fk_entity_comments_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_comments
    ADD CONSTRAINT fk_entity_comments_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: entity_comments fk_entity_comments_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_comments
    ADD CONSTRAINT fk_entity_comments_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: entity_tags fk_entity_tags_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_tags
    ADD CONSTRAINT fk_entity_tags_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id) ON DELETE CASCADE;


--
-- Name: entity_tags fk_entity_tags_tag; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.entity_tags
    ADD CONSTRAINT fk_entity_tags_tag FOREIGN KEY (tag_id) REFERENCES public.tags(tag_name) ON DELETE CASCADE;


--
-- Name: liked_entities fk_liked_entities_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.liked_entities
    ADD CONSTRAINT fk_liked_entities_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id);


--
-- Name: liked_entities fk_liked_entities_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.liked_entities
    ADD CONSTRAINT fk_liked_entities_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: saved_entities fk_saved_entities_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.saved_entities
    ADD CONSTRAINT fk_saved_entities_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id);


--
-- Name: saved_entities fk_saved_entities_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.saved_entities
    ADD CONSTRAINT fk_saved_entities_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: user_relationships fk_user_relationships_first_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.user_relationships
    ADD CONSTRAINT fk_user_relationships_first_user FOREIGN KEY (first_user_id) REFERENCES public.users(id);


--
-- Name: user_relationships fk_user_relationships_second_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.user_relationships
    ADD CONSTRAINT fk_user_relationships_second_user FOREIGN KEY (second_user_id) REFERENCES public.users(id);


--
-- Name: wish_stories fk_wish_stories_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wish_stories
    ADD CONSTRAINT fk_wish_stories_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id);


--
-- Name: wishes fk_wishes_category; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wishes
    ADD CONSTRAINT fk_wishes_category FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: wishes fk_wishes_entity; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wishes
    ADD CONSTRAINT fk_wishes_entity FOREIGN KEY (entity_id) REFERENCES public.entities(id);


--
-- Name: wishes fk_wishes_user; Type: FK CONSTRAINT; Schema: public; Owner: wish
--

ALTER TABLE ONLY public.wishes
    ADD CONSTRAINT fk_wishes_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO wish;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

