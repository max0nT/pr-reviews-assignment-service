CREATE TABLE IF NOT EXISTS teams(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    team_name VARCHAR(50) REFERENCES teams(name),
    is_active BOOLEAN DEFAULT true
);


CREATE TABLE IF NOT EXISTS pull_requests (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    is_merged boolean DEFAULT false,
    created_by_id VARCHAR(50) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    merged_at TIMESTAMP
);


CREATE TABLE IF NOT EXISTS reviewers (
    id SERIAL PRIMARY KEY,
    reviewer_id VARCHAR(50) REFERENCES users(id),
    pr_id VARCHAR(50) REFERENCES pull_requests(id),
    UNIQUE (reviewer_id, pr_id)
);


CREATE INDEX IF NOT EXISTS idx_pull_requests_created_by ON pull_requests(created_by_id);
CREATE INDEX IF NOT EXISTS idx_reviewers_reviewer_id ON reviewers(reviewer_id);
CREATE INDEX IF NOT EXISTS idx_reviewers_pr_id ON reviewers(pr_id);
CREATE INDEX IF NOT EXISTS idx_users_team_id ON users(team_name);
