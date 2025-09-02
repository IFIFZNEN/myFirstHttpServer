# myFirstHttpServer
Мой первый сервер, написанный на http модуле



Запуск докер компос
1) Запустить и закрыть соединение
    docker compose up -d
    docker compose down


Убить процесс:

1) sudo lsof -nP -iTCP:8081 | grep LISTEN

2) kill -9 25274 

