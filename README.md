## GoLinks
Приложение получает от пользователя ссылку на веб страницу и скачивает все изображения с нее.

PUT /link - добавляет ссылку в очередь обработки и возвращает task id  
GET /task/:id - возвращает статус и адреса изображений  

## GUI
[http://localhost:8080](http://localhost:8080) - Application swagger  
[http://localhost:15672](http://localhost:15672) - RabbitMQ management (root / root)  

## Migration
Подготовить адрес соединения с бд:  
`export POSTGRESQL_URL='postgres://golinks:golinks@localhost:5432/golinks?sslmode=disable'`

Создать новую миграцию:  
`migrate create -ext sql -dir internal/db/migrations -seq <name>`

Запустить миграции:  
`migrate -database ${POSTGRESQL_URL} -path internal/db/migrations up`

Откатить последнюю миграцию:  
`migrate -database ${POSTGRESQL_URL} -path internal/db/migrations down 1`

## TODO
Сделать миграцию ALTER TABLE link REPLICA IDENTITY FULL;

![](docs/architecture.png)