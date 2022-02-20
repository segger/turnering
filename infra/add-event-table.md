Migrate

```
create table event (id uuid not null, name varchar(50) not null, class_nbr smallint not null, sort_order int not null, contest_id uuid not null, primary key (id), foreign key (contest_id) references contest (id));
```

```
alter table result add column event_id uuid;
alter table result alter column time set not null;
```

```
insert into event select uuid_generate_v4(), event_name, class_nbr, row_number() over() as sort_order, contest_id from result join participant on result.participant_id=participant.id where contest_id = '<>' group by event_name, class_nbr, contest_id;
```

```
update result set event_id = event.id from event where result.contest_id=event.contest_id and result.event_name=event.name;

alter table result alter column event_id set not null;
alter table result drop column event_name;
```
