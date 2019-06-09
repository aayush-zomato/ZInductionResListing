package post

import (
	"context"
	"database/sql"

	models "Zinduction/reslisting/models"
	pRepo  "Zinduction/reslisting/repository"
)

type mysqlPostRepo struct{
	Conn *sql.DB
}

func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) (
	[]*models.Post, error){
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err!=nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Post, 0)
	for rows.Next() {
		data := new(models.Post)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Cusine,
			&data.Distance,
			&data.Address,
			&data.CFT,
			&data.Rating,
			&data.Timing,
		)

		if err!=nil{
			return nil, err
		}
		
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*models.Post, error){
	query := "Select id, name, cusine, distance, address, cft, rating, timing From posts limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlPostRepo) GetByID(ctx context.Context, id int64) (*models.Post, error){
	query := "Select id, name, cusine, distance, address, cft, rating, timing From posts where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Post{}
	if len(rows) > 0 {
		payload = rows[0]
	} else{
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlPostRepo) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "Insert posts SET name=?, cusine=?, distance=?, address=?, cft=?, rating=?, timing=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil{
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.Cusine, p.Distance, p.Address, p.CFT, p.Rating, p.Timing)
	defer stmt.Close()

	if err!=nil{
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlPostRepo) Update(ctx context.Context, p *models.Post) (*models.Post, error){
	query := "Update posts set name=?, cusine=?, distance=?, address=?, cft=?, rating=?, timing=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err!=nil{
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Cusine, 
		p.Distance, 
		p.Address, 
		p.CFT, 
		p.Rating, 
		p.Timing,
	)

	if err!=nil{
		return nil, err
	}
	defer stmt.Close()
	return p, nil
}

func (m *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error){
	query := "Delete from posts Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err!=nil{
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err!=nil{
		return false, err
	}
	return true, nil
}