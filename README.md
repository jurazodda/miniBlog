# Mini-Blog Platform

Платформа для мини-блога на Go с функционалом создания, редактирования, удаления постов, комментариев, лайков и репостов.

------------------------------------------------------

## Функционал

- Аутентификация: регистрация и вход.
- Посты: создание, редактирование, удаление.
- Комментарии: добавление и удаление комментариев.
- Лайки: лайк постов.
- Репосты: репост с soft delete.

## Технологии

- **Go**: основной язык.
- **Gin**: веб-фреймворк.
- **Gorm**: ORM для работы с PostgreSQL.
- **PostgreSQL**: база данных.
- **Zerolog**: логгирование.

## Установка

### 1. Клонировать репозиторий:

   ```bash
   git clone https://github.com/jurazodda/mini-blog
   cd mini-blog
   ```

### 2. Установить зависимости:

```bash
go mod download
```

### 3. Настроить базу данных и выполнить миграции:

```bash
go run cmd/migrate.go
```

### 4. Запустить приложение:

```bash
go run cmd/main.go
```