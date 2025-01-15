CREATE SEQUENCE IF NOT EXISTS entity_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY entity_tags ADD COLUMN id BIGINT DEFAULT nextval('entity_tags_id_seq'::regclass);

CREATE SEQUENCE IF NOT EXISTS liked_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY liked_entities ADD COLUMN id BIGINT DEFAULT nextval('liked_entities_id_seq'::regclass);

CREATE SEQUENCE IF NOT EXISTS saved_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY saved_entities ADD COLUMN id BIGINT DEFAULT nextval('saved_entities_id_seq'::regclass);

CREATE SEQUENCE IF NOT EXISTS tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY tags ADD COLUMN id BIGINT DEFAULT nextval('tags_id_seq'::regclass);

CREATE SEQUENCE IF NOT EXISTS user_relationships_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY user_relationships ADD COLUMN id BIGINT DEFAULT nextval('user_relationships_id_seq'::regclass);

CREATE SEQUENCE IF NOT EXISTS wishes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY wishes ADD COLUMN id BIGINT DEFAULT nextval('wishes_id_seq'::regclass);