create database poketracker;
\connect poketracker;

create table Pokemon
(
    dex             int           not null,
    name            varchar(255)  not null,
    types           varchar(50)[] not null,
    shiny           boolean       not null,
    normal          boolean       not null,
    universal       boolean       not null,
    regional        boolean       not null,
    userId          varchar(255)  not null references UserInfo (userId),
    normalSpriteUrl varchar(255)  not null,
    shinySpriteUrl  varchar(255)  not null,
    PRIMARY KEY (dex, userId)
);

create table Edition
(
    name varchar(255) unique primary key,
    gen  int not null
);

create table PokemonEditionRelation
(
    dex         int          not null,
    editionName varchar(255) not null,
    userId      varchar(255) not null,
    FOREIGN KEY (dex, userid) REFERENCES Pokemon (dex, userid),
    FOREIGN KEY (editionName) REFERENCES Edition (name),
    UNIQUE (dex, userId, editionName)
);

create table UserInfo
(
    userId    varchar(255) not null primary key,
    avatarUrl varchar(255) not null,
    bulkMode  boolean      not null
);

insert into edition
values ('Rot', 1);
insert into edition
values ('Blau', 1);
insert into edition
values ('Gelb', 1);
insert into edition
values ('Gold', 2);
insert into edition
values ('Silber', 2);
insert into edition
values ('Kristall', 2);
insert into edition
values ('Sphir', 3);
insert into edition
values ('Rubin', 3);
insert into edition
values ('Smaragd', 3);
insert into edition
values ('Feuerrot', 3);
insert into edition
values ('Blattgrün', 3);
insert into edition
values ('Diamant', 4);
insert into edition
values ('Perl', 4);
insert into edition
values ('Platin', 4);
insert into edition
values ('Heart Gold', 4);
insert into edition
values ('Soul Silver', 4);
insert into edition
values ('Schwarz', 5);
insert into edition
values ('Weiß', 5);
insert into edition
values ('Schwarz 2', 5);
insert into edition
values ('Weiß 2', 5);
insert into edition
values ('X', 6);
insert into edition
values ('Y', 6);
insert into edition
values ('Omega Rubin', 6);
insert into edition
values ('Alpha Saphir', 6);
insert into edition
values ('Sonne', 7);
insert into edition
values ('Mond', 7);
insert into edition
values ('Ultra Sonne', 7);
insert into edition
values ('Ultra Mond', 7);
insert into edition
values ('Let''s Go, Pikachu!', 7);
insert into edition
values ('Let''s Go, Evoli!', 7);
insert into edition
values ('Schwert', 8);
insert into edition
values ('Schild', 8);
insert into edition
values ('Strahlender Diamant', 8);
insert into edition
values ('Leuchtende Perle', 8);
insert into edition
values ('Legenden: Arceus', 8);
insert into edition
values ('Karmesin', 9);
insert into edition
values ('Purpur', 9);