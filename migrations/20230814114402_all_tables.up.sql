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
    id varchar(14) PRIMARY KEY,
    name varchar(45) NOT NULL,
    email varchar(45) NOT NULL,
    phone_number varchar(14) NOT NULL,
    password varchar(75) NOT NULL,
    is_verified boolean DEFAULT FALSE,
    role varchar(15) DEFAULT 'user',
    UNIQUE(email)
);

CREATE TABLE notifications(
    id varchar(10) PRIMARY KEY,
    user_id varchar(14) NOT NULL,
    message TEXT NOT NULL,
    mark_as_read boolean DEFAULT FALSE,
    CONSTRAINT fk_notif_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);
