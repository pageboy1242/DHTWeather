docker run -p 127.0.0.1:3306:3306/tcp --name dht-mysql -e MYSQL_ROOT_PASSWORD=password -d dht-mysql
docker run -d -p 9090:9090 --link dht-mysql:localhost dht-weather