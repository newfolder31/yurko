
CREATE TABLE usr (
  id SERIAL,
  email character varying(255) NOT NULL,
  password character varying(255) NOT NULL,
  is_admin boolean NOT NULL,
  is_active boolean NOT NULL,
  first_name character varying(255),
  last_name character varying(255),
  fathers_name character varying(255),

  CONSTRAINT usr_pkey PRIMARY KEY (id),
  CONSTRAINT usr_email_key UNIQUE (email)
);

ALTER TABLE usr OWNER TO yurkorole;