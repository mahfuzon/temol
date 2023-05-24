CREATE TABLE IF NOT EXISTS users (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    email varchar(255) UNIQUE DEFAULT NULL,
    password varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL UNIQUE,
    address varchar(255) DEFAULT NULL,
    profile_image varchar(255) DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB;