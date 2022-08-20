# api-postgresql
GO LANG 

https://www.youtube.com/watch?v=MD7b-iQMC24&ab_channel=AprendaGolang

Lenvantando container com postgresql :
```
docker run -d --name api-todo -p 5432:5452 -e POSTGRES_PASSWORD=1234 postgresql:13.5
```

Criando Database para Desenvolvimento local :
```
docker exec -it api-todo psql -U postgres create database api_todo;
docker exec -it api-todo psql -U postgres create user user_todo;
```