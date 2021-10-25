BEGIN;
CREATE TYPE time_span AS ENUM('four weeks', 'six months', 'all time');
COMMIT;
BEGIN;
CREATE TYPE track AS (
  track_name text,
  artist_name text,
  spotify_track_id text,
  track_image_url text,
  track_rank int
);
COMMIT;
CREATE TYPE artist AS (
  artist_name text,
  artist_spotify_id text,
  artist_image_url text,
  artist_rank int
);
BEGIN;
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  spotify_user_id text UNIQUE NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  is_premium boolean NOT NULL,
  updated_at timestamp NOT NULL
);
COMMIT;
BEGIN;
CREATE TABLE tracks (
  id SERIAL PRIMARY KEY,
  spotify_user_id text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  time_span time_span NOT NULL,
  items track [] NOT NULL
);
COMMIT;
BEGIN;
CREATE TABLE artists (
  id SERIAL PRIMARY KEY,
  spotify_user_id text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  time_span time_span NOT NULL,
  items artist [] NOT NULL
);
COMMIT;
BEGIN;
CREATE INDEX ON tracks (spotify_user_id);
COMMIT;
BEGIN;
CREATE INDEX ON artists (spotify_user_id);
COMMIT;