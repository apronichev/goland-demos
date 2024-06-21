Assuming you're in `docker-env-compose` folder.
To run service:
```
docker-compose -f ./${COMPOSE_FILE_NAME?} pull ${SERVICE_NAME?}
```

To run MySQL 8.2:
```bash
docker-compose -f ./docker-compose.mysql.yml up -d mysql_8.2.X
```
