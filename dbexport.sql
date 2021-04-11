--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 13.2

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

--
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
	NEW.updated_at = now();
	RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: activities; Type: TABLE; Schema: public; Owner: gcunningham
--

CREATE TABLE public.activities (
    id integer NOT NULL,
    name character varying(255),
    description character varying(255),
    metric character varying(255),
    multiplier integer,
    custom boolean,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.activities OWNER TO gcunningham;

--
-- Name: activities_id_seq; Type: SEQUENCE; Schema: public; Owner: gcunningham
--

CREATE SEQUENCE public.activities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.activities_id_seq OWNER TO gcunningham;

--
-- Name: activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gcunningham
--

ALTER SEQUENCE public.activities_id_seq OWNED BY public.activities.id;


--
-- Name: courses; Type: TABLE; Schema: public; Owner: gcunningham
--

CREATE TABLE public.courses (
    id integer NOT NULL,
    name character varying(255),
    institution character varying(255),
    credit_hours integer,
    length integer,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.courses OWNER TO gcunningham;

--
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: gcunningham
--

CREATE SEQUENCE public.courses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.courses_id_seq OWNER TO gcunningham;

--
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gcunningham
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- Name: gopg_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gopg_migrations (
    id integer NOT NULL,
    version bigint,
    created_at timestamp with time zone
);


ALTER TABLE public.gopg_migrations OWNER TO postgres;

--
-- Name: gopg_migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gopg_migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gopg_migrations_id_seq OWNER TO postgres;

--
-- Name: gopg_migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gopg_migrations_id_seq OWNED BY public.gopg_migrations.id;


--
-- Name: module_activities; Type: TABLE; Schema: public; Owner: gcunningham
--

CREATE TABLE public.module_activities (
    id integer NOT NULL,
    input integer,
    notes character varying(255),
    module_id integer NOT NULL,
    activity_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.module_activities OWNER TO gcunningham;

--
-- Name: module_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: gcunningham
--

CREATE SEQUENCE public.module_activities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.module_activities_id_seq OWNER TO gcunningham;

--
-- Name: module_activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gcunningham
--

ALTER SEQUENCE public.module_activities_id_seq OWNED BY public.module_activities.id;


--
-- Name: modules; Type: TABLE; Schema: public; Owner: gcunningham
--

CREATE TABLE public.modules (
    id integer NOT NULL,
    name character varying(255),
    number integer,
    course_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.modules OWNER TO gcunningham;

--
-- Name: modules_id_seq; Type: SEQUENCE; Schema: public; Owner: gcunningham
--

CREATE SEQUENCE public.modules_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.modules_id_seq OWNER TO gcunningham;

--
-- Name: modules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gcunningham
--

ALTER SEQUENCE public.modules_id_seq OWNED BY public.modules.id;


--
-- Name: activities id; Type: DEFAULT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.activities ALTER COLUMN id SET DEFAULT nextval('public.activities_id_seq'::regclass);


--
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- Name: gopg_migrations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gopg_migrations ALTER COLUMN id SET DEFAULT nextval('public.gopg_migrations_id_seq'::regclass);


--
-- Name: module_activities id; Type: DEFAULT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.module_activities ALTER COLUMN id SET DEFAULT nextval('public.module_activities_id_seq'::regclass);


--
-- Name: modules id; Type: DEFAULT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.modules ALTER COLUMN id SET DEFAULT nextval('public.modules_id_seq'::regclass);


--
-- Data for Name: activities; Type: TABLE DATA; Schema: public; Owner: gcunningham
--

COPY public.activities (id, name, description, metric, multiplier, custom, created_at, updated_at) FROM stdin;
1	Reading (understand)	130 wpm; 10 pages per hour	# of pages	6	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
2	Reading (study guide)	65 wpm; 5 pages per hour	# of pages	12	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
3	Writing (research)	6 hours per page (500 words, single-spaced)	# of pages	360	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
4	Writing (reflection)	90 minutes per page (500 words, single-spaced)	# of pages	90	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
5	Learning Objects (matching/multiple choice)	10 minutes per object	# of LOs	10	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
6	Learning Objects (case study)	20 minutes per object	# of LOs	20	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
7	Lecture	Factor 1.25x the actual lecture runtime	# of minutes	1	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
8	Videos	Factor the full length of video	# of minutes	1	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
9	Websites	10-20 minutes		1	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
10	Discussion Boards	250 words/60 minutes for initial post or 2 replies	# of discussion boards	60	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
11	Quizzes	Average 1.5 minutes per question	# of questions	2	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
12	Exams	Average 1.5 minutes per question	# of questions	2	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
13	Self Assessments	Average 1 minute per question	# of questions	1	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
14	Miscellaneous	any additional assignments not listed		1	f	2021-04-10 17:26:54.140713-06	2021-04-10 17:26:54.140713-06
\.


--
-- Data for Name: courses; Type: TABLE DATA; Schema: public; Owner: gcunningham
--

COPY public.courses (id, name, institution, credit_hours, length, created_at, updated_at) FROM stdin;
1	Foundations of Nursing	Colorado Nursing College	3	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
\.


--
-- Data for Name: gopg_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.gopg_migrations (id, version, created_at) FROM stdin;
1	1	2021-04-10 11:35:23.107688-06
2	2	2021-04-10 11:35:23.112611-06
3	3	2021-04-10 11:35:23.116442-06
4	4	2021-04-10 11:35:23.119579-06
5	5	2021-04-10 11:35:23.122399-06
6	4	2021-04-10 11:37:03.420702-06
7	3	2021-04-10 11:37:03.423981-06
8	2	2021-04-10 11:37:03.427902-06
9	1	2021-04-10 11:37:03.429619-06
10	0	2021-04-10 11:37:03.431275-06
11	1	2021-04-10 11:37:30.515279-06
12	2	2021-04-10 11:37:30.520529-06
13	3	2021-04-10 11:37:30.523784-06
14	4	2021-04-10 11:37:30.526957-06
15	5	2021-04-10 11:37:30.52956-06
16	6	2021-04-10 12:14:57.12417-06
17	5	2021-04-10 17:20:16.749841-06
18	4	2021-04-10 17:20:16.752667-06
19	3	2021-04-10 17:20:16.754034-06
20	2	2021-04-10 17:20:16.757863-06
21	1	2021-04-10 17:20:16.760297-06
22	0	2021-04-10 17:20:16.761858-06
23	1	2021-04-10 17:21:07.428152-06
24	2	2021-04-10 17:21:07.432831-06
25	3	2021-04-10 17:21:07.436119-06
26	4	2021-04-10 17:21:07.439305-06
27	5	2021-04-10 17:21:07.442078-06
28	6	2021-04-10 17:21:07.444404-06
29	5	2021-04-10 17:26:09.597597-06
30	4	2021-04-10 17:26:09.600088-06
31	3	2021-04-10 17:26:09.601482-06
32	2	2021-04-10 17:26:09.60439-06
33	1	2021-04-10 17:26:09.606776-06
34	0	2021-04-10 17:26:09.608509-06
35	1	2021-04-10 17:26:54.123954-06
36	2	2021-04-10 17:26:54.128352-06
37	3	2021-04-10 17:26:54.13138-06
38	4	2021-04-10 17:26:54.135125-06
39	5	2021-04-10 17:26:54.138378-06
40	6	2021-04-10 17:26:54.140713-06
41	7	2021-04-11 08:18:53.811535-06
\.


--
-- Data for Name: module_activities; Type: TABLE DATA; Schema: public; Owner: gcunningham
--

COPY public.module_activities (id, input, notes, module_id, activity_id, created_at, updated_at) FROM stdin;
1	107	\N	1	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
2	6	\N	1	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
3	7	\N	1	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
4	95	\N	1	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
5	1	\N	1	10	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
6	450	\N	1	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
7	50	\N	1	13	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
8	53	\N	2	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
9	5	\N	2	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
10	5	\N	2	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
11	71	\N	2	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
12	1	\N	2	10	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
13	100	\N	2	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
14	66	\N	3	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
15	4	\N	3	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
16	1	\N	3	4	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
17	4	\N	3	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
18	2	\N	3	6	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
19	86	\N	3	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
20	1	\N	3	10	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
21	240	\N	3	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
22	50	\N	3	13	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
23	105	\N	4	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
24	7	\N	4	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
25	2	\N	4	4	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
26	3	\N	4	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
27	75	\N	4	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
28	390	\N	4	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
29	50	\N	4	13	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
30	52	\N	5	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
31	5	\N	5	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
32	1	\N	5	4	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
33	5	\N	5	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
34	1	\N	5	6	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
35	62	\N	5	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
36	1	\N	5	10	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
37	300	\N	5	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
38	36	\N	6	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
39	5	\N	6	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
40	5	\N	6	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
41	1	\N	6	6	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
42	40	\N	6	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
43	1	\N	6	10	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
44	90	\N	6	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
45	50	\N	6	13	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
46	88	\N	7	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
47	5	\N	7	2	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
48	4	\N	7	5	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
49	2	\N	7	6	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
50	42	\N	7	8	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
51	240	\N	7	11	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
52	3	\N	8	3	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
53	100	\N	8	13	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
\.


--
-- Data for Name: modules; Type: TABLE DATA; Schema: public; Owner: gcunningham
--

COPY public.modules (id, name, number, course_id, created_at, updated_at) FROM stdin;
1	Module 1	1	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
2	Module 2	2	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
3	Module 3	3	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
4	Module 4	4	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
5	Module 5	5	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
6	Module 6	6	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
7	Module 7	7	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
8	Module 8	8	1	2021-04-11 08:18:53.811535-06	2021-04-11 08:18:53.811535-06
\.


--
-- Name: activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gcunningham
--

SELECT pg_catalog.setval('public.activities_id_seq', 1, false);


--
-- Name: courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gcunningham
--

SELECT pg_catalog.setval('public.courses_id_seq', 2, true);


--
-- Name: gopg_migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gopg_migrations_id_seq', 41, true);


--
-- Name: module_activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gcunningham
--

SELECT pg_catalog.setval('public.module_activities_id_seq', 53, true);


--
-- Name: modules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gcunningham
--

SELECT pg_catalog.setval('public.modules_id_seq', 1, false);


--
-- Name: activities activities_pkey; Type: CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_pkey PRIMARY KEY (id);


--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: module_activities module_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.module_activities
    ADD CONSTRAINT module_activities_pkey PRIMARY KEY (id);


--
-- Name: modules modules_pkey; Type: CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT modules_pkey PRIMARY KEY (id);


--
-- Name: activities update_activity_updated_at; Type: TRIGGER; Schema: public; Owner: gcunningham
--

CREATE TRIGGER update_activity_updated_at BEFORE UPDATE ON public.activities FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: courses update_course_updated_at; Type: TRIGGER; Schema: public; Owner: gcunningham
--

CREATE TRIGGER update_course_updated_at BEFORE UPDATE ON public.courses FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: module_activities update_module_activity_updated_at; Type: TRIGGER; Schema: public; Owner: gcunningham
--

CREATE TRIGGER update_module_activity_updated_at BEFORE UPDATE ON public.module_activities FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: modules update_module_updated_at; Type: TRIGGER; Schema: public; Owner: gcunningham
--

CREATE TRIGGER update_module_updated_at BEFORE UPDATE ON public.modules FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: module_activities fk_activity; Type: FK CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.module_activities
    ADD CONSTRAINT fk_activity FOREIGN KEY (activity_id) REFERENCES public.activities(id);


--
-- Name: modules fk_course; Type: FK CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.modules
    ADD CONSTRAINT fk_course FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: module_activities fk_module; Type: FK CONSTRAINT; Schema: public; Owner: gcunningham
--

ALTER TABLE ONLY public.module_activities
    ADD CONSTRAINT fk_module FOREIGN KEY (module_id) REFERENCES public.modules(id);


--
-- PostgreSQL database dump complete
--

