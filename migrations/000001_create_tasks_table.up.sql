CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    assignee VARCHAR(20) NOT NULL,
    title TEXT NOT NULL,
    summary TEXT NOT NULL,
    deadline DATE NOT NULL DEFAULT CURRENT_DATE,
    status TEXT DEFAULT 'active'
);

