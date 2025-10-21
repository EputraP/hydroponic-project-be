

create schema hydroponic_system;

-- Master Tables
CREATE TABLE hydroponic_system.process (
	id uuid DEFAULT public.uuid_generate_v4(),
	process_name varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT proccess_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.remarks (
	id uuid DEFAULT public.uuid_generate_v4(),
	remark varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT remark_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.plants (
	id uuid DEFAULT public.uuid_generate_v4(),
	plant varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT plant_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.uom (
	id uuid DEFAULT public.uuid_generate_v4(),
	asset_id uuid NOT NULL,
    uom varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT uom_pkey PRIMARY KEY (id)
);

CREATE TABLE hydroponic_system.assets (
	id uuid DEFAULT public.uuid_generate_v4(),
	uom_id uuid NOT NULL,
    plant_id uuid NOT NULL,
    asset_name varchar NOT NULL,
    price int NOT NULL,
    asset_type varchar NOT NULL,
    "value" int NOT NULL,
    cycle int NOT NULL,
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
    treatment int NOT NULL,
    cond_after_tret varchar NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT upt_pkey PRIMARY KEY (id)
);


-- Table Relationships
ALTER TABLE ONLY hydroponic_system.assets ADD CONSTRAINT fk_assets_uom FOREIGN KEY (uom_id) REFERENCES hydroponic_system.uom(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets ADD CONSTRAINT fk_assets_plant FOREIGN KEY (plant_id) REFERENCES hydroponic_system.plants(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.plant_growth ADD CONSTRAINT fk_pg_asset FOREIGN KEY (tower_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.plant_growth ADD CONSTRAINT fk_pg_process FOREIGN KEY (process_id) REFERENCES hydroponic_system.process(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.unhealthy_plant_treatment ADD CONSTRAINT fk_upt_asset FOREIGN KEY (tower_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets_ops_transaction ADD CONSTRAINT fk_aot_asset FOREIGN KEY (asset_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY hydroponic_system.assets_ops_transaction ADD CONSTRAINT fk_aot_tower FOREIGN KEY (tower_id) REFERENCES hydroponic_system.assets(id) ON UPDATE CASCADE ON DELETE SET NULL;