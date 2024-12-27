-- name: GetTotal :one
select count(*) from visits;

-- name: GetLast24h :one
select COUNT(*)
from visits
where t >= now() - interval '24 HOURS';

-- name: InsertVisit :exec
insert into visits (t) values (now());