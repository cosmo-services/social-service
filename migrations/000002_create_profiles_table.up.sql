CREATE TABLE IF NOT EXISTS profiles (
    id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    display_name VARCHAR(255),
    avatar_url TEXT,
    is_active BOOLEAN,
    is_deleted BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT profiles_user_id_unique UNIQUE (user_id),
    CONSTRAINT profiles_username_unique UNIQUE (username),
    CONSTRAINT profiles_email_unique UNIQUE (email)
);

CREATE INDEX idx_profiles_user_id ON profiles(user_id) WHERE is_deleted = false;
CREATE INDEX idx_profiles_username ON profiles(username) WHERE is_deleted = false;
CREATE INDEX idx_profiles_email ON profiles(email) WHERE is_deleted = false;
CREATE INDEX idx_profiles_created_at ON profiles(created_at);