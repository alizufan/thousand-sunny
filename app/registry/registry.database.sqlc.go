package registry

import (
	"github.com/Mind2Screen-Dev-Team/thousand-sunny/gen/repo"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

var (
	RepoGenerationSqlc = fx.Options(
		fx.Module("dependency:database:sqlc:repo",
			fx.Provide(func(db *pgxpool.Pool) repo.DBTX {
				return db
			}),
			fx.Provide(repo.New),
		),
	)
)
