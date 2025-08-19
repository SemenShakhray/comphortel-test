-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,         
    login VARCHAR(100) NOT NULL UNIQUE,  
    full_name VARCHAR(200) NOT NULL,     
    gender VARCHAR(10),                 
    age INT,                            
    contacts TEXT,                       
    avatar_url TEXT,                     
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), 
    active BOOLEAN DEFAULT true  
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
