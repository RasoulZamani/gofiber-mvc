# Sample API with GoFiber, Gorm and PSQL in MVC Architecture


# Usage:


# DB

you can run psql by:
```
docker run --name my_postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=mydatabase \
  -p 5444:5432 \
  -v pgdata:/var/lib/postgresql/data \
  -d postgres:15
```
and set url for db in .env file.