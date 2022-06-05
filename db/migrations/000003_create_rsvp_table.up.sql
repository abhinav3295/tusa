create table rsvp (
    id uuid primary key,
    eventId uuid,
    userId uuid,
    date timestamp,
    attending integer,
    numberOfGuests integer
);