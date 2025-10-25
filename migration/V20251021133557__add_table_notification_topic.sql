create table if not exists notification_topic
(
    id         uuid primary key      default gen_random_uuid(),
    name       varchar(255) not null,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

create table if not exists user_notification_topic
(
    id                    uuid primary key   default gen_random_uuid(),
    user_id               uuid      not null references users (id) on delete cascade,
    notification_topic_id uuid      not null references notification_topic (id) on delete cascade,
    subscribed_at         timestamp not null default now(),
    unique (user_id, notification_topic_id)
);