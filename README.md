# pr-reviews-assignment-service

## Setup and run

- To use service just run
```bash
docker compose up -d
```
- After that you may test app by using swagger which is available on that [link](http://0.0.0.0:8080/swagger/index.html)

## Performance tests

- `GET /api/v1/user/stats` is tested

- 3000 PRs, 600 users, 60 teams

- P(99) < 100ms

- [Benchmarks](./docs/performance/user_stats.jpg)

## Useful links
- [Erd](./docs/ERD.md)
