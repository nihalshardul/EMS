USE EMS;


CREATE TABLE IF NOT EXISTS emp (
    id                  INT               PRIMARY KEY,
    first_name          VARCHAR(25)    NOT NULL,
    last_name           VARCHAR(25)    NOT NULL,
    designation           VARCHAR(25)    NOT NULL,
    email               VARCHAR(64)    NOT NULL UNIQUE,
    password            VARBINARY(128)    NOT NULL,
    salary             INT      NOT NULL, 
    lvs                 TINYINT(1)   NOT NULL DEFAULT 0 ,
    is_emp_root         TINYINT(1)   NOT NULL DEFAULT 0 ,
    creation_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    last_update         DATETIME    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by          INT         NOT NULL DEFAULT 0,
    deleted             TINYINT(1)  NOT NULL DEFAULT 0
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE IF NOT EXISTS groups (
    id            INT          PRIMARY KEY,
    group_name           VARCHAR(50)    NOT NULL,            
    creation_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    last_update         DATETIME    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by          INT         NOT NULL DEFAULT 0,
    deleted             TINYINT(1)  NOT NULL DEFAULT 0
    
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS notice (
    id    INT     PRIMARY KEY,
    info    TEXT
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS task(
    id            INT            PRIMARY KEY,
    task_name    VARCHAR(50)    NOT NULL, 
    comment    TEXT    NOT NULL, 
    pending         TINYINT(1)   NOT NULL DEFAULT 1 ,
    complete        TINYINT(1)   NOT NULL DEFAULT 0 ,
    creation_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    last_update         DATETIME    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by          INT         NOT NULL DEFAULT 0,
    deleted             TINYINT(1)  NOT NULL DEFAULT 0
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE IF NOT EXISTS attendance (
    id       INT     AUTO_INCREMENT      PRIMARY KEY,
    emp_id            INT NOT NULL,
    attend_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(emp_id) REFERENCES emp(id) 
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS leaves (
    id       INT     AUTO_INCREMENT      PRIMARY KEY,
    emp_id            INT NOT NULL,
    from_date       DATE  NOT NULL,
    to_date       DATE    NOT NULL,
    FOREIGN KEY(emp_id) REFERENCES emp(id) 
)ENGINE = INNODB CHARACTER SET=utf8;

CREATE TABLE IF NOT EXISTS query (
    id    INT     PRIMARY KEY,
    emp_id    INT NOT NULL,
    comment    TEXT
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS ans_to_query (
    id    INT     PRIMARY KEY,
    q_id    INT NOT NULL,
    comment    TEXT   NOT NULL,
    FOREIGN KEY(q_id) REFERENCES query(id) 
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS emp_group (
    id       INT      PRIMARY KEY,
    gid            INT NOT NULL,
    emp_id             INT  NOT NULL,
    FOREIGN KEY(emp_id) REFERENCES emp(id) ,
    FOREIGN KEY(gid) REFERENCES groups(id) 
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS group_task (
    id       INT      PRIMARY KEY,
    gid            INT NOT NULL,
    tsk_id             INT  NOT NULL,
    FOREIGN KEY(tsk_id) REFERENCES task(id) ,
    FOREIGN KEY(gid) REFERENCES groups(id) 
)ENGINE = INNODB CHARACTER SET=utf8;


CREATE TABLE IF NOT EXISTS login (
    id    INT     PRIMARY KEY,
    emp_id    INT NOT NULL,
    login_date    DATE,
    FOREIGN KEY(emp_id) REFERENCES emp(id) 
)ENGINE = INNODB CHARACTER SET=utf8;