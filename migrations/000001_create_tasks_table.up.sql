CREATE TABLE IF NOT EXISTS tasks(
    id UUID NOT NULL,
    assignee VARCHAR(20) NOT NULL,
    title TEXT NOT NULL,
    summary TEXT NOT NULL,
    deadline timestamp DEFAULT CURRENT_DATE,
    status TEXT DEFAULT 'active',
    created_at timestamp DEFAULT CURRENT_DATE,
    updated_at timestamp DEFAULT CURRENT_DATE,
    deleted_at timestamp
);


UPDATE tasks SET new_id = CAST(LPAD(TO_HEX(id), 32, '0') AS UUID);