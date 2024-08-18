-- +goose Up
-- +goose StatementBegin

-- Animes

CREATE TABLE IF NOT EXISTS animes (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

    title VARCHAR NOT NULL,
    description TEXT,
    poster VARCHAR(512),
    episodes INTEGER,
    episodes_released INTEGER NOT NULL,
    
    mal_id BIGINT,
    shiki_id BIGINT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_animes_mal_id ON animes (mal_id);
CREATE UNIQUE INDEX idx_animes_shiki_id ON animes (shiki_id);

-- Related animes

CREATE TABLE IF NOT EXISTS animes_related (
    anime_id BIGINT,
    related_id BIGINT,
    PRIMARY KEY(anime_id, related_id),
    CONSTRAINT fk_animes_related_anime FOREIGN KEY (anime_id) REFERENCES animes (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_animes_related_related FOREIGN KEY (related_id) REFERENCES animes (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- User

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

    username VARCHAR(256) NOT NULL,
    password VARCHAR(256),
    last_online TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- User actions

CREATE TABLE user_actions (
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  
  user_id BIGINT NOT NULL,
  title VARCHAR NOT NULL,
  description VARCHAR NOT NULL,
  
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  
  CONSTRAINT fk_user_actions_user FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Tokens

CREATE TABLE IF NOT EXISTS tokens (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    
    token VARCHAR(256),
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tokens_token ON tokens(token);


-- Collections

CREATE TABLE IF NOT EXISTS collections (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    
    title VARCHAR(512),
    description VARCHAR,
    is_public BOOLEAN NOT NULL DEFAULT TRUE,
    user_id BIGINT,

    updated_at TIMESTAMP  NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), 

    CONSTRAINT fk_users_collections FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE SET NULL
);

-- Animes collections

CREATE TABLE IF NOT EXISTS animes_collections (
    collection_id BIGINT,
    anime_id BIGINT,
    PRIMARY KEY (collection_id, anime_id),
    CONSTRAINT fk_animes_lists_collection FOREIGN KEY (collection_id) REFERENCES collections (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_animes_lists_anime FOREIGN KEY (anime_id) REFERENCES animes (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE animes_collections;

DROP TABLE collections;

DROP INDEX idx_tokens_token;
DROP TABLE tokens;

DROP TABLE user_actions;

DROP TABLE users;

DROP TABLE animes_related;

DROP INDEX idx_animes_mal_id;
DROP INDEX idx_animes_shiki_id;
DROP TABLE animes;

-- +goose StatementEnd
