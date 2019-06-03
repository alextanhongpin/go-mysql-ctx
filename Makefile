include .env
export

mysql: # Access the MySQL cli.
	@mysql -h ${DB_HOST} -u ${DB_USER} -p ${DB_NAME}

dump: # Backup the database.
	@mysqldump -h ${DB_HOST} -u ${DB_USER} -p ${DB_NAME} > backup.sql

recover: # Recover the backup.
	@mysql -u root -p -h ${DB_HOST} ${DB_NAME} < backup.sql

clean: # Clear the temporary MySQL folder.
	@rm -rf tmp

up: # Start the docker.
	@docker-compose up -d

down: # Stop the docker.
	@docker-compose down 
