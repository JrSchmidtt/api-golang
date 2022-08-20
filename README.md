# api-postgresql Golang Boilerplate

https://www.youtube.com/watch?v=MD7b-iQMC24&ab_channel=AprendaGolang


## 1 - Criando abiente para Desenvolvimento local

1.1 - Criando container docker com postgres:
```bash
docker run -d --name db-api-todo -p 5432:5452 -e POSTGRES_PASSWORD=1234 postgres
```

1.2 - Conectando ao container docker acessando Postgres:
```bash
docker exec -it db-api-todo psql -U postgres
```

1.3 - Criando Configuração do banco
```postgres
create database api_todo;
create user user_todo;
alter user user_todo with encrypted password '0808';
grant all privileges on database api_todo to user_todo;
grant all privileges on all tables in schema public to user_todo;
```

1.4 - Lista de database 
```postgres
\l
```
1.5 - Conectando ao banco de dados "api_todo"
```postgres
\c api_todo
```
1.6 - Lista de tabelas 
```postgres
\dt
```
1.7 - Criando tabela "todos"
```postgres
create table todos (id serial primary key,name varchar, description text, done bool);
```
1.8 - Checando tabela todos
```postgres
\dt
```
1.9 - Rodando Aplicação (Bash)
```bash
go run main.go
```
