package gateway

import "CleanArchitecture-Go-web/entity"

func (r *postgresRepository) GetImage(uid int) (*entity.Image, error) {
	var i entity.Image
	rows, err := r.client.db.Query("SELECT name FROM image WHERE uid = $1", uid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&i.Name)
		if err != nil {
			panic(err)
		}
	}
	return &i, nil
}
