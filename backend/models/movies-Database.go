package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date,rating, runtime, created_at, updated_at from movies where id = $1 `

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	query = `select
	           mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
			   movies_genres mg
			   left join genres g on (g.id = mg.genre_id)
			where
			    mg.movie_id = $1
	`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.Genre.ID,
			&mg.Genre.GenreName,
		)

		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres
	return &movie, nil
}

func (m *DBModel) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date,rating, runtime, created_at, updated_at from movies order by title`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Rating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		genreQuery := `select
		mg.id, mg.movie_id, mg.genre_id, g.genre_name
	 from
		movies_genres mg
		left join genres g on (g.id = mg.genre_id)
	 where
		 mg.movie_id = $1
`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.Genre.ID,
				&mg.Genre.GenreName,
			)

			if err != nil {
				return nil, err
			}
			genres[mg.ID] = mg.Genre.GenreName
		}

		genreRows.Close()

		movie.MovieGenre = genres

		movies = append(movies, &movie)

	}
	return movies, nil
}

func (m *DBModel) GenresAll() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from genres order by genre_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var g Genre
		err := rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAt,
			&g.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, &g)
	}

	return genres, nil
}
