-- master data
INSERT INTO hydroponic_system.plants (
    plant_name,
    scientific_name,
    variety,
    plant_type,
    description,
    pH_min,
    pH_max,
    ppm_min,
    ppm_max,
    light_hours,
    optimal_temperature_min,
    optimal_temperature_max,
    harvest_days,
    germination_days,
    hss_days,
    hst_days
) VALUES
-- 🥬 Lettuce
('Lettuce', 'Lactuca sativa', 'Butterhead', 'Leafy',
 'Most popular hydroponic leafy crop. High demand and fast turnover.',
 5.5, 6.5, 700, 900, 14.0, 18.0, 24.0, 40, 4, 10, 30),

-- 🥬 Pak Choi
('Pak Choi', 'Brassica rapa var. chinensis', 'Green Pak Choi', 'Leafy',
 'Fast-growing leafy vegetable widely used in Asian cuisine.',
 6.0, 7.0, 900, 1200, 12.0, 20.0, 28.0, 35, 3, 7, 28),

-- 🥬 Kale
('Kale', 'Brassica oleracea var. sabellica', 'Curly Kale', 'Leafy',
 'Nutrient-dense green with high market value for healthy food products.',
 5.5, 6.5, 1000, 1400, 14.0, 16.0, 24.0, 60, 5, 14, 46),

-- 🥬 Spinach
('Spinach', 'Spinacia oleracea', 'Green Spinach', 'Leafy',
 'Quick to harvest and suitable for warm climates.',
 6.0, 7.0, 800, 1000, 12.0, 18.0, 26.0, 30, 4, 7, 23),

-- 🥬 Mustard Green
('Mustard Green', 'Brassica juncea', 'Sawi Hijau', 'Leafy',
 'Commonly grown in hydroponics with strong local market demand.',
 6.0, 7.5, 900, 1400, 12.0, 20.0, 30.0, 35, 3, 8, 27),

-- 🥬 Swiss Chard
('Swiss Chard', 'Beta vulgaris subsp. cicla', 'Rainbow Chard', 'Leafy',
 'Colorful and nutritious leafy crop; medium growth speed.',
 6.0, 7.0, 1000, 1400, 14.0, 18.0, 27.0, 55, 5, 12, 43),
 
 -- 🥬 Water Spinach
 ('Water Spinach', 'Ipomoea aquatica', 'Green Kangkung', 'Leafy',
 'Fast-growing leafy vegetable ideal for hydroponic systems; high market demand in Southeast Asia.',
 6.0, 7.5, 900, 1200, 12.0, 22.0, 32.0, 25, 2, 5, 20);


INSERT INTO hydroponic_system.process (process_name, description,type,process_id, created_at, updated_at)
VALUES
('Seed Preparation', 'Selecting and soaking lettuce seeds before germination begins.',NULL,NULL, now(), NULL),
('Germination', 'Sprouting seeds using rockwool or sponge media until roots appear.', NULL,NULL,now(), NULL),
('Seedling Care', 'Growing young seedlings under low EC and moderate light until ready for transplant.', NULL,NULL,now(), NULL),
('Transplanting', 'Moving seedlings into the NFT or tower hydroponic system.',NULL,NULL, now(), NULL),
('Vegetative Growth', 'Main growth phase — monitoring pH, EC, light intensity, and temperature.', NULL,NULL,now(), NULL),
('Nutrient Monitoring', 'Regular checking and adjustment of EC and pH to maintain optimal nutrient levels.',NULL,NULL, now(), NULL),
('Pest & Disease Check', 'Inspecting plants for pests, leaf damage, or fungal infections.', NULL,NULL,now(), NULL),
('System Maintenance', 'Routine maintenance — cleaning filters, checking pumps, and ensuring good water flow.', NULL,NULL,now(), NULL),
('Pre-Harvest Check', 'Inspecting plant size, leaf color, and system readiness before harvest.',NULL,NULL, now(), NULL),
('Harvesting', 'Cutting lettuce heads at maturity and cleaning roots.', NULL,NULL,now(), NULL),
('Post-Harvest Handling', 'Trimming, cleaning, and packaging harvested produce for sale or storage.',NULL,NULL, now(), NULL),
('System Sanitization', 'Cleaning and disinfecting all system parts before the next planting cycle.', NULL,NULL,now(), NULL),
('Planning', 'Module', 1,NULL,now(), NULL),
('Growing', 'Module', 1,NULL,now(), NULL),
('Harvesting', 'Module', 1,NULL,now(), NULL),
('Transaction', 'Module', 1,NULL,now(), NULL),
('Waste Management', 'Module', 1,NULL,now(), NULL),
('Energy Monitoring', 'Module', 1,NULL,now(), NULL),
('Admin', 'Module', 1,NULL,now(), NULL),
('Investation', 'Module', 1,NULL,now(), NULL);

INSERT INTO hydroponic_system.process (process_name, description,type,process_id, created_at, updated_at)
VALUES
('Growth Monitoring', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Growing'),now(), NULL),
('Unhealthy Plant Treatment', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Growing'),now(), NULL),
('Asset Ops. Transaction', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Transaction'),now(), NULL),
('Process', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Admin'),now(), NULL),
('UoM', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Admin'),now(), NULL),
('Remarks', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Admin'),now(), NULL),
('Assets', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Admin'),now(), NULL),
('Plants', 'Sub Module', 2,(SELECT id FROM hydroponic_system.process WHERE process_name = 'Admin'),now(), NULL);

INSERT INTO hydroponic_system.process (process_name,description,type,process_id,created_at)
VALUES
('Inspect & Identify', 
 'Inspect leaves, stems, roots, and environment to determine cause of plant stress before taking action.', 
 3, NULL, NOW()),
('Water Management', 
 'Adjust watering frequency, improve drainage, or increase aeration to prevent root suffocation or rot.', 
 3, NULL, NOW()),
('pH Adjustment', 
 'Measure and correct pH to optimal range (5.8–6.5 for hydroponics) for nutrient absorption and root health.', 
 3, NULL, NOW()),
('Flush Medium / Salt Removal', 
 'Flush the medium or reservoir with clean, pH-balanced water to remove nutrient or salt buildup.', 
 3, NULL, NOW()),
('Refeed with Balanced Nutrients', 
 'Provide a balanced nutrient mix after flushing or deficiency correction to restore proper growth.', 
 3, NULL, NOW()),
('Root Rescue / Replant', 
 'Trim damaged roots, rinse them, and replant in clean medium to recover from root rot or oxygen deprivation.', 
 3, NULL, NOW()),
('Light & Temperature Optimization', 
 'Adjust light distance and ensure stable temperature and humidity to reduce stress and leaf burn.', 
 3, NULL, NOW()),
('Pest Control — Organic Treatment', 
 'Use neem oil or insecticidal soap to remove pests and prevent further infestation.', 
 3, NULL, NOW()),
('Fungal Control — Airflow & Baking Soda Spray', 
 'Improve air circulation and apply mild fungicide or baking soda spray to treat fungal infections.', 
 3, NULL, NOW()),
('Sterilize Tools & Clean Equipment', 
 'Disinfect tools, trays, and reservoirs to prevent disease spread among plants.', 
 3, NULL, NOW()),
('Prune & Remove Dead Tissue', 
 'Cut away diseased or dead leaves to redirect energy and reduce infection risk.', 
 3, NULL, NOW()),
('Foliar Feeding', 
 'Apply dilute foliar nutrient spray to correct deficiencies quickly through leaf absorption.', 
 3, NULL, NOW()),
('Isolation & Quarantine', 
 'Move unhealthy plants away from healthy ones to prevent disease or pest spread.', 
 3, NULL, NOW());



INSERT INTO hydroponic_system.remarks (process_id, remark, description, created_at, updated_at, deleted_at)
VALUES
-- Seeding Stage
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Seedling Care'), 'Rockwool too wet', 'The rockwool is oversaturated, causing low oxygen for seed germination.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Seedling Care'), 'Uneven germination', 'Some seeds sprout slower or fail to germinate due to uneven moisture.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Seedling Care'), 'Algae growth in nursery tray', 'Algae forming on tray surface due to excessive light exposure or high humidity.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Seedling Care'), 'Damping off', 'Seedlings collapse due to fungal infection caused by overwatering or poor airflow.', NOW(), NULL, NULL),

-- Transplanting Stage
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Transplanting'), 'Root damage during transplanting', 'Roots were torn or pinched when moving seedlings to NFT cups.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Transplanting'), 'Netpot not seated properly', 'Netpot not fitted evenly causing poor water contact.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Transplanting'), 'Low survival rate after transplant', 'Plants wilted after transplanting due to stress or shock.', NOW(), NULL, NULL),

-- Vegetative Growth Stage
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'Uneven growth', 'Some plants grow faster than others due to uneven light or nutrient flow.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'PPM too high', 'Nutrient concentration too high causing leaf burn.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'PPM too low', 'Nutrient solution too diluted causing yellow leaves.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'pH drift', 'pH levels increase or decrease rapidly due to unbalanced nutrient uptake.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'Pump failure', 'Pump stopped working, affecting water circulation.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'Nutrient solution temperature too high', 'Warm water reduces oxygen and stresses roots.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Vegetative Growth'), 'Lettuce tip burn', 'Leaf edges burned due to calcium deficiency or high EC.', NOW(), NULL, NULL),

-- Harvesting Stage
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Harvesting'), 'Bolting', 'Lettuce begins flowering prematurely due to high temperature.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Harvesting'), 'Leaf wilting before harvest', 'Leaves droop from high temperature or lack of water.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Harvesting'), 'Mature lettuce head too tight', 'Overgrown lettuce forming hard compact head.', NOW(), NULL, NULL),
((SELECT id FROM hydroponic_system.process WHERE process_name = 'Harvesting'), 'Uneven harvest size', 'Different plant sizes make harvesting inconsistent.', NOW(), NULL, NULL);


INSERT INTO hydroponic_system.uom (name, symbol, description,created_at) VALUES
-- Nutrient & Solution Measurement
('Milliliter', 'mL', 'For measuring liquid nutrients and additives', NOW()),
('Liter', 'L', 'For nutrient solution volume or tank capacity', NOW()),
('Part Per Million', 'ppm', 'For measuring nutrient concentration (EC/TDS)', NOW()),
('Microsiemens per centimeter', 'µS/cm', 'For measuring electrical conductivity (EC)', NOW()),
('pH', 'pH', 'For acidity or alkalinity of nutrient solution', NOW()),

-- Plant Growth & Space Measurement
('Centimeter', 'cm', 'For measuring plant height and spacing', NOW()),
('Meter', 'm', 'For measuring tray or pipe length', NOW()),
('Square Meter', 'm²', 'For area coverage of growing beds', NOW()),

-- Water Flow & Volume Measurement
('Liter per Hour', 'L/h', 'Common unit for pump flow rate', NOW()),
('Liter per Minute', 'L/min', 'Used for faster flow rate measurement', NOW()),

-- Weight Measurement
('Gram', 'g', 'For measuring nutrients or seed weight', NOW()),
('Kilogram', 'kg', 'For harvest weight or bulk material measurement', NOW()),

-- Light & Environment Measurement
('Lux', 'lx', 'For light intensity measurement', NOW()),
('Micromole per square meter per second', 'µmol/m²/s', 'For PAR/PPFD light measurement', NOW()),
('Degree Celsius', '°C', 'For temperature measurement', NOW()),
('Percent', '%', 'For humidity or concentration ratios', NOW());


INSERT INTO hydroponic_system.asset_types (type_name, description, created_at) VALUES
('Infrastructure', 'Structural and growing system components', NOW()),
('Equipment', 'Mechanical and electrical tools for operations', NOW()),
('Consumable', 'Short-term items used regularly', NOW()),
('Monitoring & IT', 'Sensors and control systems for automation', NOW()),
('Operational Support', 'Support tools for daily activities', NOW()),
('Investment / Financial', 'Initial and long-term assets', NOW()),
('Intangible', 'Non-physical assets such as brand or IP', NOW());
