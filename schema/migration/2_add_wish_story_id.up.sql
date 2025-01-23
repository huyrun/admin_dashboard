CREATE SEQUENCE IF NOT EXISTS wish_stories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY wish_stories ADD COLUMN id BIGINT DEFAULT nextval('wish_stories_id_seq'::regclass);