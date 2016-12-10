drop database gocrud;
create database gocrud;
drop user gocrud;
create user gocrud with password 'gocrud';
grant all privileges on database gocrud to gocrud;