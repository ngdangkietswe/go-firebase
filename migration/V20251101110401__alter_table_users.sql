alter table users
    add column if not exists status                int     not null default 1, -- 0: Unspecified, 1: Active, 2: Inactive
    add column if not exists last_login_at         timestamptz,
    add column if not exists last_login_ip         varchar(45),
    add column if not exists last_login_user_agent varchar(512),
    add column if not exists failed_login_attempts int     not null default 0,
    add column if not exists deleted               boolean not null default false,
    add column if not exists deleted_at            timestamptz,
    add column if not exists deleted_by            uuid    null references users (id),
    add column if not exists created_by            uuid    null references users (id),
    add column if not exists updated_by            uuid    null references users (id);