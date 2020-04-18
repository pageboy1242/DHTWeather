-- Place all database creation and update scripting in this file
CREATE DATABASE IF NOT EXISTS DHT;

USE DHT;

CREATE TABLE IF NOT EXISTS readings (
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    createdate DATETIME NOT NULL,
    nodeid INTEGER NOT NULL,
    temperature DOUBLE,
    humidity DOUBLE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE readings ADD INDEX(createdate);
