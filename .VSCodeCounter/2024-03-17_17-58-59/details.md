# Details

Date : 2024-03-17 17:58:59

Directory /Users/lightbeag2/go/pkg/vktest

Total : 49 files,  8259 codes, 225 comments, 811 blanks, all 9295 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [Dockerfile](/Dockerfile) | Docker | 8 | 0 | 1 | 9 |
| [Makefile](/Makefile) | Makefile | 5 | 0 | 0 | 5 |
| [cmd/main.go](/cmd/main.go) | Go | 64 | 8 | 9 | 81 |
| [cover.html](/cover.html) | HTML | 2,613 | 0 | 287 | 2,900 |
| [docker-compose.yaml](/docker-compose.yaml) | YAML | 19 | 0 | 1 | 20 |
| [docs/docs.go](/docs/docs.go) | Go | 581 | 2 | 5 | 588 |
| [docs/swagger.json](/docs/swagger.json) | JSON | 563 | 0 | 0 | 563 |
| [docs/swagger.yaml](/docs/swagger.yaml) | YAML | 362 | 0 | 1 | 363 |
| [go.mod](/go.mod) | Go Module File | 47 | 0 | 4 | 51 |
| [go.sum](/go.sum) | Go Checksum File | 177 | 0 | 1 | 178 |
| [internal/config/config.yaml](/internal/config/config.yaml) | YAML | 8 | 0 | 0 | 8 |
| [internal/errs/error.go](/internal/errs/error.go) | Go | 10 | 0 | 4 | 14 |
| [internal/errs/error_test.go](/internal/errs/error_test.go) | Go | 13 | 0 | 4 | 17 |
| [internal/handler/actors.go](/internal/handler/actors.go) | Go | 101 | 39 | 15 | 155 |
| [internal/handler/actors_test.go](/internal/handler/actors_test.go) | Go | 264 | 0 | 23 | 287 |
| [internal/handler/auth.go](/internal/handler/auth.go) | Go | 165 | 40 | 12 | 217 |
| [internal/handler/auth_test.go](/internal/handler/auth_test.go) | Go | 301 | 0 | 27 | 328 |
| [internal/handler/cinema.go](/internal/handler/cinema.go) | Go | 145 | 48 | 11 | 204 |
| [internal/handler/cinema_test.go](/internal/handler/cinema_test.go) | Go | 309 | 0 | 12 | 321 |
| [internal/handler/handler.go](/internal/handler/handler.go) | Go | 26 | 0 | 6 | 32 |
| [internal/handler/handler_test.go](/internal/handler/handler_test.go) | Go | 11 | 0 | 4 | 15 |
| [internal/handler/middleware.go](/internal/handler/middleware.go) | Go | 63 | 0 | 8 | 71 |
| [internal/handler/middleware_test.go](/internal/handler/middleware_test.go) | Go | 110 | 0 | 9 | 119 |
| [internal/migrate/000001_init_mg.down.sql](/internal/migrate/000001_init_mg.down.sql) | SQL | 5 | 0 | 0 | 5 |
| [internal/migrate/000001_init_mg.up.sql](/internal/migrate/000001_init_mg.up.sql) | SQL | 32 | 0 | 7 | 39 |
| [internal/models/models.go](/internal/models/models.go) | Go | 52 | 0 | 9 | 61 |
| [internal/repository/actor_repository.go](/internal/repository/actor_repository.go) | Go | 89 | 0 | 11 | 100 |
| [internal/repository/actor_repository_test.go](/internal/repository/actor_repository_test.go) | Go | 210 | 0 | 21 | 231 |
| [internal/repository/auth_postgres.go](/internal/repository/auth_postgres.go) | Go | 48 | 0 | 9 | 57 |
| [internal/repository/auth_postgres_test.go](/internal/repository/auth_postgres_test.go) | Go | 229 | 0 | 32 | 261 |
| [internal/repository/cinema_repository.go](/internal/repository/cinema_repository.go) | Go | 172 | 0 | 15 | 187 |
| [internal/repository/cinema_repository_test.go](/internal/repository/cinema_repository_test.go) | Go | 210 | 0 | 33 | 243 |
| [internal/repository/mocksr/mockr.go](/internal/repository/mocksr/mockr.go) | Go | 202 | 43 | 44 | 289 |
| [internal/repository/repository.go](/internal/repository/repository.go) | Go | 42 | 1 | 8 | 51 |
| [internal/repository/repository_test.go](/internal/repository/repository_test.go) | Go | 15 | 0 | 6 | 21 |
| [internal/server/server.go](/internal/server/server.go) | Go | 21 | 0 | 5 | 26 |
| [internal/service/actor_service.go](/internal/service/actor_service.go) | Go | 54 | 0 | 8 | 62 |
| [internal/service/actor_service_test.go](/internal/service/actor_service_test.go) | Go | 144 | 0 | 28 | 172 |
| [internal/service/auth_service.go](/internal/service/auth_service.go) | Go | 74 | 0 | 14 | 88 |
| [internal/service/cinema_service.go](/internal/service/cinema_service.go) | Go | 148 | 0 | 10 | 158 |
| [internal/service/cinema_service_test.go](/internal/service/cinema_service_test.go) | Go | 192 | 0 | 27 | 219 |
| [internal/service/mocks/mock.go](/internal/service/mocks/mock.go) | Go | 202 | 43 | 44 | 289 |
| [internal/service/service.go](/internal/service/service.go) | Go | 35 | 1 | 7 | 43 |
| [internal/service/service_test.go](/internal/service/service_test.go) | Go | 17 | 0 | 6 | 23 |
| [internal/utils/utils.go](/internal/utils/utils.go) | Go | 13 | 0 | 4 | 17 |
| [internal/utils/utils_test.go](/internal/utils/utils_test.go) | Go | 13 | 0 | 4 | 17 |
| [pkg/migr.go](/pkg/migr.go) | Go | 16 | 0 | 4 | 20 |
| [pkg/newdb.go](/pkg/newdb.go) | Go | 25 | 0 | 5 | 30 |
| [pkg/newdb_test.go](/pkg/newdb_test.go) | Go | 34 | 0 | 6 | 40 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)