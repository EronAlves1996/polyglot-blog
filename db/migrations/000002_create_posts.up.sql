CREATE TABLE blog.posts (
  id serial PRIMARY KEY,
  title varchar(255) NOT NULL,
  content text NOT NULL,
  -- The opening_text gonna be used to repost
  -- blog link on other social media platforms
  opening_text varchar(240) NOT NULL,
  covering_image varchar(255) NOT NULL
);
