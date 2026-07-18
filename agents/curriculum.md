# Curriculum — карта тем (0 → Senior / highload)

Порядок — рекомендуемый. Epic = срез marketplace + 1–2 трека. Перед каждым epic — согласование с пользователем.

Прогресс трека = закрытые Issues с соответствующим `track:*`.

## Фазы

| Фаза | Track labels | Фокус | Пример среза marketplace |
|---|---|---|---|
| 0 | `track:lang`, `track:docker` | Окружение, modules, структура репо, базовый CLI/HTTP hello | Каркас сервиса, `go mod`, layout |
| 1 | `track:lang` | Типы, interfaces, errors, generics, concurrency basics | Доменные типы каталога/заказа (без полной БД) |
| 2 | `track:lang` | Testing, table-driven, fuzz (по мере готовности) | Тесты на доменные правила |
| 3 | `track:api` | HTTP router, middleware, config, graceful shutdown | Public API каталога (in-memory или stub) |
| 4 | `track:db` | pgx/sqlc, миграции, транзакции | Persistence каталога / inventory |
| 5 | `track:auth` | JWT/sessions, password hashing, authz | Регистрация/логин продавца и покупателя |
| 6 | `track:cache` | Redis: кэш, locks, rate limit | Кэш карточки товара, lock резерва |
| 7 | `track:docker` | Compose multi-service, healthchecks | Локальный стек API+DB+Redis |
| 8 | `track:queue` | Kafka/NATS, outbox, consumers | Order placed → асинхронная обработка |
| 9 | `track:observability` | structured logs, metrics, tracing | SLO-friendly эндпоинты заказа |
| 10 | `track:highload` | пулы, таймауты, идемпотентность, backpressure | Конкурентное резервирование inventory |

Фазы можно частично перекрывать, но не прыгать в highload без базы языка/API/DB.

## Банк источников (предпочитать официальные)

| Тема | Ссылки |
|---|---|
| Язык | https://go.dev/doc/effective_go — Effective Go |
| Tour | https://go.dev/tour/ — A Tour of Go |
| Spec / modules | https://go.dev/ref/mod — Go Modules Reference |
| Errors | https://go.dev/blog/go1.13-errors — Working with Errors |
| Context | https://pkg.go.dev/context — package context |
| Testing | https://go.dev/doc/tutorial/add-a-test |
| Postgres | https://www.postgresql.org/docs/current/ |
| Redis | https://redis.io/docs/ |
| Docker | https://docs.docker.com/ |
| Kafka | https://kafka.apache.org/documentation/ |

В каждой Issue — **свой** учебный блок с терминами и аннотированными ссылками под задачу; этот файл — карта, не замена Learning-секции.

## Типичные термины по трекам (шпаргалка для Issues)

- **lang:** goroutine, channel, interface satisfaction, zero value, error wrapping, `context.Context`
- **api:** middleware, handler, graceful shutdown, request-scoped context
- **db:** connection pool, migration, isolation level, upsert, `FOR UPDATE`
- **auth:** JWT claims, refresh vs access, bcrypt/argon2, principle of least privilege
- **cache:** cache-aside, TTL, stampede, distributed lock
- **queue:** at-least-once, idempotency key, outbox, consumer lag
- **highload:** backpressure, timeouts, bulkhead, idempotent retry
