alter table customers
  add column created_at time without time zone,
  add column updated_at time without time zone;

update customers
set
  created_at = now(),
  updated_at = now();

insert into customer_statuses (name)
values ('verified');
insert into customer_statuses (name)
values ('created');
