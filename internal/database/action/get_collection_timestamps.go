package action

import (
	"context"
	"database/sql"
	"time"
)

type GetCollectionTimestamps struct {
	pool *sql.DB
	sql  string
}

func NewGetCollectionTimestamps(pool *sql.DB) *GetCollectionTimestamps {
	return &GetCollectionTimestamps{
		pool: pool,
		sql: `select t2.name as name,t1.modified as modified 
from user_collections t1
inner join collections t2 on t2.id=t1.collection_id AND t1.user_id = ?
order by t2.id;`,
	}
}

func (p *GetCollectionTimestamps) Run(ctx context.Context, uid uint64) (map[string]time.Time, error) {
	type Result struct {
		Name     string    `db:"name"`
		Modified time.Time `db:"modified"`
	}
	conn, connError := p.pool.Conn(ctx)
	if connError != nil {
		return nil, connError
	}
	defer conn.Close()

	rows, rowsError := conn.QueryContext(ctx, p.sql, uid)
	if rowsError != nil {
		return nil, rowsError
	}
	defer rows.Close()

	out := make(map[string]time.Time)
	for rows.Next() {
		var tmp Result
		if scanError := rows.Scan(&tmp.Name, &tmp.Modified); scanError != nil {
			return nil, scanError
		}
		out[tmp.Name] = tmp.Modified
	}
	return out, nil
}
