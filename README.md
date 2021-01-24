# Vincivia Mono Repo
Is made up of a golang server and a react client
 
# Client
Todo

# Server


## Steps to setup the a database locally (on Mac)
```
check config.go for ENV_VARS required
brew install postgresql@11
brew services start postgresql@11
cat /usr/local/var/log/postgresql@11.log
rm /usr/local/var/postgresql@11/postmaster.pid
brew services restart postgresql@11
pgadmin4
password: password
create database: vincivia

