CREATE TYPE time_span AS ENUM('4 weeks', '6 months', 'all time');

CREATE TYPE track AS (
  track_name text NOT NULL,
  artist_name text NOT NULL,
  spotify_track_id text NOT NULL unique,
  track_image_url text NOT NULL,
  track_rank int NOT NULL
);

CREATE TYPE artist AS (
  artist_name text NOT NULL,
  artist_spotify_id text NOT NULL,
  artist_image_url text NOT NULL,
  artist_rank int NOT NULL
);

CREATE TABLE users (
  id int PRIMARY KEY SERIAL,
  spotify_user_id text UNIQUE NOT NULL,
  created_at timestamp NOT NULL,
  is_premium boolean NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE tracks (
  id int PRIMARY KEY SERIAL,
  spotify_user_id text NOT NULL,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  time_span time_span NOT NULL,
  tracks track[] NOT NULL
);

CREATE TABLE artists (
  id int PRIMARY KEY SERIAL,
  spotify_user_id text NOT NULL,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  time_span time_span NOT NULL,
  tracks artist[] NOT NULL
);