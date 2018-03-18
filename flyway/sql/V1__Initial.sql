CREATE TABLE sources
(
  id    INT AUTO_INCREMENT,
  title VARCHAR(300),
  url   VARCHAR(300) CHARSET latin1,
  PRIMARY KEY (id),
  UNIQUE u1 (url)
);

CREATE TABLE articles
(
  id          INT AUTO_INCREMENT,
  source_id   INT,
  title       VARCHAR(300),
  link        VARCHAR(300) CHARSET latin1,
  description TEXT,
  PRIMARY KEY (id),
  UNIQUE u1 (link),
  CONSTRAINT articles_sources_id_fk FOREIGN KEY (source_id) REFERENCES sources (id)
);
