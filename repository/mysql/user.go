package mysql

import (
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"time"
)

func (d MysqlDb) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantQueryResult).WithKind(richerror.KindUnexpected)
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

func (d MysqlDb) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgNotFound)
		}

		// TODO - log unexpected error for better observability
		return entity.User{}, richerror.New(op).WithErr(err).
			WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgCantQueryResult)
	}

	return user, nil
}

func (d MysqlDb) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"

	row := d.db.QueryRow(`select * from users where id = ?`, userID)
	// TODO - use a function for scan user
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	user := entity.User{}
	var createdAt time.Time
	if err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password); err != nil {
		return entity.User{}, err
	}

	return user, nil
}
