package mysqlaccesscontrol

import (
	"game-app/entity"
	"game-app/repository/mysql"
	"time"
)

func scanPermission(scanner mysql.Scanner) (entity.Permission, error) {
	var createdAt time.Time
	var permission entity.Permission

	scanner.Scan(&permission.ID, &permission.Title, &createdAt)

	return permission, nil
}
