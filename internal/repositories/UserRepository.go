package repositories

import (
	"AuthBeatsPro/internal/db"
	"AuthBeatsPro/internal/models"
	"database/sql"
	"fmt"
	"time"
)

type UserRepository struct {
	store db.Store
}

func NewUserRepository(store db.Store) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (userRepository *UserRepository) CreateUser(user *models.User) (int, error) {
	var id int64
	sql := "INSERT INTO %s " +
		"(`login`, `password`, `email`, `name`, `lastName`, `type`, `dateOfBirth`, " +
		"`isDeleted`, `isBanned`, `isConfirmed`) " +
		"VALUES " +
		"('?', '?', '?', '?', '?', ?, '?', ?, ?, ?);"

	query := fmt.Sprintf(sql, USERS_TABLE)
	result, err := userRepository.store.Exec(query,
		user.Login,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
		user.Type,
		user.DateOfBirth.Format("YYYY-MM-DD"),
		user.IsDeleted,
		user.IsBanned,
		user.IsConfirmed)

	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (userRepository *UserRepository) UpdateUser(user *models.User) error {
	sql := "UPDATE %s SET " +
		"`login` = '?', " +
		"`password` = '?', " +
		"`email` = '?', " +
		"`name` = '?', " +
		"`lastName` = '?', " +
		"`type` = ?, " +
		"`dateTimeCreation` = '?', " +
		"`dateOfBirth` = '?', " +
		"`isDeleted` = ?, " +
		"`isBanned` = ?, " +
		"`isConfirmed` = ? " +
		"WHERE `id` = ?;"

	query := fmt.Sprintf(sql, USERS_TABLE)

	_, err := userRepository.store.Exec(query,
		user.Login,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
		user.Type,
		user.DateTimeCreation,
		user.DateOfBirth,
		user.IsDeleted,
		user.IsBanned,
		user.IsConfirmed)

	if err != nil {
		return err
	}

	return nil
}

func (userRepository *UserRepository) GetById(id int) (*models.User, error) {
	sql := "SELECT " +
		"`id`, `login`, `password`, `email`, `name`, `lastName`, `type`, " +
		"`dateTimeCreation`, `dateOfBirth`, `isDeleted`, `isBanned`, `isConfirmed` " +
		"FROM %s " +
		"WHERE id = ?;"

	query := fmt.Sprintf(sql, USERS_TABLE)
	row := userRepository.store.QueryRow(query, id)
	user, err := userRepository.makeFromRow(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) GetByCredentials(email string, password string) (*models.User, error) {
	sql := "SELECT " +
		"`id`, `login`, `password`, `email`, `name`, `lastName`, `type`, " +
		"`dateTimeCreation`, `dateOfBirth`, `isDeleted`, `isBanned`, `isConfirmed` " +
		"FROM %s " +
		"WHERE `email` = ? AND `password` = ?;"

	query := fmt.Sprintf(sql, USERS_TABLE)
	row := userRepository.store.QueryRow(query, email, password)

	user, err := userRepository.makeFromRow(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) makeFromRow(row *sql.Row) (*models.User, error) {
	user := models.NewUser()
	var dateOfBirth string
	var dateTimeCreation string
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.LastName,
		&user.Type,
		&dateTimeCreation,
		&dateOfBirth,
		&user.IsDeleted,
		&user.IsBanned,
		&user.IsConfirmed)

	if err != nil {
		return nil, err
	}

	user.DateOfBirth, _ = time.Parse("YYYY-MM-DD", dateOfBirth)
	user.DateTimeCreation, _ = time.Parse("YYYY-MM-DD", dateTimeCreation)

	return user, nil
}
