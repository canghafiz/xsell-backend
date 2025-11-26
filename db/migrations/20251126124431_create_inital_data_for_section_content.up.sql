INSERT INTO content_sections (section_key, title, subtitle, section_type, config, is_active, sort_order)
VALUES
-- Newly listed items
('newly_listed', 'Newly Listed', 'Just posted, not sold yet!', 'predefined', '{"strategy": "newly_listed", "limit": 12}'::JSONB, true, 1),

-- Like New condition
('like_new_items', 'Like New Condition', 'Minimal use, looks brand new!', 'dynamic', '{"filters": {"condition": "Like New", "status": "Available"}, "limit": 10}'::JSONB, true, 2),

-- Good condition
('good_condition', 'Good Condition', 'Fully functional, great value', 'dynamic', '{"filters": {"condition": "Good", "status": "Available"}, "limit": 10}'::JSONB, true, 3),

-- Trending (most viewed)
('trending', 'Trending Now', 'Most viewed this week', 'predefined', '{"strategy": "trending", "limit": 10}'::JSONB, true, 4),

-- Electronics
('used_electronics', 'Used Electronics', 'Phones, laptops, cameras & accessories', 'dynamic', '{"filters": {"category_slug": "electronics", "status": "Available"}, "limit": 10}'::JSONB, true, 5),

-- Furniture & Home
('home_items', 'Home & Furniture', 'Tables, chairs, decor, and more', 'dynamic', '{"filters": {"category_slug": "furniture", "status": "Available"}, "limit": 8}'::JSONB, true, 6),

-- Fashion
('fashion_items', 'Fashion & Accessories', 'Clothing, shoes, bags, and watches', 'dynamic', '{"filters": {"category_slug": "fashion", "status": "Available"}, "limit": 8}'::JSONB, true, 7),

-- Trusted sellers (optional)
('trusted_sellers', 'From Trusted Sellers', 'Highly rated sellers', 'predefined', '{"strategy": "trusted_sellers", "limit": 8}'::JSONB, true, 8),

-- Editor's picks (manual curation)
('editor_pick', 'Editor''s Picks', 'Handpicked just for you', 'fixed', '{"product_ids": []}'::JSONB, true, 9);

-- ==========================================================
-- SEED DATA: page_layouts
-- ==========================================================

-- Home Page
INSERT INTO page_layouts (page_key, section_id, sort_order, is_active)
VALUES
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'newly_listed'), 1, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'trending'), 2, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'like_new_items'), 3, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'good_condition'), 4, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'used_electronics'), 5, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'home_items'), 6, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'fashion_items'), 7, true),
    ('home', (SELECT section_id FROM content_sections WHERE section_key = 'editor_pick'), 8, true);

-- Product Detail Page
INSERT INTO page_layouts (page_key, section_id, sort_order, is_active)
VALUES
    ('product_detail', (SELECT section_id FROM content_sections WHERE section_key = 'trending'), 1, true),
    ('product_detail', (SELECT section_id FROM content_sections WHERE section_key = 'like_new_items'), 2, true);

-- Category: Electronics
INSERT INTO page_layouts (page_key, section_id, sort_order, is_active)
VALUES
    ('category_electronics', (SELECT section_id FROM content_sections WHERE section_key = 'used_electronics'), 1, true),
    ('category_electronics', (SELECT section_id FROM content_sections WHERE section_key = 'like_new_items'), 2, true),
    ('category_electronics', (SELECT section_id FROM content_sections WHERE section_key = 'good_condition'), 3, true);

-- Category: Furniture
INSERT INTO page_layouts (page_key, section_id, sort_order, is_active)
VALUES
    ('category_furniture', (SELECT section_id FROM content_sections WHERE section_key = 'home_items'), 1, true),
    ('category_furniture', (SELECT section_id FROM content_sections WHERE section_key = 'good_condition'), 2, true);

-- Category: Fashion
INSERT INTO page_layouts (page_key, section_id, sort_order, is_active)
VALUES
    ('category_fashion', (SELECT section_id FROM content_sections WHERE section_key = 'fashion_items'), 1, true),
    ('category_fashion', (SELECT section_id FROM content_sections WHERE section_key = 'like_new_items'), 2, true);