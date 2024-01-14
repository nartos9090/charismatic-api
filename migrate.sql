create table user
(
    id          int auto_increment
        primary key,
    email       varchar(255) not null,
    fullname    varchar(255) not null,
    passwd      varchar(255) null,
    passwdSalt  varchar(30)  null,
    provider    varchar(30)  null,
    provider_id varchar(30)  null,
    constraint user_pk2
        unique (email)
);

