CREATE TABLE movies
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	Name varchar(150) PRIMARY KEY,
	Description varchar(1000),
	realize_date timestamp,
	rating integer
);

CREATE TABLE actors
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	Name varchar(150) PRIMARY KEY,
	sex boolean,
	born timestamp
);

CREATE TABLE lists
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	Movie varchar(150),
	Actor varchar(150),
	FOREIGN KEY (Actor) REFERENCES actors(Name),
	FOREIGN KEY (Movie) REFERENCES movies(Name)
);

CREATE TABLE users
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	id integer,
	login TEXT PRIMARY KEY,
	role TEXT
);

INSERT INTO users (login, role) VALUES ('login', 'admin');
INSERT INTO users (login, role) VALUES ('newlogin', 'user');