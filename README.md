# docker-mysql-starter

DRY-way to setup MySQL locally with Docker.

```bash
# Start docker.
$ make up

# Stop docker.
$ make down

# For the brave who uses mysql cli...
$ make mysql

# Else...
# Use Sequel Pro on macOS

# Remove the tmp/ directory created by Docker.
$ make clean
```

## References

- https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30
