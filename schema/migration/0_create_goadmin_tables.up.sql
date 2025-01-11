--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_menu_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_menu_myid_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: goadmin_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_menu (
                                     id integer DEFAULT nextval('goadmin_menu_myid_seq'::regclass) NOT NULL,
                                     parent_id integer DEFAULT 0 NOT NULL,
                                     type integer DEFAULT 0,
                                     "order" integer DEFAULT 0 NOT NULL,
                                     title character varying(50) NOT NULL,
                                     header character varying(100),
                                     plugin_name character varying(100) NOT NULL,
                                     icon character varying(50) NOT NULL,
                                     uri character varying(3000) NOT NULL,
                                     uuid character varying(100),
                                     created_at timestamp without time zone DEFAULT now(),
                                     updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_menu OWNER TO postgres;

--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_operation_log_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_operation_log_myid_seq OWNER TO postgres;

--
-- Name: goadmin_operation_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_operation_log (
                                              id integer DEFAULT nextval('goadmin_operation_log_myid_seq'::regclass) NOT NULL,
                                              user_id integer NOT NULL,
                                              path character varying(255) NOT NULL,
                                              method character varying(10) NOT NULL,
                                              ip character varying(15) NOT NULL,
                                              input text NOT NULL,
                                              created_at timestamp without time zone DEFAULT now(),
                                              updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_operation_log OWNER TO postgres;

--
-- Name: goadmin_site_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_site_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_site_myid_seq OWNER TO postgres;

--
-- Name: goadmin_site; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_site (
                                     id integer DEFAULT nextval('goadmin_site_myid_seq'::regclass) NOT NULL,
                                     key character varying(100) NOT NULL,
                                     value text NOT NULL,
                                     type integer DEFAULT 0,
                                     description character varying(3000),
                                     state integer DEFAULT 0,
                                     created_at timestamp without time zone DEFAULT now(),
                                     updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_site OWNER TO postgres;

--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_permissions_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_permissions_myid_seq OWNER TO postgres;

--
-- Name: goadmin_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_permissions (
                                            id integer DEFAULT nextval('goadmin_permissions_myid_seq'::regclass) NOT NULL,
                                            name character varying(50) NOT NULL,
                                            slug character varying(50) NOT NULL,
                                            http_method character varying(255),
                                            http_path text NOT NULL,
                                            created_at timestamp without time zone DEFAULT now(),
                                            updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_permissions OWNER TO postgres;

--
-- Name: goadmin_role_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_role_menu (
                                          role_id integer NOT NULL,
                                          menu_id integer NOT NULL,
                                          created_at timestamp without time zone DEFAULT now(),
                                          updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_role_menu OWNER TO postgres;

--
-- Name: goadmin_role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_role_permissions (
                                                 role_id integer NOT NULL,
                                                 permission_id integer NOT NULL,
                                                 created_at timestamp without time zone DEFAULT now(),
                                                 updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_role_permissions OWNER TO postgres;

--
-- Name: goadmin_role_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_role_users (
                                           role_id integer NOT NULL,
                                           user_id integer NOT NULL,
                                           created_at timestamp without time zone DEFAULT now(),
                                           updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_role_users OWNER TO postgres;

--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_roles_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_roles_myid_seq OWNER TO postgres;

--
-- Name: goadmin_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_roles (
                                      id integer DEFAULT nextval('goadmin_roles_myid_seq'::regclass) NOT NULL,
                                      name character varying NOT NULL,
                                      slug character varying NOT NULL,
                                      created_at timestamp without time zone DEFAULT now(),
                                      updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_roles OWNER TO postgres;

--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_session_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_session_myid_seq OWNER TO postgres;

--
-- Name: goadmin_session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_session (
                                        id integer DEFAULT nextval('goadmin_session_myid_seq'::regclass) NOT NULL,
                                        sid character varying(50) NOT NULL,
                                        "values" character varying(3000) NOT NULL,
                                        created_at timestamp without time zone DEFAULT now(),
                                        updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_session OWNER TO postgres;

--
-- Name: goadmin_user_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_user_permissions (
                                                 user_id integer NOT NULL,
                                                 permission_id integer NOT NULL,
                                                 created_at timestamp without time zone DEFAULT now(),
                                                 updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_user_permissions OWNER TO postgres;

--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE goadmin_users_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE goadmin_users_myid_seq OWNER TO postgres;

--
-- Name: goadmin_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE goadmin_users (
                                      id integer DEFAULT nextval('goadmin_users_myid_seq'::regclass) NOT NULL,
                                      username character varying(100) NOT NULL,
                                      password character varying(100) NOT NULL,
                                      name character varying(100) NOT NULL,
                                      avatar character varying(255),
                                      remember_token character varying(100),
                                      created_at timestamp without time zone DEFAULT now(),
                                      updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE goadmin_users OWNER TO postgres;

--
-- Data for Name: goadmin_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_menu (id, parent_id, type, "order", title, plugin_name, header, icon, uri, created_at, updated_at)
VALUES
    (1, 0, 1, 2, 'Admin', '', NULL, 'fa-tasks', '', now(), now()),
    (2, 1, 1, 2, 'Users', '', NULL, 'fa-users', '/info/manager', now(), now()),
    (3, 1, 1, 3, 'Roles', '', NULL, 'fa-user', '/info/roles', now(), now()),
    (4, 1, 1, 4, 'Permission', '', NULL, 'fa-ban', '/info/permission', now(), now()),
    (5, 1, 1, 5, 'Menu', '', NULL, 'fa-bars', '/menu', now(), now()),
    (6, 1, 1, 6, 'Operation log', '', NULL, 'fa-history', '/info/op', now(), now()),
    (7, 0, 1, 1, 'Dashboard', '', NULL, 'fa-bar-chart', '/', now(), now());

INSERT INTO goadmin_permissions (id, name, slug, http_method, http_path, created_at, updated_at)
VALUES
    (1, 'All permission', '*', '', '*', now(), now()),
    (2, 'Dashboard', 'dashboard', 'GET,PUT,POST,DELETE', '/', now(), now());


--
-- Data for Name: goadmin_role_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_role_menu (role_id, menu_id, created_at, updated_at)
VALUES
    (1, 1, now(), now()),
    (1, 7, now(), now()),
    (2, 7, now(), now());


--
-- Data for Name: goadmin_role_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_role_permissions (role_id, permission_id, created_at, updated_at)
VALUES
    (1, 1, now(), now()),
    (1, 2, now(), now()),
    (2, 2, now(), now());


--
-- Data for Name: goadmin_role_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_role_users (role_id, user_id, created_at, updated_at)
VALUES
    (1, 1, now(), now()),
    (2, 2, now(), now());


--
-- Data for Name: goadmin_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_roles (id, name, slug, created_at, updated_at)
VALUES
    (1, 'Administrator', 'administrator', now(), now()),
    (2, 'Operator', 'operator', now(), now());

--
-- Data for Name: goadmin_user_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_user_permissions (user_id, permission_id, created_at, updated_at)
VALUES
    (1, 1, now(), now()),
    (2, 2, now(), now());

--
-- Data for Name: goadmin_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO goadmin_users (id, username, password, name, avatar, remember_token, created_at, updated_at)
VALUES
    (1, 'admin', '$2a$10$OxWYJJGTP2gi00l2x06QuOWqw5VR47MQCJ0vNKnbMYfrutij10Hwe', 'Admin', NULL, 'tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh', now(), now()),
    (2, 'operator', '$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.', 'Operator', NULL, NULL, now(), now());


--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_menu_myid_seq', 7, true);


--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_operation_log_myid_seq', 1, true);


--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_permissions_myid_seq', 2, true);


--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_roles_myid_seq', 2, true);


--
-- Name: goadmin_site_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_site_myid_seq', 1, true);


--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_session_myid_seq', 1, true);


--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('goadmin_users_myid_seq', 2, true);


--
-- Name: goadmin_menu goadmin_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_menu
    ADD CONSTRAINT goadmin_menu_pkey PRIMARY KEY (id);


--
-- Name: goadmin_operation_log goadmin_operation_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_operation_log
    ADD CONSTRAINT goadmin_operation_log_pkey PRIMARY KEY (id);


--
-- Name: goadmin_permissions goadmin_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_permissions
    ADD CONSTRAINT goadmin_permissions_pkey PRIMARY KEY (id);


--
-- Name: goadmin_roles goadmin_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_roles
    ADD CONSTRAINT goadmin_roles_pkey PRIMARY KEY (id);


--
-- Name: goadmin_site goadmin_site_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_site
    ADD CONSTRAINT goadmin_site_pkey PRIMARY KEY (id);


--
-- Name: goadmin_session goadmin_session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_session
    ADD CONSTRAINT goadmin_session_pkey PRIMARY KEY (id);


--
-- Name: goadmin_users goadmin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY goadmin_users
    ADD CONSTRAINT goadmin_users_pkey PRIMARY KEY (id);
