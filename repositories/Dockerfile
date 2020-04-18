FROM mariadb
# Add a database
ENV MYSQL_DATABASE DHT

COPY ./CREATE_DATABASE.sql /docker-entrypoint-initdb.d/
