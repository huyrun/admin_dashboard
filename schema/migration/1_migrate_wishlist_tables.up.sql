ALTER TABLE ONLY entity_tags ALTER COLUMN tag_id TYPE BIGINT USING tag_id::BIGINT;
CREATE SEQUENCE entity_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY entity_tags ADD COLUMN id BIGINT DEFAULT nextval('entity_tags_id_seq'::regclass);

CREATE SEQUENCE liked_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY liked_entities ADD COLUMN id BIGINT DEFAULT nextval('liked_entities_id_seq'::regclass);

CREATE SEQUENCE saved_entities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY saved_entities ADD COLUMN id BIGINT DEFAULT nextval('saved_entities_id_seq'::regclass);

CREATE SEQUENCE tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY tags ADD COLUMN id BIGINT DEFAULT nextval('tags_id_seq'::regclass);

CREATE SEQUENCE user_relationships_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY user_relationships ADD COLUMN id BIGINT DEFAULT nextval('user_relationships_id_seq'::regclass);

CREATE SEQUENCE wishes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY wishes ADD COLUMN id BIGINT DEFAULT nextval('wishes_id_seq'::regclass);

ALTER TABLE wish_stories ALTER COLUMN image TYPE VARCHAR(1024) USING image::VARCHAR(1024);