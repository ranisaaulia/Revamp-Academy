package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	curi "codeid.revampacademy/models"
	mod "codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
)

const createProgEntityDesc = `-- name: CreateProgEntityDesc :one

INSERT INTO curriculum.program_entity_description (pred_prog_entity_id, 
pred_item_learning, 
pred_description, 
pred_target_level) 
VALUES($1,$2,$3,$4)
RETURNING *
`

type CreateProgEntityDescParams struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
	PredTargetLevel  sql.NullString `db:"pred_target_level" json:"predTargetLevel"`
}

func (q *Queries) CreateProgEntityDesc(ctx context.Context, arg CreateProgEntityDescParams) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	row := q.db.QueryRowContext(ctx, createProgEntityDesc,
		arg.PredProgEntityID,
		arg.PredItemLearning,
		arg.PredDescription,
		arg.PredTargetLevel,
	)
	i := mod.CurriculumProgramEntityDescription{}
	err := row.Scan(
		&i.PredProgEntityID,
		&i.PredItemLearning,
		&i.PredDescription,
		&i.PredTargetLevel,
	)
	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.CurriculumProgramEntityDescription{
		PredProgEntityID: i.PredProgEntityID,
		PredItemLearning: i.PredItemLearning,
		PredDescription:  i.PredDescription,
		PredTargetLevel:  i.PredTargetLevel,
	}, nil
}

const deleteprogram_entity_description = `-- name: Deleteprogram_entity_description :exec
DELETE FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

func (q *Queries) Deleteprogram_entity_description(ctx context.Context, predProgEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteprogram_entity_description, predProgEntityID)
	return err
}

const getprogram_entity_description = `-- name: Getprogram_entity_description :one

SELECT pred_prog_entity_id, pred_item_learning, pred_description, pred_target_level FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

func (q *Queries) Getprogram_entity_description(ctx context.Context, predProgEntityID int32) (curi.CurriculumProgramEntityDescription, error) {
	row := q.db.QueryRowContext(ctx, getprogram_entity_description, predProgEntityID)
	var i curi.CurriculumProgramEntityDescription
	err := row.Scan(
		&i.PredProgEntityID,
		&i.PredItemLearning,
		&i.PredDescription,
		&i.PredTargetLevel,
	)
	return i, err
}

const listprogram_entity_description = `-- name: Listprogram_entity_description :many
SELECT pred_prog_entity_id,
pred_item_learning, 
pred_description, 
pred_target_level FROM curriculum.program_entity_description
ORDER BY pred_prog_entity_id
limit $1 offset $2
`

func (q *Queries) Listprogram_entity_description(ctx context.Context, metadata *features.Metadata) ([]curi.CurriculumProgramEntityDescription, error) {
	rows, err := q.db.QueryContext(ctx, listprogram_entity_description, metadata.PageSize, metadata.PageNo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []curi.CurriculumProgramEntityDescription
	for rows.Next() {
		var i curi.CurriculumProgramEntityDescription
		if err := rows.Scan(
			&i.PredProgEntityID,
			&i.PredItemLearning,
			&i.PredDescription,
			&i.PredTargetLevel,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateprogram_entity_description = `-- name: Updateprogram_entity_description :exec
UPDATE curriculum.program_entity_description
  set pred_item_learning= $2,
  pred_description = $3,
  pred_target_level = $4
WHERE pred_prog_entity_id= $1
`

type UpdateProgEntityDescParams struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
	PredTargetLevel  sql.NullString `db:"pred_target_level" json:"predTargetLevel"`
}

func (q *Queries) UpdateProgEntityDesc(ctx context.Context, arg UpdateProgEntityDescParams) error {
	_, err := q.db.ExecContext(ctx, updateprogram_entity_description, arg.PredProgEntityID, arg.PredItemLearning, arg.PredDescription, arg.PredTargetLevel)
	return err
}

const deleteprogEntityDesc = `-- name: DeleteProgEntityDesc :exec
DELETE FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

func (q *Queries) DeleteProgEntityDesc(ctx context.Context, predProgEntityId int32) error {
	_, err := q.db.ExecContext(ctx, deleteprogEntityDesc, predProgEntityId)
	return err
}
