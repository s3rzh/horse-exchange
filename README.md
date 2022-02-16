ТЗ:

Биржа лошадей

Цель
Цель задания - выяснить уровень владения ООП, навыки построения абстракций и знание принципов.
Разработать модель на языке программирования Golang для данной предметной области. Описать типы сущностей, сервисов и интерфейсы для абстрагирования от слоя хранения данных.
Реализовывать графический интерфейс с контроллерами и представлениями и хранение данных в БД не обязательно. Можно ограничиться интерфейсами репозиториев или заглушками.

Все что не описано на совести и фантазии разработчика.

Описание

Биржа лошадей позволяет покупать или брать в аренду лошадей для различных задач:
земледелие
работа на каменоломне
перевозки грузов
и т.п.

Для удобства клиентов весь ассортимент разбит по категориям, например:
породы
назначение

Для описания лошади используются характеристики
кличка
пол
возраст
вес
порода
место рождения
повадки
ставка почасовой аренды
стоимость покупки

Функционал
Довольно примитивный и не сложный

Покупка лошади
пользователь выбирает лошадь из каталога
пользователь покупает лошадь, биржа проверяет свободна ли эта лошадь, если она занята (то есть уже арендована), то такую лошадь нельзя купить

Аренда лошади
пользователь находит подходящую лошадь в каталоге
пользователь может создать запрос на аренду лошади
если в эти даты лошадь доступна, то запрос удовлетворяется и пользователю дается договор аренды со стоимостью
пользователь подписывает и оплачивает договор
биржа проверяет если договор не оплачен, то запрос пользователя на аренду не учитывается и лошадь считается свободной для аренды

ПО: Git, Docker, Postman

Клонировать проект:
git clone git@gitlab.com:dev-area/horse-exchange.git

Перейти в папку проекта и cоздать в корне файл .env

Добавить в него
DB_PASSWORD=qwerty

В корне проекта, выполнить след. команды:
docker-compose up -d --build horse-exchange-app

Создание образа для миграций
docker build -t migrator ./api/migrator

Запустить миграции
docker run --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" up

В Postman

(получение списка лошадей)
GET http://localhost:8000/api/horses


(список лошадей по каталогам)
GET http://localhost:8000/api/catalog/breeds
GET http://localhost:8000/api/catalog/breeds/2

GET http://localhost:8000/api/catalog/tasks
GET http://localhost:8000/api/catalog/tasks/3


купить лошадь
PATCH http://localhost:8000/api/horses/2/buy


запросить аренду
POST http://localhost:8000/api/horses/rent

В Body/raw/JSON добавить

{
	"horse_id": 1, 
	"since": "2021-06-23 00:00:00", 
	"till": "2021-06-23 23:59:59"
}

оплатить аренду (ссылку получить из предыдущего запроса)
PATCH http://localhost:8000/api/horses/rent/1/pay

docker exec user_clickhouse bash -c "clickhouse-client -mn < /var/lib/clickhouse/migrations/1_init_schema.up.sql"


docker exec user_clickhouse bash -c "echo $PATH"

Для отката миграций
docker run --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" down -all




docker rm user_postgres user_redis user_rabbitmq user_clickhouse user_logger user_app


дата: лето 2021