CREATE TABLE airports(
    id varchar(8) PRIMARY KEY,
    airport_code varchar(3) NOT NULL,
    airport_name varchar(40) NOT NULL,
    location text NOT NULL, 
    location_acronym varchar(3) NOT NULL,
    UNIQUE(airport_name),
    UNIQUE(location)
);

CREATE TABLE airlines(
    id varchar(8) PRIMARY KEY, 
    airline_code varchar(2) NOT NULL,
    airline_name varchar(40) NOT NULL,
    airline_image varchar(65) NOT NULL,
    UNIQUE(airline_name)
);

CREATE TABLE users(
    id uuid PRIMARY KEY,
    name varchar(45) NOT NULL,
    email varchar(45) NOT NULL,
    phone_number varchar(14) NOT NULL,
    password varchar(40) NOT NULL,
    UNIQUE(email)
);
