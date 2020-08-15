package gateway

import "CleanArchitecture-Go-web/entity"

var dbImage = map[int]*entity.Image{}

func (r *fileRepository) MapUidToImage(uid int, name string) (string, error) {
	dbImage[uid] = &entity.Image{name}
	return dbImage[uid].Name, nil
}
