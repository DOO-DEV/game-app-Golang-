package mysqlanswer

import (
	"context"
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"time"
)

func (d DB) InsertAnswers(ctx context.Context, answers []entity.PossibleAnswer) error {
	const op = "mysql.InsertAnswers"

	var query string
	for idx, ans := range answers {
		query += fmt.Sprintf(`("%s", %d)`, ans.Text, ans.QuestionID)
		if idx == len(answers)-1 {
			query += ";"
		} else {
			query += ","
		}
	}
	query = `insert into answers(text, question_id) values ` + query
	fmt.Println(query)
	_, err := d.conn.Conn().ExecContext(ctx, query)
	if err != nil {
		fmt.Println("mysql error", err)
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
func (d DB) UpdateAnswer(ctx context.Context, answer entity.PossibleAnswer) (entity.PossibleAnswer, error) {
	const op = "mysql.UpdateAnswer"

	_, err := d.conn.Conn().ExecContext(ctx, `update answers set text = ?, question_id = ? where id = ?`, answer.Text, answer.QuestionID, answer.ID)
	if err != nil {
		return entity.PossibleAnswer{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	row := d.conn.Conn().QueryRowContext(ctx, `select * from answers where id = ?`, answer.ID)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return entity.PossibleAnswer{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return entity.PossibleAnswer{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	var ans entity.PossibleAnswer
	var createdAt time.Time
	var choice uint
	if err := row.Scan(&ans.ID, &ans.Text, &createdAt, &choice, &ans.QuestionID); err != nil {
		return entity.PossibleAnswer{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return ans, nil
}
func (d DB) DeleteAnswer(ctx context.Context, id uint) error {
	const op = "mysql.DeleteAnswer"

	if _, err := d.conn.Conn().ExecContext(ctx, `delete from answers where question_id = ?`, id); err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) GetAnswers(ctx context.Context, id uint) ([]entity.PossibleAnswer, error) {
	const op = "mysql.GetAnswers"

	rows, err := d.conn.Conn().QueryContext(ctx, `select * from answers where question_id = ?`, id)
	if err != nil {
		return []entity.PossibleAnswer{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()
	answers := make([]entity.PossibleAnswer, 0)
	for rows.Next() {
		ans := entity.PossibleAnswer{}
		var createdAt time.Time
		if err := rows.Scan(&ans.ID, &ans.Text, &createdAt, &ans.QuestionID); err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		answers = append(answers, ans)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return answers, nil
}
