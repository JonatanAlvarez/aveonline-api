--
-- PostgreSQL database dump
--

-- Dumped from database version 11.12 (Debian 11.12-1.pgdg100+1)
-- Dumped by pg_dump version 13.3 (Debian 13.3-1.pgdg100+1)

-- Started on 2021-10-24 20:03:03 -04

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

--
-- TOC entry 200 (class 1259 OID 36900)
-- Name: factura_medicamentos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.factura_medicamentos (
    factura_id integer NOT NULL,
    medicamento_id integer NOT NULL
);


ALTER TABLE public.factura_medicamentos OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 36907)
-- Name: facturas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.facturas (
    id integer NOT NULL,
    fecha_crear timestamp with time zone,
    pago_total numeric,
    promocion_id integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);


ALTER TABLE public.facturas OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 36905)
-- Name: facturas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.facturas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.facturas_id_seq OWNER TO postgres;

--
-- TOC entry 2940 (class 0 OID 0)
-- Dependencies: 201
-- Name: facturas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.facturas_id_seq OWNED BY public.facturas.id;


--
-- TOC entry 197 (class 1259 OID 36874)
-- Name: medicamentos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.medicamentos (
    id integer NOT NULL,
    nombre character varying(50),
    precio numeric,
    ubicacion character varying(50),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);


ALTER TABLE public.medicamentos OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 36872)
-- Name: medicamentos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.medicamentos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.medicamentos_id_seq OWNER TO postgres;

--
-- TOC entry 2941 (class 0 OID 0)
-- Dependencies: 196
-- Name: medicamentos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.medicamentos_id_seq OWNED BY public.medicamentos.id;


--
-- TOC entry 199 (class 1259 OID 36888)
-- Name: promocions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promocions (
    id integer NOT NULL,
    descripcion character varying(100),
    porcentaje numeric,
    fecha_inicio timestamp with time zone,
    fecha_fin timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);


ALTER TABLE public.promocions OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 36886)
-- Name: promocions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.promocions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.promocions_id_seq OWNER TO postgres;

--
-- TOC entry 2942 (class 0 OID 0)
-- Dependencies: 198
-- Name: promocions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.promocions_id_seq OWNED BY public.promocions.id;


--
-- TOC entry 2800 (class 2604 OID 36910)
-- Name: facturas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.facturas ALTER COLUMN id SET DEFAULT nextval('public.facturas_id_seq'::regclass);


--
-- TOC entry 2794 (class 2604 OID 36877)
-- Name: medicamentos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medicamentos ALTER COLUMN id SET DEFAULT nextval('public.medicamentos_id_seq'::regclass);


--
-- TOC entry 2797 (class 2604 OID 36891)
-- Name: promocions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promocions ALTER COLUMN id SET DEFAULT nextval('public.promocions_id_seq'::regclass);


--
-- TOC entry 2810 (class 2606 OID 36904)
-- Name: factura_medicamentos factura_medicamentos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.factura_medicamentos
    ADD CONSTRAINT factura_medicamentos_pkey PRIMARY KEY (factura_id, medicamento_id);


--
-- TOC entry 2812 (class 2606 OID 36917)
-- Name: facturas facturas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.facturas
    ADD CONSTRAINT facturas_pkey PRIMARY KEY (id);


--
-- TOC entry 2805 (class 2606 OID 36884)
-- Name: medicamentos medicamentos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medicamentos
    ADD CONSTRAINT medicamentos_pkey PRIMARY KEY (id);


--
-- TOC entry 2808 (class 2606 OID 36898)
-- Name: promocions promocions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promocions
    ADD CONSTRAINT promocions_pkey PRIMARY KEY (id);


--
-- TOC entry 2813 (class 1259 OID 36918)
-- Name: idx_facturas_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_facturas_deleted_at ON public.facturas USING btree (deleted_at);


--
-- TOC entry 2803 (class 1259 OID 36885)
-- Name: idx_medicamentos_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_medicamentos_deleted_at ON public.medicamentos USING btree (deleted_at);


--
-- TOC entry 2806 (class 1259 OID 36899)
-- Name: idx_promocions_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_promocions_deleted_at ON public.promocions USING btree (deleted_at);


-- Completed on 2021-10-24 20:03:05 -04

--
-- PostgreSQL database dump complete
--

