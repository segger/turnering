#### 2022-02-20

##### Create event table
```
create table event (id uuid not null, name varchar(50) not null, class_nbr smallint not null, sort_order int not null, contest_id uuid not null, primary key (id), foreign key (contest_id) references contest (id));
```
##### Prepare result table
```
alter table result add column event_id uuid;
alter table result alter column time set not null;
```

###### Migrate data
```
insert into event select uuid_generate_v4(), event_name, class_nbr, row_number() over() as sort_order, contest_id from result join participant on result.participant_id=participant.id where contest_id = '<>' group by event_name, class_nbr, contest_id;
```
```
update result set event_id = event.id from event where result.contest_id=event.contest_id and result.event_name=event.name;
```
##### Update result table
```
alter table result alter column event_id set not null;
alter table result drop column event_name;
```

##### Add events
```
insert into event (id, name, class_nbr, sort_order, contest_id) values (uuid_generate_v4(), 'Sök 1', 2, 1, '98bee440-42b8-41b7-9d58-4889586d9a99');
insert into event (id, name, class_nbr, sort_order, contest_id) values (uuid_generate_v4(), 'Sök 2', 2, 2, '98bee440-42b8-41b7-9d58-4889586d9a99');
insert into event (id, name, class_nbr, sort_order, contest_id) values (uuid_generate_v4(), 'Sök 3', 2, 3, '98bee440-42b8-41b7-9d58-4889586d9a99');
```

#### Set results as deleted
```
alter table participant add column deleted boolean not null default false;
```