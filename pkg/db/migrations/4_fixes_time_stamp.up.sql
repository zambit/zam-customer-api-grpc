alter table customers
  alter column registered_at type timestamp without time zone using current_date + registered_at;
alter table customers
  alter column created_at type timestamp without time zone using current_date + created_at;
alter table customers
  alter column updated_at type timestamp without time zone using current_date + updated_at;