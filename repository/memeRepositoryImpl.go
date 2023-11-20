package repository

import "my-meme/model"

type MemeRepoImpl[T model.Meme] struct {
}

func (r *MemeRepoImpl[Meme]) GetAll() ([]model.Meme, error) {
	meme := []model.Meme{}
	return meme, nil
}

func (r *MemeRepoImpl[Meme]) GetByID(id int) (*model.Meme, error) {
	meme := new(model.Meme)
	return meme, nil
}

func (r *MemeRepoImpl[Meme]) Create(t model.Meme) (*model.Meme, error) {
	meme := new(model.Meme)
	return meme, nil
}

func (r *MemeRepoImpl[Meme]) Update(t model.Meme) (*model.Meme, error) {
	meme := new(model.Meme)
	return meme, nil
}

func (r *MemeRepoImpl[Meme]) Delete(id int) error {
	return nil
}
