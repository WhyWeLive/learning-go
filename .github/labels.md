# Справочник GitHub labels

Создаются через `gh label create`. Цвета — рекомендация для единообразия.

## Тип

| Label | Color | Описание |
|---|---|---|
| `epic` | `#6F42C1` | Эпик / фаза обучения |
| `task` | `#0366D6` | Атомарная задача |
| `chore-learning` | `#BFD4F2` | Учебный chore без продуктовой фичи |

## Track

| Label | Color | Описание |
|---|---|---|
| `track:lang` | `#0E8A16` | Язык Go |
| `track:db` | `#1D76DB` | PostgreSQL / persistence |
| `track:api` | `#5319E7` | HTTP API |
| `track:auth` | `#D93F0B` | AuthN/AuthZ |
| `track:cache` | `#FBCA04` | Redis / кэш |
| `track:queue` | `#006B75` | Очереди / messaging |
| `track:docker` | `#0052CC` | Docker / compose |
| `track:observability` | `#C5DEF5` | Logs / metrics / tracing |
| `track:highload` | `#B60205` | Highload-паттерны |

## Статус

| Label | Color | Описание |
|---|---|---|
| `status:ready` | `#C2E0C6` | Готово к взятию |
| `status:in-progress` | `#FEF2C0` | В работе (есть ветка) |
| `status:blocked` | `#E99695` | Ждёт другую Issue / решение |
| `status:review` | `#D4C5F9` | Открыт PR на review |

## Размер

| Label | Color | Описание |
|---|---|---|
| `size:S` | `#EDEDED` | ~1–2 часа |
| `size:M` | `#CCCCCC` | ~2–4 часа |
| `size:L` | `#999999` | Слишком крупно — дробить |

## Bootstrap-команды

```bash
gh label create "epic" --color "6F42C1" --description "Эпик / фаза обучения" --force
gh label create "task" --color "0366D6" --description "Атомарная задача" --force
gh label create "chore-learning" --color "BFD4F2" --description "Учебный chore" --force

gh label create "track:lang" --color "0E8A16" --description "Язык Go" --force
gh label create "track:db" --color "1D76DB" --description "PostgreSQL / persistence" --force
gh label create "track:api" --color "5319E7" --description "HTTP API" --force
gh label create "track:auth" --color "D93F0B" --description "AuthN/AuthZ" --force
gh label create "track:cache" --color "FBCA04" --description "Redis / кэш" --force
gh label create "track:queue" --color "006B75" --description "Очереди / messaging" --force
gh label create "track:docker" --color "0052CC" --description "Docker / compose" --force
gh label create "track:observability" --color "C5DEF5" --description "Logs / metrics / tracing" --force
gh label create "track:highload" --color "B60205" --description "Highload-паттерны" --force

gh label create "status:ready" --color "C2E0C6" --description "Готово к взятию" --force
gh label create "status:in-progress" --color "FEF2C0" --description "В работе" --force
gh label create "status:blocked" --color "E99695" --description "Заблокировано" --force
gh label create "status:review" --color "D4C5F9" --description "На review" --force

gh label create "size:S" --color "EDEDED" --description "~1-2 часа" --force
gh label create "size:M" --color "CCCCCC" --description "~2-4 часа" --force
gh label create "size:L" --color "999999" --description "Дробить" --force
```
