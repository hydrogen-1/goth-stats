create table if not exists visits(
    id bigserial primary key,
    t timestamptz not null
);