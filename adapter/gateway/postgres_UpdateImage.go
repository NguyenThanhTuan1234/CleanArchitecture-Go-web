package gateway

func (r *postgresRepository) UpdateImage(imgName string, uid int) error {
	_, err := r.client.db.Exec("UPDATE image SET name = $1 WHERE uid = $2", imgName, uid)
	if err != nil {
		return err
	}
	return nil
}
