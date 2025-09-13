# Subscription Aggregator

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Chi](https://img.shields.io/badge/chi-%23000000.svg?style=for-the-badge&logo=&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

---

## Описание проекта

Subscription Tracker на Golang с использованием роутера Chi, реализующий RESTful API.

---

## Особенности и функциональные требования

- **Роутинг и API:**
```
POST    http://localhost:8080/subscriptions        # Создать подписку
GET     http://localhost:8080/subscriptions/{id}   # Получить подписку по ID
GET     http://localhost:8080/subscriptions         # Получить список всех подписок
GET     http://localhost:8080/subscriptions/sum     # Получить сумму цен всех подписок
PUT     http://localhost:8080/subscriptions/{id}   # Обновить подписку по ID
DELETE  http://localhost:8080/subscriptions/{id}   # Удалить подписку по ID

```

## Установка и запуск

1. Клонируйте репозиторий

```
git clone https://github.com/RoiDuNord/subscription_aggregator_api.git
```

2. Запустите Docker

```
docker-compose up -d --build
```

3. Передайте подписку с помощью json

```
POST  http://localhost:8080/subscriptions
```
```
{
    "service_name": "Sber Prime",
    "price": 200,
    "user_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "start_date": "2024-07-15"
}
```

4. Получите подписку по ID:

```
GET  http://localhost:8080/subscriptions/{id}
```

5. Получите список всех подписок:

```
GET  http://localhost:8080/subscriptions
```

6. Получите сумму цен всех подписок:

```
DELETE  http://localhost:8080/subscriptions/sum
```

7. Обновите подписку по ID:

```
PUT http://localhost:8080/subscriptions/{id}
```
```
{
    "service_name": "Sber Plus",
    "price": 250,
    "user_id": "a1b2c3d4-e5f6-7890-abcd-ef1296567890",
    "start_date": "2024-08-01"
}
```

8. Удалите подписку по ID:

```
DELETE  http://localhost:8080/subscriptions/{id}
```