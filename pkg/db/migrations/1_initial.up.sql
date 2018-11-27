create table customer_statuses (
  id   serial primary key,
  name varchar(63)
);
create table customers (
  id            serial primary key,
  phone         varchar(255),
  password      text,
  registered_at time without time zone,
  referrer_id   integer,
  status_id     integer,
  constraint customer_referrer_id_fk foreign key (referrer_id) references customers(id)
    on delete set null,
  constraint customer_statuses_id_fk foreign key (status_id) references customer_statuses(id)
);

insert into customer_statuses (name)
values ('pending');
insert into customer_statuses (name)
values ('active');