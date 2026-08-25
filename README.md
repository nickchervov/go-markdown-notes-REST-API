# Markdown Notes REST API

REST API для управления заметками в формате Markdown с полнотекстовым поиском и rate limiting.

---

## Технологии

- **Go** + chi (роутинг)
- **PostgreSQL** + sqlx + pgx (БД)
- **Redis** (rate limiter)
- **golang-migrate** (миграции)
- **go-playground/validator** (валидация)

---

## Структура проекта

```
cmd/server/main.go              — точка входа
internal/
  adapters/
    postgres.go                 — репозиторий (SQL запросы)
    redis.go                    — Redis клиент (rate limiter)
  connectors/
    routes.go                   — маршруты
    handler.go                  — хендлер
    create_note.go              — POST   /api/notes
    get_note.go                 — GET    /api/notes/{id}
    list_notes.go               — GET    /api/notes
    list_notes_by_search.go     — GET    /api/notes/search?q=
    update_note.go              — PUT    /api/notes/{id}
    delete_note.go              — DELETE /api/notes/{id}
    export_note.go              — GET    /api/notes/{id}/export
    rate_limit_middleware.go    — middleware rate limiter
  domain/
    domain.go                   — модель Note, валидация
    errors.go                   — кастомные ошибки
  dto/                          — входные/выходные структуры
  service/                      — бизнес-логика
pkg/
  migrations/                   — SQL миграции
  httpserver/                   — обёртка над http.Server
  render/                       — JSON рендер
```

---

## Эндпоинты


| Метод | Путь                 | Описание                                                |
| ---------- | ------------------------ | --------------------------------------------------------------- |
| `POST`     | `/api/notes`             | Создать заметку                                   |
| `GET`      | `/api/notes`             | Список (пагинация:`?page=&limit=`)               |
| `GET`      | `/api/notes/{id}`        | Получить одну заметку                        |
| `PUT`      | `/api/notes/{id}`        | Обновить заметку                                 |
| `DELETE`   | `/api/notes/{id}`        | Удалить заметку                                   |
| `GET`      | `/api/notes/search?q=`   | Полнотекстовый поиск по заголовку |
| `GET`      | `/api/notes/{id}/export` | Скачать заметку как`.md` файл              |

## Примеры запросов

```bash
# Создать заметку
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "Go interfaces", "content": "# Interface\nAn interface is...", "tags": "go,basics"}'

# Список заметок
curl "http://localhost:8080/api/notes?page=1&limit=10"

# Поиск
curl "http://localhost:8080/api/notes/search?q=interface"

# Скачать как .md
curl -O http://localhost:8080/api/notes/1/export
```

---

## Rate Limiter

Каждый клиент ограничен **20 запросами в минуту**. При превышении возвращается `429 Too Many Requests` с заголовком `Retry-After`.

---

## Архитектура

Проект построен по принципу **Clean Architecture**:

- **domain** — модель данных и бизнес-правила (не зависит от внешних пакетов)
- **service** — бизнес-логика, работает через интерфейсы
- **adapters** — конкретные реализации (PostgreSQL, Redis)
- **connectors** — HTTP хендлеры, маршрутизация, middleware
- **dto** — структуры для передачи данных между слоями
