-- Create Users Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Create Groups Table
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Mapping Table: Users in Groups
CREATE TABLE IF NOT EXISTS group_members (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id INT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    UNIQUE(user_id, group_id)
);


--  Stored Procedures 
--  To create user
CREATE OR REPLACE FUNCTION create_user(p_name TEXT)
RETURNS TABLE (id INT, name TEXT) AS $$
BEGIN
    RETURN QUERY
    INSERT INTO users(name)
    VALUES (p_name)
    RETURNING users.id, users.name;
END;
$$ LANGUAGE plpgsql;


--  Get User by Id
CREATE OR REPLACE FUNCTION get_user_by_id(p_id INT)
RETURNS TABLE (id INT, name TEXT) AS $$
BEGIN
    RETURN QUERY
    SELECT users.id, users.name
    FROM users
    WHERE users.id = p_id;
END;
$$ LANGUAGE plpgsql;


--  Create Group 

CREATE OR REPLACE FUNCTION create_group(p_name TEXT)
RETURNS TABLE (id INT, name TEXT) AS $$
BEGIN
    RETURN QUERY
    INSERT INTO groups(name)
    VALUES (p_name)
    RETURNING groups.id, groups.name;
END;
$$ LANGUAGE plpgsql;


--  add user to Group 

CREATE OR REPLACE FUNCTION add_user_to_group(p_user_id INT, p_group_id INT)
RETURNS VOID AS $$
BEGIN
    INSERT INTO group_members(user_id, group_id)
    VALUES (p_user_id, p_group_id)
    ON CONFLICT (user_id, group_id) DO NOTHING;
END;
$$ LANGUAGE plpgsql;
