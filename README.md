# adoetl
Extract data from ADO and save it to a database for better data analysis

# Development
execute `run.sh` to run the appliation

# install postgres
```
docker container rm --force postgres-local
docker run --name postgres-local -e POSTGRES_PASSWORD=P@ssw0rd -p 5432:5432 -d postgres:latest
docker start postgres-local
```

