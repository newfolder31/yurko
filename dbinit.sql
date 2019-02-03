CREATE TABLE IF NOT EXISTS "user" (
  id SERIAL PRIMARY KEY
)

CREATE TABLE IF NOT EXISTS scheduler (
  id SERIAL PRIMARY KEY,
  user_id SERIAL NOT NULL REFERENCES user(id),
  profession_type varchar(50)
)


CREATE TABLE IF NOT EXISTS "interval" (
  id SERIAL PRIMARY KEY,
	scheduler_id SERIAL NOT NULL REFERENCES "user"(id),
	"from" INTEGER,
	"to" INTEGER,
	"date" TIMESTAMP,
	weekDay SMALLINT
)