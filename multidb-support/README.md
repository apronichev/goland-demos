1. In the 'docker-compose' folder, open YML files and pull and run images of several databases (preferably, MYSQL, PostgreSQL).
2. Create a local sqlite instance in the Database tool window. For more information about installation, refer to [Connect to an SQLite database](https://www.jetbrains.com/help/go/2024.2/sqlite.html#connect-to-sqlite-database).
3. In *Run/debug configuration* dialog, create 'Database script' run configuration for MySQL, PostgreSQL, and sqlite. Include in scripts the following files from the 'dumps' folder:
   1. sakila-schema.sql
   2. sakila-data.sql
   3. *-sakila-insert-data.sql
4. Run created run configurations, ignore warnings and possible errors.
5. Add all databases that are run in the Docker to the 'Database' tool window. You can do it by using the JDBC link in YML files (for example, jdbc:mysql://localhost:33082/guest?user=guest&password=guest). For more information about adding databases, refer to the docs:
   1. [MySQL](https://www.jetbrains.com/help/go/2024.2/mysql.html#connect-to-mysql-database)
   2. [PostgreSQL](https://www.jetbrains.com/help/go/2024.2/postgresql.html#connect-to-postgresql-database)
   3. [SQLite](https://www.jetbrains.com/help/go/2024.2/sqlite.html#connect-to-sqlite-database)