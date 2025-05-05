-- Crear la base de datos
CREATE DATABASE series_tracker;

-- Usar la base de datos
\c series_tracker;

-- Crear la tabla para las series
CREATE TABLE tv_series (
    series_id SERIAL PRIMARY KEY,                -- Identificador único de la serie
    series_title VARCHAR(255) NOT NULL,          -- Título de la serie
    viewing_status VARCHAR(50) NOT NULL,         -- Estado de la serie (e.g., Watching, Completed)
    episodes_watched INT DEFAULT 0 CHECK (episodes_watched >= 0), -- Número de episodios vistos (no negativo)
    total_episode_count INT DEFAULT 0 CHECK (total_episode_count >= 0), -- Número total de episodios (no negativo)
    user_rating INT DEFAULT 0 CHECK (user_rating >= 0) -- Ranking de la serie (no negativo)
);
