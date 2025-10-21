alter table if exists notification
    alter column user_id drop not null,
    add column if not exists notification_topic_id uuid references notification_topic (id) on delete set null;

alter table if exists notification_topic
    add column if not exists description varchar(512);