package mysqluser

import (
	"context"
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"game-app/repository/mysql"
	"time"
)

func (d DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"

	row := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
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

func (d DB) Register(u entity.User) (entity.User, error) {
	res, err := d.conn.Conn().Exec(`insert into  users(name, phone_number, password, role) values(?, ?, ?, ?)`, u.Name, u.PhoneNumber, u.Password, u.Role.String())
	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command: %w", err)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (d DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.conn.Conn().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
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

func (d DB) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"

	row := d.conn.Conn().QueryRowContext(ctx, `select * from users where id = ?`, userID)
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

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	user := entity.User{}
	var createdAt time.Time
	var roleStr string
	if err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr); err != nil {
		return entity.User{}, err
	}

	user.Role = entity.MapToEntityRole(roleStr)

	return user, nil
}
