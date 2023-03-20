CREATE TABLE IF NOT EXISTS users(
    id varchar(250) PRIMARY KEY NOT NULL,
    name varchar(250) NOT NULL,
    project_v2 varchar(250)
);

CREATE TABLE IF NOT EXISTS repositories(
    id varchar(250) PRIMARY KEY NOT NULL,
    owner varchar(250) NOT NULL,
    name varchar(250) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	FOREIGN KEY (owner) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS issues(
    id varchar(250) PRIMARY KEY NOT NULL,
    url varchar(250) NOT NULL,
    title varchar(250) NOT NULL,
    closed INTEGER NOT NULL DEFAULT 0,
    number INTEGER NOT NULL,
    repository varchar(250) NOT NULL,
    CHECK (closed IN (0, 1)),
	FOREIGN KEY (repository) REFERENCES repositories(id)
);

CREATE TABLE IF NOT EXISTS projects(
    id varchar(250) PRIMARY KEY NOT NULL,
    title varchar(250) NOT NULL,
    url varchar(250) NOT NULL,
    owner varchar(250) NOT NULL,
    FOREIGN KEY (owner) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS pullrequests(
    id varchar(250) PRIMARY KEY NOT NULL,
    base_ref_name varchar(250) NOT NULL,
    closed INTEGER NOT NULL DEFAULT 0,
    head_ref_name varchar(250) NOT NULL,
    url varchar(250) NOT NULL,
    number INTEGER NOT NULL,
    repository varchar(250) NOT NULL,
    CHECK (closed IN (0, 1)),
	FOREIGN KEY (repository) REFERENCES repositories(id)
);

CREATE TABLE IF NOT EXISTS projectcards(
    id varchar(250) PRIMARY KEY NOT NULL,
    project varchar(250) NOT NULL,
    issue varchar(250),
    pullrequest varchar(250),
    FOREIGN KEY (project) REFERENCES projects(id),
	FOREIGN KEY (issue) REFERENCES issues(id),
	FOREIGN KEY (pullrequest) REFERENCES pullrequests(id),
	CHECK (issue IS NOT NULL OR pullrequest IS NOT NULL)
);