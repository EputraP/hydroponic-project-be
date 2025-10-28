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


 INSERT INTO hydroponic_system.process (process_name, description, created_at, updated_at)
VALUES
('Seed Preparation', 'Selecting and soaking lettuce seeds before germination begins.', now(), NULL),
('Germination', 'Sprouting seeds using rockwool or sponge media until roots appear.', now(), NULL),
('Seedling Care', 'Growing young seedlings under low EC and moderate light until ready for transplant.', now(), NULL),
('Transplanting', 'Moving seedlings into the NFT or tower hydroponic system.', now(), NULL),
('Vegetative Growth', 'Main growth phase — monitoring pH, EC, light intensity, and temperature.', now(), NULL),
('Nutrient Monitoring', 'Regular checking and adjustment of EC and pH to maintain optimal nutrient levels.', now(), NULL),
('Pest & Disease Check', 'Inspecting plants for pests, leaf damage, or fungal infections.', now(), NULL),
('System Maintenance', 'Routine maintenance — cleaning filters, checking pumps, and ensuring good water flow.', now(), NULL),
('Pre-Harvest Check', 'Inspecting plant size, leaf color, and system readiness before harvest.', now(), NULL),
('Harvesting', 'Cutting lettuce heads at maturity and cleaning roots.', now(), NULL),
('Post-Harvest Handling', 'Trimming, cleaning, and packaging harvested produce for sale or storage.', now(), NULL),
('System Sanitization', 'Cleaning and disinfecting all system parts before the next planting cycle.', now(), NULL);\


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


INSERT INTO hydroponic_system.uom (name, symbol, description) VALUES
-- Nutrient & Solution Measurement
('Milliliter', 'mL', 'For measuring liquid nutrients and additives'),
('Liter', 'L', 'For nutrient solution volume or tank capacity'),
('Part Per Million', 'ppm', 'For measuring nutrient concentration (EC/TDS)'),
('Microsiemens per centimeter', 'µS/cm', 'For measuring electrical conductivity (EC)'),
('pH', 'pH', 'For acidity or alkalinity of nutrient solution'),

-- Plant Growth & Space Measurement
('Centimeter', 'cm', 'For measuring plant height and spacing'),
('Meter', 'm', 'For measuring tray or pipe length'),
('Square Meter', 'm²', 'For area coverage of growing beds'),

-- Water Flow & Volume Measurement
('Liter per Hour', 'L/h', 'Common unit for pump flow rate'),
('Liter per Minute', 'L/min', 'Used for faster flow rate measurement'),

-- Weight Measurement
('Gram', 'g', 'For measuring nutrients or seed weight'),
('Kilogram', 'kg', 'For harvest weight or bulk material measurement'),

-- Light & Environment Measurement
('Lux', 'lx', 'For light intensity measurement'),
('Micromole per square meter per second', 'µmol/m²/s', 'For PAR/PPFD light measurement'),
('Degree Celsius', '°C', 'For temperature measurement'),
('Percent', '%', 'For humidity or concentration ratios');


INSERT INTO hydroponic_system.asset_types (type_name, description, created_at) VALUES
('Infrastructure', 'Structural and growing system components', NOW()),
('Equipment', 'Mechanical and electrical tools for operations', NOW()),
('Consumable', 'Short-term items used regularly', NOW()),
('Monitoring & IT', 'Sensors and control systems for automation', NOW()),
('Operational Support', 'Support tools for daily activities', NOW()),
('Investment / Financial', 'Initial and long-term assets', NOW()),
('Intangible', 'Non-physical assets such as brand or IP', NOW());
