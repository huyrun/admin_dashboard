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
-- Name: atlas_schema_revisions; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA atlas_schema_revisions;


ALTER SCHEMA atlas_schema_revisions OWNER TO postgres;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS '';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: atlas_schema_revisions; Type: TABLE; Schema: atlas_schema_revisions; Owner: postgres
--

CREATE TABLE atlas_schema_revisions.atlas_schema_revisions (
                                                               version character varying NOT NULL,
                                                               description character varying NOT NULL,
                                                               type bigint DEFAULT 2 NOT NULL,
                                                               applied bigint DEFAULT 0 NOT NULL,
                                                               total bigint DEFAULT 0 NOT NULL,
                                                               executed_at timestamp with time zone NOT NULL,
                                                               execution_time bigint NOT NULL,
                                                               error text,
                                                               error_stmt text,
                                                               hash character varying NOT NULL,
                                                               partial_hashes jsonb,
                                                               operator_version character varying NOT NULL
);


ALTER TABLE atlas_schema_revisions.atlas_schema_revisions OWNER TO postgres;

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
-- Data for Name: atlas_schema_revisions; Type: TABLE DATA; Schema: atlas_schema_revisions; Owner: postgres
--

COPY atlas_schema_revisions.atlas_schema_revisions (version, description, type, applied, total, executed_at, execution_time, error, error_stmt, hash, partial_hashes, operator_version) FROM stdin;
20231231122354		2	16	16	2024-04-06 21:08:28.468012+01	245750			vBod8NLafJkwZTiKalsl9/5as6ICe33ayHbS4rtAawU=	["h1:yI7NNtpXu5TgklkMmt2qaBDN2fpa/JOxIbD94x+bgZM=", "h1:PBfKdvfLxkN/i3waiAezUZc6H8TATsXKt/T/LmrYr5E=", "h1:+6grhnvzfSX4LqOpOOiYzpJ8XYCKESpzKQMj+mvyrPI=", "h1:qBvwkojBxdSE4cjuQDvR776AeAA9jBpb8dnGZgTTr+E=", "h1:xc4iy3scQUyMhgSu/Mx4Wjk0PYjFSkB68F5Vyyrm5Xg=", "h1:Grdaq2/ZBCBTAtp+W2uUHZAOIt35xaqKG8z/N+2uaro=", "h1:dnRBJJSgnCEBCu0ZI/9w9tNRZ5uqtTDmoplDcmuK1g4=", "h1:erRKIN1jz/DODnT+UtuWhwK5kv7kXMwNLGV1ZaIhAcQ=", "h1:57O510QD3am/UL5qOLLRFjXrlrd55sV4NW34Hwbgc5A=", "h1:fWw8xgFj4q87wSiV7pUcOmTT4S9rlMWonL6qVXek/uw=", "h1:lR5ggQhIirRE2fE5YN7wh4l0lHXaGdt89QH73brGW58=", "h1:bN68Yao+UQeDZtahhKZef3Q1+8M7+mLa8J3XkdYbzNU=", "h1:F1r5/MypFRomWyiWyLwr7fWRKuNzXnfJSNxmjce/rgI=", "h1:Jn/QdQKkq0KT7oKHVDCuzrnqEYdREwxzWYuDhwNwqbo=", "h1:kRHRLdkSqRs0ZUQNui33O8/PX2rjEuprJJbF6EIcPFQ=", "h1:w5y1CMzKGfyoEcS2XNxUu7/tMwhC9DAgi8mNAqaH8XU="]	Atlas CLI v0.21.2-58d847f-canary
20240610142529		2	15	15	2024-06-10 20:41:51.23829+01	173083			Ww6BYgvlF6ZlRlmkqi1lp2EYuKD9TFVUYl3Gp8xWlkk=	["h1:sn1lsj5qX4zF/VG6H7QEfY7bZJ5vwN1adZVxgC5Si0g=", "h1:HiYI3UA+1+OFYcqM3mQNL5FMVbywPEgb62QiO37ghwg=", "h1:qhTVW+RcM7HaPGTOpAVctpSzaCWUODmqUvytcLnqiLI=", "h1:PztyLuK42FM16FAFeTlq5r0Hfz6xGjwpNcP5+QKRxPE=", "h1:0zPISKk/jnCKTilstYz6hnaRwknEKd4F3EIZMk8eXko=", "h1:2t2oLw1VU6gyYc4HolkliozXSiKnQ+wXXIfQPlX8l6Y=", "h1:I0kFxBpdIqlIUJH4TaWk+M2XjiR3epqpXMsDCxd8DoY=", "h1:6YJrocSDH3aqmPkBdkr/KD5EDTTtYe2NtbadkMeUNTw=", "h1:JiDxuTLGXOonTTNlKBZYMPSNyS5ismHjBLSAa0wlFgA=", "h1:fxuI8soUmpvEPVjqc8CW13aylD9milmye2lNb/tmTvo=", "h1:WArUuLXxphse601FMCkdd60JX920l5MfRbNvWvfyD2w=", "h1:KcVTfuyWyLPXy5s0MxtF5UcukkDU2O/XaSqMIl3fj3M=", "h1:EsFlHGHygVfe7XTD3OHR0J9ciXNFmoNbJG79W1EuZNw=", "h1:n1scolv8IgQcKcPoTIa7bBt1GUJv3AVEFwvn3lPK6/k=", "h1:/HL5wfRseOD6TNriOFWAXZnUtqkpK1T6BWJclR8o0V8="]	Atlas CLI v0.23.1-7b27d81-canary
\.


--
-- Data for Name: activities; Type: TABLE DATA; Schema: public; Owner: wish
--

COPY public.activities (id, action, points, user_id, entity_id, created_at) FROM stdin;
1	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	1	2024-10-30 21:44:07.476211+00
2	wish	1	\\x0192dfa356d3b1820e93b90a9c4d217c	2	2024-10-31 00:10:02.978819+00
3	wish	1	\\x0192dfa356d3b1820e93b90a9c4d217c	3	2024-11-02 15:48:07.88525+00
4	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354569	2024-11-05 22:01:40.862004+00
5	like	1	\\x0192ded76aa9e70cd8aaf41391111946	354568	2024-11-06 18:14:47.709271+00
6	like	1	\\x0192ded76aa9e70cd8aaf41391111946	354569	2024-12-01 17:24:11.409185+00
7	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354570	2024-12-03 14:02:52.868466+00
8	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354571	2024-12-03 14:05:34.254207+00
9	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354572	2024-12-03 14:16:18.866123+00
10	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354573	2024-12-03 14:19:16.196101+00
11	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354574	2024-12-03 16:05:44.965796+00
12	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354575	2024-12-03 16:06:35.622323+00
13	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354576	2024-12-06 10:59:20.251963+00
14	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354577	2024-12-06 11:01:16.016069+00
15	like	1	\\x0192ded76aa9e70cd8aaf41391111946	354577	2024-12-14 15:20:19.639411+00
16	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354578	2024-12-15 17:32:35.941337+00
17	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354579	2024-12-16 22:44:34.880422+00
18	wish	1	\\x0192ded76aa9e70cd8aaf41391111946	354580	2024-12-18 12:30:35.807199+00
\.

--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: wish
--

COPY public.categories (id, name) FROM stdin;
1	wellness
2	home
3	family
4	travel
5	tech
6	style
7	finances
8	food
9	hobbies
10	fun
11	growth
12	motorvehicles
13	other
\.


--
-- Data for Name: entities; Type: TABLE DATA; Schema: public; Owner: wish
--

COPY public.entities (id) FROM stdin;
1
2
3
4
5
6
7
8
9
10
\.

COPY public.users (created_at, updated_at, id, username, first_name, last_name, email, role, password_hash, age, dob, sex, country, city, points, avatar_url, google_sub, fb_id, status) FROM stdin;
2024-11-02 16:23:56.955203+00	2024-11-02 16:23:56.955203+00	\\x0192edb0c11bce3fe576899a15df0d4e	Stehr1764	Eddie	Rolfson	moriahconnelly@huel.net	user		0	0001-01-01	0			0		\N	\N	1
2024-11-02 16:23:56.962637+00	2024-11-02 16:23:56.962637+00	\\x0192edb0c12230784f6fd9d9a2049b67	Cummings8902	Dion	Douglas	tremainesanford@mohr.net	user		0	0001-01-01	0			0		\N	\N	1
2024-10-30 21:49:46.568642+00	2024-10-30 21:49:47.346458+00	\\x0192df67fac890e3251477f2835d9048	chipenes305	Xander	Pokhylenko	xander.pokhylenko@gmail.com	user		0	0001-01-01	0			0	https://s3.mywish.is//users/01JBFPFYP8J3HJA53QYA1NV428/profile/avatar/	105303970959694004792	\N	1
2024-11-02 16:23:56.963374+00	2024-11-02 16:23:56.963374+00	\\x0192edb0c1234a76acd0ef47f31340cf	Tillman4560	Ricky	Aufderhar	kristinastroman@pacocha.info	user		0	0001-01-01	0			0		\N	\N	1
2024-11-02 16:23:56.964155+00	2024-11-02 16:23:56.964155+00	\\x0192edb0c1240bcf38b507830b129ea3	Cummerata7134	Kelly	Homenick	grantgerlach@toy.io	user		0	0001-01-01	0			0		\N	\N	1
2024-10-30 19:11:52.489575+00	2024-12-18 12:30:35.809811+00	\\x0192ded76aa9e70cd8aaf41391111946	thiter145	Xanderx	Pokhylenkox	pro100gt@gmail.com	user		0	0001-01-01	0			16		114408320438981253567	533518888188023	1
2024-11-02 15:53:38.008528+00	2024-11-02 15:53:38.008528+00	\\x0192ed94ffd8a6d9aea7a027752a8188	Hodkiewicz7583	Zander	Gleason	savannahermiston@veum.name	user		0	0001-01-01	0			0		\N	\N	1
2024-11-02 15:53:38.012335+00	2024-11-02 15:53:38.012335+00	\\x0192ed94ffdc33823e7539d17e649667	Wilkinson9030	Baylee	Schultz	kaylicormier@gleason.net	user		0	0001-01-01	0			0		\N	\N	1
2024-11-02 15:53:38.012687+00	2024-11-02 15:53:38.012687+00	\\x0192ed94ffdc33823e7539d23ff208f8	Lowe1629	Kelsi	Crooks	austenfunk@labadie.io	user		0	0001-01-01	0			0		\N	\N	1
2024-11-02 15:53:38.013168+00	2024-11-02 15:53:38.013168+00	\\x0192ed94ffdd972ca39070d53f931077	Dickens9932	Laverne	Towne	dasialangosh@kshlerin.info	user		0	0001-01-01	0			0		\N	\N	1
2024-11-02 15:53:38.013555+00	2024-11-02 15:53:38.013555+00	\\x0192ed94ffdd972ca39070d58db120ed	Christiansen9119	Deja	Gorczany	hellenemard@langworth.io	user		0	0001-01-01	0			0		\N	\N	1
\.

COPY public.entity_comments (created_at, updated_at, entity_id, user_id, comment_no, comment) FROM stdin;
2024-11-02 18:51:40.649433+00	2024-11-02 18:51:40.649433+00	263089	\\x0192ee32a93fffbe964114e4eac6ee30	1	Which is lastly.
\.

COPY public.liked_entities (entity_id, user_id, amount) FROM stdin;
263089	\\x0192ee32a93fffbe964114e4eac6ee30	0
263090	\\x0192ee32a93fffbe964114e4eac6ee30	9
263091	\\x0192ee32a93fffbe964114e4eac6ee30	4
263092	\\x0192ee32a93fffbe964114e4eac6ee30	1
263093	\\x0192ee32a93fffbe964114e4eac6ee30	0
263094	\\x0192ee32a93fffbe964114e4eac6ee30	9
263095	\\x0192ee32a93fffbe964114e4eac6ee30	3
263096	\\x0192ee32a93fffbe964114e4eac6ee30	3
263097	\\x0192ee32a93fffbe964114e4eac6ee30	6
263098	\\x0192ee32a93fffbe964114e4eac6ee30	2
263099	\\x0192ee32a946235938092b63addaa39f	7
263100	\\x0192ee32a946235938092b63addaa39f	9
263101	\\x0192ee32a946235938092b63addaa39f	7
263102	\\x0192ee32a946235938092b63addaa39f	6
\.

COPY public.saved_entities (entity_id, user_id, type) FROM stdin;
\.

COPY public.tags (tag_name) FROM stdin;
demo
surfing
test
tag
testing
tags
example
image
picture
photo
image test
sample
text
tag233
yacht
boat
sea
ship
luxury
cloud
computing
storage
internet
technology
sky
weather
nature
meteorology
\.

-- COPY public.user_relationships (first_user_id, second_user_id, first_to_second_status, second_to_first_status) FROM stdin;
-- \\x0192ded76aa9e70cd8aaf41391111946	\\x0192ee32b4adb467155fdaddcb095cf8 subscribed  subscribed
-- \\x0192ded76aa9e70cd8aaf41391111946	\\x0192dfa356d3b1820e93b90a9c4d217c subscribed  subscribed
-- \\x0192ded76aa9e70cd8aaf41391111946	\\x0192f9681151d19405b5f0e89e8b1333	subscribed  subscribed
-- \\x0192ded76aa9e70cd8aaf41391111946	\\x0192f9675ceff02178785fec1a469220	subscribed  subscribed
-- \\x0192ded76aa9e70cd8aaf41391111946	\\x0192f96f686f8a7fead84d6c94233956	subscribed  subscribed
-- \.

COPY public.wish_stories (created_at, updated_at, entity_id, body, image, status) FROM stdin;
2024-11-02 15:53:40.106433+00	2024-11-02 15:53:40.106433+00	4	Since Monacan elegant should when that anyone Swiss refrigerator one. Run these hourly crew us her riches were our remain. Owing mob daily what e.g. thing awful another group them. Be how abroad utterly wash the ourselves none that contradict. Nobody that smoke furthermore hundred to including school you follow. Within be either stand all there deceit which block whom. Ream bale lately which does time movement are through patrol.\nThemselves my anything before first elsewhere upon party group knit. Me exemplified after himself whose this company just e.g. these. Be wad all this been fortnightly yours those they she. Carry strongly numerous these no muddy sparse Peruvian peep batch. Horde those much ill consequently laughter team it will next. Orchard line quaint then its I flock first fortnightly till. Of it irritably horde cast from today to an without.\nYay catalog smoothly cough usually does fact yours hotel for. Board conclude eventually before group dark that out previously hand. Head they village there e.g. off half finally is them. These Confucian his while already almost goal line flock answer. Significant such begin annoyance consequently finally soon lastly so close. That themselves ours in when had either have besides these. Wow then for hourly had awfully yet talented that honestly.		published
\.

COPY public.wishes (created_at, updated_at, entity_id, type, title, description, story, price, currency, category_id, visible_by, image, user_id, status) FROM stdin;
2024-10-30 21:44:07.462979+00	2024-10-30 21:49:09.049407+00	1	wish	mk7	\N	\N	\N	\N	12	1	https://s3.mywish.is/develop/users/01JBFDETN9WW6DHAQM2E8H26A6/wishes/MDFKQkZQNFBCNlRSS1lIQzZUWFRSOTEyVjA=/preview.jpg	\\x0192ded76aa9e70cd8aaf41391111946	completed
