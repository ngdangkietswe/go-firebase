create table if not exists users
(
    id         uuid primary key      default gen_random_uuid(),
    email      varchar(255) not null unique,
    first_name varchar(100),
    last_name  varchar(100),
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

create table if not exists device_token
(
    id         uuid primary key      default gen_random_uuid(),
    user_id    uuid         not null references users (id) on delete cascade,
    token      varchar(512) not null unique,
    platform   varchar(50)  not null, -- e.g., 'iOS', 'Android', 'Web'
    last_seen  timestamp    not null default now(),
    is_active  boolean      not null default true,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

create index idx_device_token_user_id on device_token (user_id);

create table if not exists notification
(
    id      uuid primary key   default gen_random_uuid(),
    user_id uuid      not null references users (id) on delete cascade,
    title   varchar(255),
    body    text,
    data    jsonb,
    sent_at timestamp not null default now(),
    is_read boolean   not null default false
);

create index idx_notification_user_id on notification (user_id);