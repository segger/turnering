### docker.io postgres

```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

##### participant
| field | type |
| ----- | ---- |
| id    | uuid |
| first_name | varchar(50) |
| dog_name   | varchar(50) |
| class_nbr  | smallint    |

##### contest
| field | type |
| ----- | ---- |
| id    | uuid |
| name  | varchar(100) |

contest-code?

##### result
| field | type |
| ----- | ---- |
| id    | uuid |
| created_at | timestamp |
| deleted | boolean |
| participant_id | uuid |
| contest_id | uuid |
| event_name | varchar(50) |
| points | real |
| errors | integer |
| time | integer |
| sse | boolean |

##### protocol
##### protocol_event
##### result
##### event
##### tournament

### google cloud console

connect ssh - settings - file transfer

### caddy

```
sudo setcap CAP_NET_BIND_SERVICE=+eip $(which caddy)
```

```
caddy stop | start
```