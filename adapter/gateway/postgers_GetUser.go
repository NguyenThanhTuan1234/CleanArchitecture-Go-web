package gateway

import "CleanArchitecture-Go-web/entity"

func (r *postgresRepository) GetUser(un string) (*entity.User, error) {
	var u entity.User
	rows, err := r.client.db.Query("SELECT username, password, first, last, role FROM test1 WHERE username = $1", un)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&u.UserName, &u.Password, &u.First, &u.Last, &u.Role)
		if err != nil {
			panic(err)
		}
	}
	return &u, nil
}
