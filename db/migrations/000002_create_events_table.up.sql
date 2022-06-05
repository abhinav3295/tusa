create table events (
    id uuid primary key,
    createdBy uuid,
    date timestamp,
    venue varchar(255),
    category varchar(50),
    description varchar(255)
);