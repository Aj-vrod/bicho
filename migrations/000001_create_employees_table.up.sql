CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country VARCHAR(50),
    job_family VARCHAR(50),
    job_title VARCHAR(50),
    business_unit VARCHAR(50),
    squad VARCHAR(50) NOT NULL,
    platoon VARCHAR(50) NOT NULL,
    battalion VARCHAR(50) NOT NULL,
    start_date VARCHAR(50)
);
