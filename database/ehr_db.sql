CREATE USER 'ehr_admin'@'localhost' IDENTIFIED WITH mysql_native_password BY '';
GRANT ALL PRIVILEGES ON *.* TO 'ehr_admin'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
