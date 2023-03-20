# shareme-go

ShareMe, written in Go

## Configuration

> environment variable will also be loaded from `.env` file

```ini
# Default port is 8080
PORT=3000

# Use Deta Base (no need to configure in deta)
DETA_PROJECT_KEY=projectid_secretkey
DETA_BASE_NAME=(leave blank to use ShareMe)

# Use MySQL
MYSQL_USERNAME=root
MYSQL_PASSWORD=password
MYSQL_HOST=(leave blank to use 127.0.0.1)
MYSQL_PORT=(leave blank to use 3306)
MYSQL_DB_NAME=database_name
MYSQL_TABLE_NAME=table_name

# Use MongoDB
MONGODB_URI=mongodb+srv://username:password@host
MONGODB_NAME=database_name
MONGODB_COLLECTION=collection_name
```

If no database is specified, the program will use memory to store data,  
which means all the data will disappear after the program's shutdown
