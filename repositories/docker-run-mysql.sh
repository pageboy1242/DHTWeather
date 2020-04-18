docker run -p 127.0.0.1:3306:3306/tcp --name dht-mysql -e MYSQL_ROOT_PASSWORD=password -v $HOME/mysql-data:/var/lib/mysql -d dht-mysql
