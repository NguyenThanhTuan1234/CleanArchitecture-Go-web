package gateway

func (r *postgresRepository) SaveImage(imgName string, uid int) error {
	_, err := r.client.db.Exec("INSERT INTO image (name, uid) VALUES($1, $2)", imgName, uid)
	if err != nil {
		return err
	}
	return nil
}
