# shareme-go

ShareMe, written in golang

# Usage

```sh
export PORT=8888; ./shareme.exe
```

# Database Config

> Config database via environment variable or `.env` file

Using MySQL

```ini
MYSQL_USERNAME=root
MYSQL_PASSWORD=password
MYSQL_HOST=(leave blank to use 127.0.0.1)
MYSQL_PORT=(leave blank to use 3306)
MYSQL_DB_NAME=database_name
MYSQL_TABLE_NAME=table_name
```

Using MongoDB

```ini
MONGO_DB_URI=mongodb+srv://username:passwrod@host
MONGO_DB_NAME=database_name
MONGO_DB_COLLECTION=collection_name
```

If no database is specified, the program use memory to store data,  
which means the data will disappear after the program is shutdown
