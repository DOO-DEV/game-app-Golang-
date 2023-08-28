package mysqlanswer

import (
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"time"
)

func (d DB) InsertAnswers(answers []entity.PossibleAnswer) error {
	const op = "mysql.InsertAnswers"

	var query string
	for _, ans := range answers {
		query += fmt.Sprintf(`insert into answers(text, quesion_id) values(%s, %d);`, ans.Text, ans.QuestionID)
	}

	_, err := d.conn.Conn().Exec(query)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}
func (d DB) UpdateAnswer(answer entity.PossibleAnswer) (entity.PossibleAnswer, error) {
	const op = "mysql.UpdateAnswer"

	_, err := d.conn.Conn().Exec(`update table answers set text = ?, question_id = ? where id = ?`, answer.Text, answer.QuestionID, answer.ID)
	if err != nil {
		return entity.PossibleAnswer{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	row := d.conn.Conn().QueryRow(`select * from answers where id = ?`, answer.ID)
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
func (d DB) DeleteAnswer(id uint) error {
	const op = "mysql.DeleteAnswer"

	if _, err := d.conn.Conn().Exec(`delete from answers where id = ?`, id); err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	
	return nil
}
