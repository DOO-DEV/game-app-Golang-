package mysqlquestion

import (
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"game-app/repository/mysql"
	"time"
)

type category struct {
	id   uint
	name string
}

func (d DB) GetQuestionByID(id uint) (entity.Question, error) {
	const op = "mysql.GetQuestionByID"

	row := d.conn.Conn().QueryRow(`select * from questions where id = ?`, id)
	question, err := scanQuestion(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Question{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return entity.Question{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return question, nil
}

func (d DB) GetQuestionsByCategory(catID uint) ([]entity.Question, error) {
	const op = "mysql.GetQuestionsByCategory"

	rows, err := d.conn.Conn().Query(`select * from questions where category_id = ?`, catID)
	if err != nil {
		return []entity.Question{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	questions := make([]entity.Question, 0)
	for rows.Next() {
		question, err := scanQuestion(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return questions, nil
}

func (d DB) GetQuestionsByCategoryName(category entity.Category) ([]entity.Question, error) {
	const op = "mysql.GetQuestionsByCategory"

	rows, err := d.conn.Conn().Query(`select q.id, q.question, q.difficulty, q.answer_id, c.name from categories as c inner join questions as q on c.name = ?`, category)
	if err != nil {
		return []entity.Question{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	questions := make([]entity.Question, 0)
	for rows.Next() {
		q := entity.Question{}
		var c string
		rows.Scan(&q.ID, &q.Question, &q.Difficulty, &q.CorrectAnswerID, &c)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		questions = append(questions, q)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).
			WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return questions, nil
}

func (d DB) InsertQuestion(question entity.Question) (entity.Question, error) {
	const op = "mysql.InsertQuestion"

	res, err := d.conn.Conn().Exec(`insert into questions(question, difficulty, answer_id, category_id) values(?, ?, ?, ?)`,
		question.Question, question.Difficulty, question.CorrectAnswerID, question.CategoryID)
	if err != nil {
		fmt.Println("errrr", err)
		return entity.Question{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	id, _ := res.LastInsertId()
	question.ID = uint(id)

	return question, nil
}
func (d DB) UpdateQuestion(q entity.Question) (entity.Question, error) {
	const op = "mysql.UpdateQuestion"

	fmt.Println(q)
	_, err := d.conn.Conn().Exec(`update questions set question = ?, difficulty = ?, answer_id = ?, category_id = ? where id = ?`,
		q.Question, q.Difficulty, q.CorrectAnswerID, q.CategoryID, q.ID)
	if err != nil {
		fmt.Println(err)
		return entity.Question{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	row := d.conn.Conn().QueryRow(`select * from questions where id = ?`, q.ID)
	question, err := scanQuestion(row)

	return question, nil
}

func (d DB) DeleteQuestion(id uint) error {
	const op = "mysql.DeleteQuestion"

	_, err := d.conn.Conn().Exec(`delete from questions where id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	return nil
}

func (d DB) GetCategoryByID(id uint) (entity.Category, error) {
	const op = "mysql.GetQuestionByID"

	row := d.conn.Conn().QueryRow(`select * from categories where id = ?`, id)
	var c category
	err := row.Scan(&c.id, &c.name)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Category(c.name), richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return entity.Category(c.name), richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return entity.Category(c.name), nil
}

func scanQuestion(scanner mysql.Scanner) (entity.Question, error) {
	q := entity.Question{}
	var createdAt time.Time

	if err := scanner.Scan(&q.ID, &q.Question, &q.Difficulty, &createdAt, &q.CorrectAnswerID, &q.CategoryID); err != nil {
		return entity.Question{}, err
	}

	return q, nil
}
