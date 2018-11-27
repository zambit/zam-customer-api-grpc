alter table customers
  drop column created_at;
alter table customers
  drop column updated_at;

update customers
set status_id = (
  select id
  from customer_statuses
  where name = 'pending'
)
where status_id = (
  select id
  from customer_statuses
  where name = 'created'
);

update customers
set status_id = (
  select id
  from customer_statuses
  where name = 'pending'
)
where status_id = (
  select id
  from customer_statuses
  where name = 'verified'
);

delete
from customer_statuses
where name = 'verified';
delete
from customer_statuses
where name = 'created';
