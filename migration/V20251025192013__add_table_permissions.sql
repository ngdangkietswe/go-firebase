create table if not exists roles
(
    id          uuid primary key      default gen_random_uuid(),
    name        varchar(255) not null unique,
    description text,
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now(),
    unique (name)
);

create table if not exists permissions
(
    id          uuid primary key      default gen_random_uuid(),
    action      varchar(255) not null, -- e.g., 'read', 'write', 'delete'
    resource    varchar(255) not null, -- e.g., 'user', 'notification_topic'
    description text,
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now(),
    unique (action, resource)
);

create table if not exists role_permissions
(
    id            uuid primary key default gen_random_uuid(),
    role_id       uuid not null references roles (id) on delete cascade,
    permission_id uuid not null references permissions (id) on delete cascade,
    unique (role_id, permission_id)
);

create table if not exists user_roles
(
    id      uuid primary key default gen_random_uuid(),
    user_id uuid not null references users (id) on delete cascade,
    role_id uuid not null references roles (id) on delete cascade,
    unique (user_id, role_id)
);

create table if not exists user_permissions
(
    id            uuid primary key default gen_random_uuid(),
    user_id       uuid not null references users (id) on delete cascade,
    permission_id uuid not null references permissions (id) on delete cascade,
    unique (user_id, permission_id)
);

insert into roles (name, description)
values ('admin', 'Administrator with full access'),
       ('user', 'Regular user with limited access')
on conflict (name) do nothing;

insert into permissions (action, resource, description)
values ('create', 'user', 'user:create'),
       ('read', 'user', 'user:read'),
       ('update', 'user', 'user:update'),
       ('delete', 'user', 'user:delete'),

       ('create', 'notification_topic', 'notification_topic:create'),
       ('read', 'notification_topic', 'notification_topic:read'),
       ('update', 'notification_topic', 'notification_topic:update'),
       ('delete', 'notification_topic', 'notification_topic:delete')
on conflict (action, resource) do nothing;

insert into role_permissions (role_id, permission_id)
values ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'create' and resource = 'user')),
       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'read' and resource = 'user')),
       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'update' and resource = 'user')),
       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'delete' and resource = 'user')),

       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'create' and resource = 'notification_topic')),
       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'read' and resource = 'notification_topic')),
       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'update' and resource = 'notification_topic')),
       ((select id from roles where name = 'admin'),
        (select id from permissions where action = 'delete' and resource = 'notification_topic')),

       ((select id from roles where name = 'user'),
        (select id from permissions where action = 'read' and resource = 'notification_topic'))
on conflict (role_id, permission_id) do nothing;

insert into user_roles (user_id, role_id)
values ((select id from users where email = 'kietnguyen17052001@gmail.com'),
        (select id from roles where name = 'admin'))
on conflict (user_id, role_id) do nothing;