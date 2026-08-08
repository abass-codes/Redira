-- Redira Initial Database Schema
-- Migration: 001

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS users (

    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    email VARCHAR(255) UNIQUE NOT NULL,

    password_hash TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);



CREATE TABLE IF NOT EXISTS links (

    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    user_id UUID REFERENCES users(id)
        ON DELETE CASCADE,

    short_code VARCHAR(20) UNIQUE NOT NULL,

    original_url TEXT NOT NULL,

    active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);



CREATE TABLE IF NOT EXISTS analytics_events (

    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    link_id UUID REFERENCES links(id)
        ON DELETE CASCADE,

    ip_address TEXT,

    user_agent TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);