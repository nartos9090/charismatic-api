create table charismatic_dev.user
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

create table charismatic_dev.copywriting_project
(
    id            int auto_increment
        primary key,
    user_id       int                                  not null,
    title         varchar(255)                         not null,
    product_image varchar(255)                         not null,
    brand_name    varchar(255)                         not null,
    market_target varchar(255)                         not null,
    superiority   text                                 not null,
    result        text                                 null,
    created_at    datetime default current_timestamp() not null,
    constraint copywriting_project_user_id_fk
        foreign key (user_id) references charismatic_dev.user (id)
);

create table charismatic_dev.video_project
(
    id            int auto_increment
        primary key,
    user_id       int                                  not null,
    product_title varchar(255)                         not null,
    brand_name    varchar(255)                         not null,
    product_type  varchar(255)                         not null,
    market_target varchar(255)                         not null,
    superiority   text                                 not null,
    duration      int                                  not null,
    created_at    datetime default current_timestamp() not null,
    constraint video_project_user_id_fk
        foreign key (user_id) references charismatic_dev.user (id)
);

create table charismatic_dev.scene
(
    id               int auto_increment
        primary key,
    video_project_id int          not null,
    sequence         int          not null,
    title            text         not null,
    narration        text         not null,
    illustration     text         not null,
    illustration_url varchar(255) null,
    voice_url        varchar(255) null,
    constraint scene_video_project_id_fk
        foreign key (video_project_id) references charismatic_dev.video_project (id)
);

