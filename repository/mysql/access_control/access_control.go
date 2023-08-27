package mysqlaccesscontrol

import (
	"game-app/entity"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"game-app/pkg/slice"
	"game-app/repository/mysql"
	"strings"
	"time"
)

func (d *DB) GetUserPermissionsTitle(userID uint, role entity.Role) ([]entity.PermissionTitle, error) {
	const op = "mysql.GetUserPermissionsTitle"

	rows, err := d.conn.Conn().Query(`select * from access_control where actor_type = ? and actor_id = ?`, entity.RoleActorType, role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer rows.Close()

	roleACL := make([]entity.AccessControl, 0)

	for rows.Next() {
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		roleACL = append(roleACL, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	userRows, err := d.conn.Conn().Query(`select * from access_control where actor_type = ? and actor_id = ?`, entity.UserActorType, userID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer userRows.Close()

	userACL := make([]entity.AccessControl, 0)

	for userRows.Next() {
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		userACL = append(userACL, acl)
	}

	if err := userRows.Err(); err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	// merge acls by permission id
	permissionIDs := make([]uint, 0)
	for _, r := range roleACL {
		if !slice.DoesExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}

	if len(permissionIDs) == 0 {
		return nil, nil
	}

	args := make([]any, len(permissionIDs))
	for i, id := range permissionIDs {
		args[i] = id
	}
	query := "select * from permission where id in (?" +
		strings.Repeat(",?", len(permissionIDs)-1) +
		")"

	pRows, err := d.conn.Conn().Query(query, args...)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer pRows.Close()

	permissionTitles := make([]entity.PermissionTitle, 0)

	for pRows.Next() {
		permission, err := scanPermission(pRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	return permissionTitles, nil
}

func scanAccessControl(scanner mysql.Scanner) (entity.AccessControl, error) {
	var acl entity.AccessControl
	var createdAt time.Time
	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)

	return acl, err
}
