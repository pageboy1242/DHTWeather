FROM mysql
# Add a database
ENV MYSQL_DATABASE DHT

COPY ./CREATE_DATABASE.sql /docker-entrypoint-initdb.d/
COPY ./CREATE_USER.sql /docker-entrypoint-initdb.d/