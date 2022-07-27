package database

import "restapi/app/models"

func (d *DB) CreatePost(p *models.Post) error {
	res, err := d.db.Exec(insertPostSchema, p.Title, p.Content, p.Author)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := d.db.Select(&posts, "SELECT * FROM posts")

	if err != nil {
		return posts, err
	}

	return posts, nil

}

func (d *DB) DeletePost(id string) error {
	_, err := d.db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetPostById(id string) (*models.Post, error) {
	var post models.Post
	err := d.db.Get(&post, "SELECT * FROM posts WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
