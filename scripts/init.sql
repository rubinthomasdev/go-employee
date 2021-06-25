-- Table: public.employee

-- DROP TABLE public.employee;

CREATE TABLE public.employee
(
    lastname character varying COLLATE pg_catalog."default" NOT NULL,
    id integer NOT NULL,
    firstname character varying COLLATE pg_catalog."default" NOT NULL,
    basesalary double precision NOT NULL,
    bonus double precision NOT NULL,
    CONSTRAINT employee_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.employee
    OWNER to postgres;


INSERT INTO public.employee(
	lastname, id, firstname, basesalary, bonus)
	VALUES ('user', 1, 'test1', 12.12, 13.13);
INSERT INTO public.employee(
	lastname, id, firstname, basesalary, bonus)
	VALUES ('user', 2, 'test2', 12.56, 13.18);