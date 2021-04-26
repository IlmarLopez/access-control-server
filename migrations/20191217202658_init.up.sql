create table building_access
(
    id          varchar(36)  null,
    building_id varchar(36)  null,
    user_id     varchar(36)  null,
    check_in    datetime     null,
    check_out   datetime     null,
    created_at  datetime     null,
    description varchar(150) null,
    update_at   datetime     null,
    is_approved int          null,
    approved_by varchar(36)  null
);

create table buildings
(
    id          varchar(36)  not null
        primary key,
    name        varchar(50)  null,
    description varchar(150) null,
    user_limit  int          null,
    created_at  datetime     null,
    updated_at  datetime     null,
    is_active   int          null
);

create table careers
(
    id        varchar(36) null,
    name      varchar(60) not null,
    is_active int         null
);

create table groups
(
    id        varchar(36) null,
    career_id varchar(36) null,
    name      varchar(5)  null,
    is_active int         null,
    constraint groups_id_uindex
        unique (id)
);

create table roles
(
    id          varchar(36)  not null
        primary key,
    name        varchar(50)  null,
    description varchar(100) null
);

create table users
(
    id                  varchar(36)  not null
        primary key,
    username            varchar(50)  not null,
    password            varchar(150) null,
    created_at          datetime     null,
    updated_at          datetime     null,
    is_active           int          null,
    role_id             varchar(36)  null,
    first_name          varchar(50)  null,
    last_name           varchar(50)  null,
    email               varchar(50)  null,
    registration_number varchar(10)  null,
    career_id           varchar(36)  null,
    group_id            varchar(36)  null,
    constraint users_username_uindex
        unique (username)
);

