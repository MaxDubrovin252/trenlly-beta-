package repository

import (
	"fmt"
	"strings"
	"trenlly"

	"github.com/jmoiron/sqlx"
)

type TrenRepository struct {
	db *sqlx.DB
}

func NewTrenRepos(db *sqlx.DB) *TrenRepository {
	return &TrenRepository{db: db}
}

func (r *TrenRepository) CreateTren(tren trenlly.Trening, userId int) (int, error) {

	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (time ,weight,height,main_group) VALUES($1,$2,$3,$4) RETURNING id", trenTable)
	var id int

	row := tx.QueryRow(query, tren.Time, tren.Weight, tren.Height, tren.MainGroup)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	queryE := fmt.Sprintf("INSERT INTO %s(user_id,tren_id) VALUES($1,$2) ", UsersList)

	_, err = tx.Exec(queryE, userId, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TrenRepository) GetAll(userId int) ([]trenlly.Trening, error) {
	var trens []trenlly.Trening

	query := fmt.Sprintf("SELECT tt.time,tt.main_group,tt.weight,tt.height 	FROM %s tt INNER JOIN %s ul ON tt.id = ul.id WHERE user_id=$1", trenTable, UsersList)

	err := r.db.Select(&trens, query, userId)

	if err != nil {
		return nil, err
	}

	return trens, nil
}

func (r *TrenRepository) GetById(userId, trenId int) (trenlly.Trening, error) {
	var tren trenlly.Trening

	query := fmt.Sprintf(`SELECT tt.time,tt.weight,tt.height, tt.main_group FROM %s tt 
	INNER JOIN %s ul ON tt.id = ul.id WHERE ul.tren_id=$1 AND ul.user_id= $2`, trenTable, UsersList)

	err := r.db.Get(&tren, query, trenId, userId)

	return tren, err
}

func (r *TrenRepository) UpdateTren(userId, trenId int, input trenlly.UpdateTrenItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Time != nil {
		setValues = append(setValues, fmt.Sprintf("time=$%d", argId))
		args = append(args, *input.Time)
		argId++
	}

	if input.Weight != nil {
		setValues = append(setValues, fmt.Sprintf("weight=$%d", argId))
		args = append(args, *input.Weight)
		argId++
	}

	if input.Height != nil {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argId))
		args = append(args, *input.Height)
		argId++
	}

	if input.MainGroup != nil {
		setValues = append(setValues, fmt.Sprintf("main_group=$%d", argId))
		args = append(args, *input.MainGroup)
		argId++
	}

	setQuery := strings.Join(setValues, " ")

	query := fmt.Sprintf("UPDATE %s tt SET %s FROM %s ul WHERE tt.id = ul.tren_id AND ul.tren_id=$%d AND ul.user_id=$%d",
		trenTable, setQuery, UsersList, argId, argId+1)
	args = append(args, trenId, userId)

	_, err := r.db.Exec(query, args...)
	return err

}

func (r *TrenRepository) Delete(userId, trenId int) error {
	query := fmt.Sprintf("DELETE FROM %s tt USING %s ul WHERE tt.id = ul.tren_id AND ul.user_id = $1 AND ul.tren_id=$2", trenTable, UsersList)

	_, err := r.db.Exec(query, userId, trenId)

	if err != nil {
		return err
	}

	return nil
}
