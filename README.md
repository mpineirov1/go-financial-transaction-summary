# go-financial-transaction-summary
This repository contains code for a system that processes a file from a mounted directory. The file contains a list of debit and credit transactions on an account and the same year. The system processes the file and sends summary information to a user in the form of an email.

Create the table

CREATE TABLE public.transactions
(
    id serial PRIMARY KEY,
    date date NOT NULL,
    transaction numeric(8, 2) NOT NULL,
    created_at timestamp without time zone
);

ALTER TABLE IF EXISTS public.transactions
    OWNER to postgres;