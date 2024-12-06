# SQLite Relational Database

## Overview of Setup Efforts
The process of setting up the SQLite relational database involved creating two main tables: `movies` and `genres`. The `movies` table includes columns for `id`, `name`, `year`, and `rank`. Meanwhile, the `genres` table includes columns for `movie_id` and `genre`, with a foreign key relationship between `movie_id` and the `movies` table.

The database was populated using data from IMDB CSV files (`IMDB-movies.csv` and `IMDB-movies_genres.csv`). The data was inserted into the database using a Go application that employed  `modernc.org/sqlite` . 

### Adding Personal Collection
An additional table can be added to track personal collections of movies. This table might be called `personal_collection`, and it could include the following columns:

- `movie_id` (INTEGER): key linking to the `movies` table.
- `location` (TEXT): A description of where the physical or digital movie is located (e.g., "Living room shelf" or "Netflix account").
- `personal_rating` (FLOAT): A personal rating for each movie.
- `comments` (TEXT): Any additional notes or comments about the movie.

## Future Plans for the Personal Movie Database
The goal of building this personal movie database is to create a customized movie application that can serve a more personalized purpose compared to general movie databases like IMDb. Here are some potential use cases and benefits of this database:

- **Track Personal Collections**: Users can manage their personal collections, including movies they own, where they are located, and personal ratings. This is particularly useful for keeping track of large collections and knowing exactly where a movie is stored.
- **Personal Ratings and Comments**: Users can add their own ratings and comments, helping them remember what they liked or disliked about a particular movie.
- **Personal Watchlists**: The application could allow users to create watchlists or categorize movies into different lists (e.g., "Favorites", "To Watch", "For Family Nights").

## Future Enhancements and Application Development
### Database Enhancements
- **Insert data faster**: Instead of inserting records one-by-one, insert multiple records in a single transaction.
- **Concurrent Insert and Read Operations**: enabling concurrent database access.

## How to Run the Code

1. Clone this repository.

2. Install dependencies using `go get` for the `modernc.org/sqlite` package.

3. Run the main Go program:

   ```
   go run main.go
   ```

4. Wait for the program finish running (It takes 25 minutes...).

## Use the database

1. Download sqlite tool: https://www.sqlite.org/download.html

2. Unzip it.

3. run `sqlite3 .\movies.db`

4. run sql:

   ```
   SELECT g.genre, AVG(m.rank) AS avg_rating
   FROM movies m
   JOIN genres g ON m.id = g.movie_id
   GROUP BY g.genre
   ORDER BY avg_rating DESC;
   ```