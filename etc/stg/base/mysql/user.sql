CREATE USER 'native'@'%' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'native'@'%';
ALTER USER 'native'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
