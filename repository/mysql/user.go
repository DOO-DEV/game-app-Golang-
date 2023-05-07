package mysql

import (
	"database/sql"
	"fmt"
	"game-app/entity"
)

func (d MysqlDb) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("cant scan query result: %w", err)
	}

	return false, nil
}

func (d MysqlDb) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into  users(name, phone_number, password) values(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command: %w", err)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (d MysqlDb) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("cant scan query result: %w", err)
	}

	return user, true, nil
}

func (d MysqlDb) GetUserByID(userID uint) (entity.User, error) {
	row := d.db.QueryRow(`select * from users where id = ?`, userID)
	// TODO - use a function for scan user
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, nil
		}

		return entity.User{}, fmt.Errorf("cant scan query result: %w", err)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	user := entity.User{}
	var createdAt []uint8
	if err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt); err != nil {
		return entity.User{}, err
	}

	return user, nil
}
