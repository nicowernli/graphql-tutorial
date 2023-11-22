package repositories

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/nicowernli/graphql-tutorial/internal/db/models"
	database "github.com/nicowernli/graphql-tutorial/pkg/db/mysql"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUsers(query *models.UserQuery) ([]*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

func NewUserRepository() UserRepository {
	return &UserMysqlRepository{}
}

type UserMysqlRepository struct{}

func (r *UserMysqlRepository) CreateUser(user *models.User) error {
	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC().Unix()
	user.UpdatedAt = user.CreatedAt

	stmt, err := database.Db.Prepare("INSERT INTO Users(ID, Email, FirstName, LastName, CreatedAt, UpdatedAt) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(user.ID.String(), user.Email, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *UserMysqlRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	stmt, err := database.Db.Prepare("SELECT ID, Email, FirstName, LastName, CreatedAt, UpdatedAt FROM Users WHERE Email = ? AND RemovedAt IS NULL")
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserMysqlRepository) FindUsers(query *models.UserQuery) ([]*models.User, error) {
	var users []*models.User

	strStmt := "SELECT ID, Email, FirstName, LastName, CreatedAt, UpdatedAt FROM Users WHERE RemovedAt IS NULL"
	if query.OrderBy != nil && query.Sort != nil {
		strStmt += " ORDER BY " + *query.OrderBy + " " + *query.Sort
	}

	stmt, err := database.Db.Prepare(strStmt)
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	if query.OrderBy == nil && query.Sort == nil {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(query.OrderBy, query.Sort)
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
