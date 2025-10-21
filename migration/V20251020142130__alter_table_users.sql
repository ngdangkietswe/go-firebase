alter table users
    add column firebase_uid varchar(255) unique,
    add column display_name varchar(255);