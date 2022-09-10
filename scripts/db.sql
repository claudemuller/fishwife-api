DROP TABLE IF EXISTS reminders;
CREATE TABLE reminders (
   id bigserial,
   user          varchar(255),
   description   text,
   done boolean,
   created_at date not null default CURRENT_TIMESTAMP,
   updated_at date,
   CONSTRAINT pk_reminders PRIMARY KEY(ID)
);
