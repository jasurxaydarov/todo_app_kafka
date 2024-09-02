create database kafka_todo_app;


CREATE TABLE users(
    user_id         serial               PRIMARY KEY,
    user_name       VARCHAR(124)       UNIQUE       NOT NULL,
    password        VARCHAR(69)                     NOT NULL,
    created_at      TIMESTAMP                       DEFAULT NOW()    
);


CREATE TABLE todo(
    todo_id         serial               PRIMARY KEY ,
    title           VARCHAR(255)                    NOT NULL,
    descreption     TEXT,
    is_completed    BOOLEAN                         DEFAULT FALSE,
    created_at      TIMESTAMP                       DEFAULT NOW(),                   
    user_id         serial                            REFERENCES  users(user_id)
);
