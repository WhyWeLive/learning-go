# Профиль ученика (живой снимок)

Обновлять после собесов и когда Teacher/PM/Reviewer выявляют устойчивые пробелы.  
**Не путать с желаемым уровнем Senior** — цель в `mission.md`, факт здесь.

## База

- Коммерческий бэкенд: ~1 год; TeamLead: ~1 год (календарь ≠ глубина).
- Go: **0** (на старте трека).
- Цель: полный путь языка → infra → highload (marketplace).
- Стек опыта: NestJS / Node; абстракции есть точечно, не «Senior во всём».

## Сильные стороны

- Nest pipeline: guard / pipe / interceptor — направление верное.
- Auth база: access vs refresh, secure storage на мобиле.
- N+1: суть понимает; unit-тесты с mock repository.
- Честность по пробелам (не выдумывает Kafka/Redis на собесе).
- Интуиция Go zero value / `error` как ценность — при нулевом опыте.

## Дыры (приоритет обучения)

- Postgres concurrency: `FOR UPDATE`, optimistic locking, резерв стока без ухода в минус.
- Redis: cache-aside, stampede, locks, rate limit.
- Kafka / outbox / at-least-once / идемпотентный consumer.
- Observability: logs vs metrics, разбор роста p99.
- Highload: пулы соединений, backpressure, скейл.
- Docker runtime: PID 1 / сигналы, liveness vs readiness.
- Миграции expand/contract.
- TL-процесс: severity в review, системная декомпозиция задач.

## Критичные misconceptions

1. **Идемпотентность** — путал с «нельзя повторить»; учить корректное определение на заказах/retry.
2. **Timeout** — считал плохой практикой; в Go-треке закреплять `context` + дедлайны как норму.
3. **CORS** — формулировка как «защита сервера»; поправлять при касании web/API.

## Следствия для агента

- **Не считать** Redis / Kafka / Postgres concurrency / highload «уже знаешь — только синтаксис Go».
- В каждой Issue — полноценный Learning-блок (см. `pm.md`), даже на «знакомых» абстракциях.
- Дробить задачи мельче; в рисках Issue явно concurrency / idempotency / timeouts где уместно.
- Teacher: лестница L0–L3; опираться на этот файл, не на завышенные ярлыки.
