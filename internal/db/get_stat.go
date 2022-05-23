package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func (c *Client) GetStat(ctx context.Context, chatID m.ID) (*m.Statistics, error) {
	const query = `SELECT c.task_number, count(s.result) FILTER ( WHERE s.result='correct' ), count(s.result)
		FROM submissions s
						JOIN problems p ON s.problem_id = p.id
						JOIN categories c ON c.id = p.category_id
		WHERE chat_id = $1
			AND (s.result != 'aborted' AND s.result != 'pending')
		group by c.task_number ORDER BY c.task_number;`

	rows, err := c.pool.Query(ctx, query, chatID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get stat from database, err: <%v>", err)
	}

	stat := m.Statistics{}

	for rows.Next() {
		taskStat := m.TaskStat{}
		err := rows.Scan(&taskStat.TaskNumber, &taskStat.Correct, &taskStat.All)
		if err != nil {
			return nil, fmt.Errorf("failed to scan a row, err: <%v>", err)
		}
		stat.Stat = append(stat.Stat, taskStat)
	}

	return &stat, nil
}
