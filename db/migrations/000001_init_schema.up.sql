CREATE TYPE time_span AS ENUM('4 weeks', '6 months', 'all time');
CREATE TYPE track AS (
  track_name text,
  artist_name text,
  spotify_track_id text,
  track_image_url text,
  track_rank int
);
CREATE TYPE artist AS (
  artist_name text,
  artist_spotify_id text,
  artist_image_url text,
  artist_rank int
);
CREATE TABLE users (
  id int PRIMARY KEY,
  spotify_user_id text UNIQUE NOT NULL,
  created_at timestamp NOT NULL,
  is_premium boolean NOT NULL,
  updated_at timestamp NOT NULL
);
CREATE TABLE tracks (
  id int PRIMARY KEY,
  spotify_user_id text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  time_span time_span NOT NULL,
  tracks track [] NOT NULL
);
CREATE TABLE artists (
  id int PRIMARY KEY,
  spotify_user_id text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  time_span time_span NOT NULL,
  tracks artist [] NOT NULL
);