package store

import (
	"database/sql"
	"fmt"

	"github.com/Harddancer/rest_api_go/internal/app/model"
)

type UserRepository struct{
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error){
if err  := u.BeforCreate();err != nil{
	return nil,err
}


	if err := r.store.db.QueryRow("INSERT INTO users (email,encrypted_password) VALUES ($1,$2) RETURNING id",
	u.Email,
	u.Password,
	).Scan(&u.ID); err != nil{
		return nil, err
	}
	fmt.Printf("ID новой записи %d",u.ID)
	return u,nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error){

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id,email,encrypted_password FROM users WHERE email = $1",
		email,
		).Scan(
			&u.ID,
			&u.Email,
			&u.Password,
		); err != nil{
			return nil,err
		}

		return u,nil
	
} 


func (r *UserRepository) FindById(id int) (*model.User, error){

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id,email,encrypted_password FROM users WHERE id = $1",
		id,
		).Scan(
			&u.ID,
			&u.Email,
			&u.Password,
		); err != nil{
			if err == sql.ErrNoRows {
				return nil, err
			}

			return nil, err
		}
		return u, nil
	
} 