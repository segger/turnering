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
| enabled | boolean |

##### event
| field | type |
| ----- | ---- |
| id    | uuid |
| name  | varchar(50) |
| class_nbr  | smallint |
| contest_id | uuid |

contest-code?

##### result
| field | type |
| ----- | ---- |
| id    | uuid |
| created_at | timestamp |
| deleted | boolean |
| participant_id | uuid |
| contest_id | uuid |
| event_id | uuid |
| points | real |
| errors | integer |
| time | integer |
| sse | boolean |

##### protocol
##### protocol_event
##### result
##### tournament

```
pg_dump -h localhost -U dbuser -Fc -f backup.pgdump
pg_restore -h localhost -U dbuser -d dbname
```

### systemd
`WorkingDirectory` to where `.env` file live
`ExecStart` to built `server`

### google cloud console

connect ssh - settings - file transfer

### caddy

```
sudo setcap CAP_NET_BIND_SERVICE=+eip $(which caddy)
```

```
caddy stop | start
```

```
root * /www
file_server
reverse_proxy /api/* localhost:8084
```
