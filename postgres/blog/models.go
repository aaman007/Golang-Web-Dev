package blog

import "github.com/aaman007/Golang-Web-Dev/postgres/config"

type Blog struct {
	Id int
	Title string
	Body string
}

func All() ([]Blog, error) {
	rows, err := config.DB.Query("SELECT * from blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	blogs := make([]Blog, 0)
	for rows.Next() {
		blog := Blog{}
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Body)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

func Get(blogId int) (Blog, error) {
	blog := Blog{}

	row := config.DB.QueryRow("SELECT * from blogs WHERE id=$1", blogId)
	err := row.Scan(&blog.Id, &blog.Title, &blog.Body)
	if err != nil {
		return blog, err
	}

	return blog, nil
}

func Create(title, body string) (Blog, error) {
	blog := Blog{}

	_, err := config.DB.Exec("INSERT INTO blogs (title, body) VALUES ($1, $2)", title, body)
	if err != nil {
		return blog, err
	}

	return blog, nil
}

func Update(title, body string, blogId int) (Blog, error) {
	blog := Blog{}

	_, err := config.DB.Exec("UPDATE blogs SET title=$1, body=$2 WHERE id=$3", title, body, blogId)
	if err != nil {
		return blog, err
	}

	return blog, nil
}

func Delete(blogId int) error {
	_, err := config.DB.Exec("DELETE FROM blogs WHERE id=$1", blogId)
	if err != nil {
		return err
	}
	return nil
}