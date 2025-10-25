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
 6.0, 7.5, 900, 1200, 12.0, 22.0, 32.0, 25, 2, 5, 20);;
