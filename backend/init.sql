CREATE TABLE tv_series (
    series_id SERIAL PRIMARY KEY,
    series_title VARCHAR(255) NOT NULL,
    viewing_status VARCHAR(50) NOT NULL CHECK (viewing_status IN ('Plan to Watch', 'Watching', 'Dropped', 'Completed')),
    episodes_watched INT DEFAULT 0,
    total_episode_count INT DEFAULT 0,
    user_rating INT DEFAULT 0
);