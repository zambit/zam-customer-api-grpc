create table profile_data_statuses (
  id   serial primary key,
  name varchar(16) not null unique
);

insert into profile_data_statuses (name)
values ('pending'),
  ('verified'),
  ('declined');

create type customer_gender as enum ('male', 'female', 'undefined');

create table customer_profile (
  id          serial primary key,
  customer_id int references customers(id)             not null unique,
  status_id   int references profile_data_statuses(id) not null,
  email       varchar(60)                              not null,
  first_name  varchar(60)                              not null,
  last_name   varchar(60)                              not null,
  birth_date  date                                     not null,
  gender      customer_gender                          not null,
  country     varchar(60)                              not null,
  address     jsonb                                    not null
);

create index on customer_profile(customer_id);