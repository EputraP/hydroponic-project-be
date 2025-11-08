CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create schema hydroponic_system;

-- Master Tables
CREATE TABLE hydroponic_system.process (
	id uuid DEFAULT public.uuid_generate_v4(),
	process_name varchar NOT NULL,
    description varchar NOT NULL,
    type int,
    process_id uuid,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT proccess_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.remarks (
	id uuid DEFAULT public.uuid_generate_v4(),
    process_id uuid NOT NULL,
	remark varchar NOT NULL,
    description varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT remark_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.plants (
    id uuid DEFAULT public.uuid_generate_v4(),
    plant_name varchar NOT NULL,                        -- e.g. Lettuce, Water Spinach
    scientific_name varchar NULL,                       -- e.g. Lactuca sativa
    variety varchar NULL,                               -- e.g. Butterhead
    plant_type varchar NULL,                            -- e.g. Leafy
    description text NULL,                              -- description of the plant
    pH_min float NULL,                                  -- minimum pH tolerance
    pH_max float NULL,                                  -- maximum pH tolerance
    ppm_min int NULL,                                   -- minimum nutrient PPM
    ppm_max int NULL,                                   -- maximum nutrient PPM
    light_hours float NULL,                             -- hours of light per day
    optimal_temperature_min float NULL,                 -- minimum optimal temperature (°C)
    optimal_temperature_max float NULL,                 -- maximum optimal temperature (°C)
    harvest_days int NULL,                              -- total days from sowing to harvest
    germination_days int NULL,                          -- days to germination
    hss_days int NULL,                                  -- days after sowing before transplant
    hst_days int NULL,                                  -- days after transplant to harvest
    created_at timestamptz DEFAULT now(),               -- record creation time
    updated_at timestamptz DEFAULT now(),               -- record last update time
    deleted_at timestamptz NULL,                        -- soft delete timestamp
    CONSTRAINT plants_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.uom (
	id uuid DEFAULT public.uuid_generate_v4(),
    "name" varchar NOT NULL,
    symbol varchar NOT NULL,
    "description" varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT uom_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.asset_types (
    id UUID DEFAULT public.uuid_generate_v4(),
    type_name VARCHAR NOT NULL,
   "description" varchar NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    CONSTRAINT asset_type_pkey PRIMARY KEY (id)
);


CREATE TABLE hydroponic_system.assets (
	id uuid DEFAULT public.uuid_generate_v4(),
	uom_id uuid,
    plant_id uuid ,
    asset_type_id uuid NOT NULL,
    asset_name varchar NOT NULL,
    price int NOT NULL,
    "value" int NOT NULL,
    cycle int ,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT asset_pkey PRIMARY KEY (id)
);

-- Process Tables
CREATE TABLE hydroponic_system.assets_ops_transaction (
	id uuid DEFAULT public.uuid_generate_v4(),
	asset_id uuid NOT NULL,
    tower_id uuid NOT NULL,
    cycle int NOT NULL,
    "value" real NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT aot_pkey PRIMARY KEY (id)
);
CREATE TABLE hydroponic_system.plant_growth (
	id uuid DEFAULT public.uuid_generate_v4(),
    tower_id uuid NOT NULL,
	process_id uuid NOT NULL,
    cycle int NOT NULL,
    hss int NOT NULL,
    hst int NOT NULL,
    height int NOT NULL,
    ph real NOT NULL,
    ppm real NOT NULL,
    water_level real NOT NULL,
    ovr_plant_condition varchar NOT NULL,
    remarks varchar NOT NULL,
    plant_amount int NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT pg_pkey PRIMARY KEY (id)
);
CREATE TABLE hydroponic_system.unhealthy_plant_treatment (
	id uuid DEFAULT public.uuid_generate_v4(),
    tower_id uuid NOT NULL,
    cycle int NOT NULL,
    tower_hole_no int NOT NULL,
    treatment varchar NOT NULL,
    after_treatment varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT upt_pkey PRIMARY KEY (id)
);


-- Table Relationships
ALTER TABLE ONLY hydroponic_system.assets ADD CONSTRAINT fk_assets_uom FOREIGN KEY (uom_id) REFERENCES hydroponic_system.uom(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets ADD CONSTRAINT fk_assets_plant FOREIGN KEY (plant_id) REFERENCES hydroponic_system.plants(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets ADD CONSTRAINT fk_assets_asset_type FOREIGN KEY (asset_type_id) REFERENCES hydroponic_system.asset_types(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.remarks ADD CONSTRAINT fk_remarks_process FOREIGN KEY (process_id) REFERENCES hydroponic_system.process(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.plant_growth ADD CONSTRAINT fk_pg_asset FOREIGN KEY (tower_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.plant_growth ADD CONSTRAINT fk_pg_process FOREIGN KEY (process_id) REFERENCES hydroponic_system.process(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.unhealthy_plant_treatment ADD CONSTRAINT fk_upt_asset FOREIGN KEY (tower_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets_ops_transaction ADD CONSTRAINT fk_aot_asset FOREIGN KEY (asset_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets_ops_transaction ADD CONSTRAINT fk_aot_tower FOREIGN KEY (tower_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
