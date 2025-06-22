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
    is_direct BOOL DEFAULT FALSE,
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



--  direct message group get
CREATE OR REPLACE FUNCTION get_or_create_direct_group(
    p_user1_id INT,
    p_user2_id INT
)
RETURNS INT AS $$
DECLARE
    group_id INT;
    tmp INT;
BEGIN
    -- Ensure lowest user ID comes first to maintain uniqueness
    IF p_user1_id > p_user2_id THEN
        tmp := p_user1_id;
        p_user1_id := p_user2_id;
        p_user2_id := tmp;
    END IF;

    -- Try to find existing direct group between these two users
    SELECT g.id INTO group_id
    FROM groups g
    INNER JOIN group_members gm1 ON gm1.group_id = g.id AND gm1.user_id = p_user1_id
    INNER JOIN group_members gm2 ON gm2.group_id = g.id AND gm2.user_id = p_user2_id
    WHERE g.is_direct = TRUE
    LIMIT 1;

    -- If found, return it
    IF group_id IS NOT NULL THEN
        RETURN group_id;
    END IF;

    -- Create a new direct group
    INSERT INTO groups (name, is_direct)
    VALUES (CONCAT('DM:', p_user1_id, '-', p_user2_id), TRUE)
    RETURNING id INTO group_id;

    -- Add both users to the group
    INSERT INTO group_members (user_id, group_id) VALUES (p_user1_id, group_id);
    INSERT INTO group_members (user_id, group_id) VALUES (p_user2_id, group_id);

    RETURN group_id;
END;
$$ LANGUAGE plpgsql;