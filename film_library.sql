CREATE TABLE movies
(
	Name varchar(150) PRIMARY KEY,
	Description varchar(1000),
	realize_date timestamp,
	rating integer
);

CREATE TABLE actors
(
	Name varchar(150) PRIMARY KEY,
	sex boolean,
	born timestamp
);

CREATE TABLE lists
(
	Movie varchar(150),
	Actor varchar(150),
	FOREIGN KEY Movie REFERENCES movies(name),
	FOREIGN KEY Actor REFERENCES movies(name)
)