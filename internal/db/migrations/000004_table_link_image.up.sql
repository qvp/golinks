
create table link_image (
    id serial primary key,
    link_id int not null,
    url varchar(2048) not null,

    constraint unique_link_id_url unique (link_id, url)
);
