create table pump_type
(
    id   serial not null
        constraint pump_type_pk
            primary key,
    name text
);

alter table pump_type
    owner to postgres;

create table sensor_type
(
    id   serial not null
        constraint sensors_pk
            primary key,
    name text
);

alter table sensor_type
    owner to postgres;

create table pump_company
(
    id           serial not null
        constraint pump_company_pk
            primary key,
    pump_id      integer,
    company_id   integer,
    created_date bigint
);

alter table pump_company
    owner to postgres;

create table pump_sensor
(
    id           serial not null
        constraint pump_sensor_pk
            primary key,
    type_id      integer,
    pump_id      integer,
    s_value      text,
    created_date bigint
);

alter table pump_sensor
    owner to postgres;

create table pump
(
    id            serial not null
        constraint pump_pk
            primary key,
    type_id       integer,
    serial_number text
        constraint un_serial_number
            unique,
    nick_name     text,
    lat           numeric,
    lng           numeric,
    created_date  bigint
);

alter table pump
    owner to postgres;

create table sensor_data
(
    id            serial not null
        constraint sensor_data_pk
            primary key,
    serial_number text,
    type_id       integer,
    s_value       text,
    update_date   bigint,
    type_text     text,
    constraint sensor_data_pk_2
        unique (type_id, serial_number)
);

alter table sensor_data
    owner to postgres;

create index sensor_data_serial_number_index
    on sensor_data (serial_number);

create table sensor_alarms
(
    id            serial not null
        constraint sensor_alarms_pk
            primary key,
    type_id       integer,
    min_value     integer,
    max_value     integer,
    alert_message text,
    created_date  bigint
);

alter table sensor_alarms
    owner to postgres;

create unique index sensor_alarms_type_id_uindex
    on sensor_alarms (type_id);

create table sensor_alarms_contacts
(
    id            serial not null
        constraint sensor_alarms_contacts_pk
            primary key,
    company_id    integer,
    cell_phone    text,
    email_address text,
    created_date  bigint
);

alter table sensor_alarms_contacts
    owner to postgres;

create index sensor_alarms_contacts_company_id_index
    on sensor_alarms_contacts (company_id);

create unique index sensor_alarms_contacts_company_id_cell_phone_uindex
    on sensor_alarms_contacts (company_id, cell_phone);