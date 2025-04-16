## GoLinks
Парсит веб страницу по переданной ссылке и возвращает url картинок из тегов img.  
REST API: ссылка сохраняется в БД, откуда попадает в очередь и обрабатывается асинхронно.  
gRPC: ссылка обрабатывается сразу и результат возвращается синхронно.  

Стек:
* Golang
* PostgreSQL
* RabbitMQ
* Debezium
* Fiber
* sqlc
* migrate
* pgx
* zerolog
* swagger

![](docs/architecture.png)

## Run
`make up` - запустить приложение в докер.  
[http://localhost:8080/docs](http://localhost:8080/docs) - Application swagger  
[http://localhost:15672](http://localhost:15672) - RabbitMQ management (root / root)
