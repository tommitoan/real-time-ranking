-- Users table
CREATE TABLE users (
                       id TEXT PRIMARY KEY,
                       username TEXT NOT NULL,
                       email TEXT NOT NULL UNIQUE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Videos table
CREATE TABLE videos (
                        id TEXT PRIMARY KEY,
                        title TEXT NOT NULL,
                        description TEXT,
                        url TEXT NOT NULL,
                        thumbnail TEXT,
                        owner_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

                        views BIGINT DEFAULT 0,
                        likes BIGINT DEFAULT 0,
                        comments BIGINT DEFAULT 0,
                        shares BIGINT DEFAULT 0,
                        watch_time BIGINT DEFAULT 0,

                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
