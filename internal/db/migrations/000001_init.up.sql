begin;

--create type link_status as ENUM('pending', 'processed', 'failed');

create table link (
    id serial primary key,
    url varchar(255) unique not null,
    status varchar(255), --link_status not null,
    created_at timestamp default now() not null,
    processed_at timestamp
);

commit;