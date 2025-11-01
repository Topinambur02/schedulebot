# Schedule bot

## Описание проекта

### Используемые технологии:

<img src="https://github.com/user-attachments/assets/50937523-df97-4fff-929e-0961f0555f4d" title="Golang" alt="golang" width="75" height="75"/>
<img src="https://github.com/user-attachments/assets/387c9b9c-ce0f-49ef-9112-10de154d1ced" title="GORM" alt="gorm" width="75" height="75"/>
<img src="https://github.com/user-attachments/assets/82b7f5a2-0559-48f7-beb9-9ecef32bd55c" title="SQLite3" alt="sqlite" width="75" height="75"/>
<img src="https://github.com/user-attachments/assets/096102ab-304b-4b89-b2d6-08f76b1aaa76" title="Redis" alt="redis" width="75" height="75"/>
<img src="https://github.com/user-attachments/assets/648aa8d0-0ba1-4362-a281-2e252e2ec97b" title="Docker" alt="docker" width="75" height="75"/>

### Основные функции:

- Загрузка расписания
- Автоматическая генерация миграций (GORM)
- Кеширование запросов к базе данных

### Предварительные требования:

- Golang v1.24.2 или выше
- SQLite v3.13.1 или выше
- Redis v8.2.2 или выше
- Docker v27.5.1 (для docker-запуска)

## Установка и запуск (без Docker)

1. Клонируйте репозиторий:

```
git clone https://github.com/Topinambur02/schedulebot
cd schedulebot
```

2. Создайте файл .env:

```
# Api токен для telegram бота
TELEGRAM_API_TOKEN=your_telegram_api_token

# Ссылка на получение расписания
URL_GET_SCHEDULE_API=https://example.com/api

# Логин и пароль от профиля Московского Политеха
ULOGIN=login
UPASSWORD=password

#  База данных Redis (замените значения на свои!)
REDIS_ADDR=your_host:your_port
```

3. Запустите с помощью Makefile скрипт:

```
make run
```

## Запуск с помощью Docker

1. Клонируйте репозиторий:

```
git clone https://github.com/Topinambur02/schedulebot
cd schedulebot
```

2. Создайте файл .env:

```
# Api токен для telegram бота
TELEGRAM_API_TOKEN=your_telegram_api_token

# Ссылка на получение расписания
URL_GET_SCHEDULE_API=https://example.com/api

# Логин и пароль от профиля Московского Политеха
ULOGIN=login
UPASSWORD=password

#  База данных Redis (замените значения на свои!)
REDIS_ADDR=your_host:your_port
```

3. Запустите проект с помощью команды:

```
make docker-build
```

## Структура проекта

```
schedulebot/
├── Dockerfile               # Конфигурация для сборки Docker-образа приложения
├── Makefile                 # Makefile для запуска скриптов
├── README.md               # Основная документация проекта
├── cacheUtils              # Дополнительные утилиты для работы с кэшом
├── configs                  # Дополнительные конфигурационные файлы
├── database.db             # База данных
├── db                      # Утилиты для работы с базами данных
├── docker-compose.yml      # Конфигурация для развёртывания многоконтейнерного приложения
├── go.mod                  # Файл с зависимостями
├── go.sum                  # Файл для управления зависимостями
├── handlers                # Обработчики событий
├── main.go                 # Главный файл для запуска
├── model                   # Модели для БД и response
├── senders                 # Функции, которые отправляют данные
└── utils                   # Дополнительные утилиты
```