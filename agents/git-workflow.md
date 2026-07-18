# Git workflow (enterprise для соло)

Одна Issue = одна ветка = один PR. Заголовки Issues/PR — русский; slug ветки — английский.

## Ветки

| Префикс | Когда |
|---|---|
| `feat/#<n>-<slug>` | продуктовая фича |
| `fix/#<n>-<slug>` | баг |
| `chore/#<n>-<slug>` | инфраструктура / tooling / учебный каркас |
| `docs/#<n>-<slug>` | документация |
| `refactor/#<n>-<slug>` | рефакторинг без смены поведения |

Правила:

- `<n>` — номер GitHub Issue.
- `<slug>` — короткий kebab-case **на английском** (`jwt-middleware`, `order-outbox`).
- Ветка от актуального `main`.
- Коммиты напрямую в `main` — запрещены по конвенции.
- При переводе Issue в `status:in-progress` имя ветки **обязательно** в секции Мета Issue.

Пример: Issue `#42` → `feat/#42-jwt-middleware`.

## Коммиты

Conventional Commits:

- `feat:`, `fix:`, `chore:`, `test:`, `docs:`, `refactor:`

В теле или футере:

- `Refs #42` — промежуточная работа;
- `Closes #42` — если один коммит/PR полностью закрывает задачу (обычно достаточно `Closes` в PR).

## Pull Request

- Заголовок: `#42 Краткий заголовок на русском`
- Base: `main`
- 1 PR = 1 Issue; раздувание PR → дробить Issue, не склеивать.

Тело (минимум):

```markdown
## Связь
Closes #42

## Что сделано
…

## Как проверить
…

## Учебный итог
Какие концепции из Issue отработаны; что могу объяснить.

## Риски / известные ограничения
…
```

## Статусы Issue ↔ git

`status:ready` → взять в работу → ветка + `status:in-progress` → PR → `status:review` → merge → **close Issue** (достаточно closed state).

Агент не мержит и не пушит код; review — через `gh pr review` (см. `reviewer.md`).

## Напоминания агенту в review

Проверять: имя ветки содержит `#<n>`, заголовок PR с `#<n>`, в теле есть `Closes #<n>` или эквивалент.
