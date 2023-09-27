-- Create a new database
CREATE DATABASE blogpostdb;

-- Connect to the database
\c blogpostdb;

-- Create a table to store blog posts
CREATE TABLE IF NOT EXISTS blog_post (
                                         id SERIAL PRIMARY KEY,
                                         title TEXT NOT NULL,
                                         content TEXT NOT NULL,
                                         created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );

-- Create an index on the 'created_at' column for faster retrieval
CREATE INDEX ON blog_post (created_at);