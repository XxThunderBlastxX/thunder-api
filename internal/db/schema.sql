CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT NOT NULL,
    stacks TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_project_name ON projects(name);

CREATE INDEX idx_project_description ON projects(description);

CREATE INDEX idx_project_link ON projects(link);

CREATE INDEX idx_project_skills ON projects(stacks);