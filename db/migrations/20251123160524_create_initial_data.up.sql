INSERT INTO categories (category_name, category_slug, description, icon)
VALUES
    ('Electronics', 'electronics', 'Phones, laptops, cameras, and accessories', 'smartphone'),
    ('Fashion', 'fashion', 'Clothing, shoes, bags, and fashion accessories', 'shirt'),
    ('Automotive', 'automotive', 'Cars, motorcycles, parts, and vehicle accessories', 'car'),
    ('Property', 'property', 'Houses, apartments, rentals, and real estate', 'home'),
    ('Furniture', 'furniture', 'Tables, chairs, beds, and home furnishings', 'bed'),
    ('Hobbies & Collectibles', 'hobbies_collectibles', 'Toys, vinyl records, collectible cards, and memorabilia', 'gamepad'),
    ('Books & Stationery', 'books_stationery', 'Textbooks, novels, notebooks, and office supplies', 'book'),
    ('Sports', 'sports', 'Bikes, fitness equipment, and outdoor gear', 'dumbbell'),
    ('Kids & Baby', 'kids_baby', 'Baby clothes, strollers, educational toys', 'baby'),
    ('Services', 'services', 'Freelance work, tutoring, repairs, and local gigs', 'briefcase'),
    ('Other', 'other', 'Items that don’t fit in other categories', 'package');

-- Electronics
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'General', 'Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'General', 'Model'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'General', 'Release Year'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Display', 'Screen Size (inches)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Display', 'Display Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Processor and Performance', 'Processor Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Processor and Performance', 'Processor Model'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Processor and Performance', 'Cores'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Memory and Storage', 'RAM (GB)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Memory and Storage', 'Internal Storage (GB)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Memory and Storage', 'Expandable Storage'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Battery', 'Battery Capacity (mAh)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Battery', 'Fast Charging'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Camera', 'Rear Camera (MP)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), TRUE, 'Camera', 'Front Camera (MP)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Connectivity', '5G Support'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Connectivity', 'Wi-Fi Standard'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Connectivity', 'Bluetooth Version'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'OS & Software', 'Operating System'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Physical', 'Weight (g)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Physical', 'Color'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Warranty', 'Warranty Period (months)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Battery', 'Battery Health (%)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Durability', 'Water Resistance'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Security', 'Biometric Security'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Connectivity', 'NFC'),
    ((SELECT category_id FROM categories WHERE category_slug = 'electronics'), FALSE, 'Connectivity', 'USB Port Type');

-- Automotive
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'General', 'Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'General', 'Model'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'General', 'Variant'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'General', 'Year'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'General', 'Body Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Registration', 'Registration Status'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Registration', 'Ownership (1st/2nd/etc)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Performance', 'Mileage (km)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Performance', 'Engine (cc)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Performance', 'Fuel Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Performance', 'Transmission Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Performance', 'Horsepower (HP)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), TRUE, 'Performance', 'Torque (Nm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Dimensions', 'Length (mm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Dimensions', 'Width (mm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Dimensions', 'Height (mm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Dimensions', 'Wheelbase (mm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Features', 'Airbags'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Features', 'ABS'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Features', 'Infotainment Screen Size (inches)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Features', 'Sunroof'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Features', 'Keyless Entry'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Condition', 'Paint Condition'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Condition', 'Interior Condition'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Service', 'Last Service (km)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'automotive'), FALSE, 'Warranty', 'Remaining Warranty (months)');

-- Property
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'General', 'Property Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'General', 'Listing Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'General', 'Location'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'Size', 'Land Area (m²)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'Size', 'Building Area (m²)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'Details', 'Bedrooms'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'Details', 'Bathrooms'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'Details', 'Floors'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), TRUE, 'Details', 'Garage'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Legal', 'Certificate Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Legal', 'Has Building Permit'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Facilities', 'Security'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Facilities', 'Swimming Pool'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Facilities', 'Garden'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Utilities', 'Electricity (Watt)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Utilities', 'Water Source'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Orientation', 'Facing Direction'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Age', 'Year Built'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Condition', 'Furnishing'),
    ((SELECT category_id FROM categories WHERE category_slug = 'property'), FALSE, 'Condition', 'Renovation Year');

-- Furniture
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'General', 'Furniture Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'General', 'Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'Material', 'Primary Material'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'Material', 'Upholstery Material'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'Dimensions', 'Height (cm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'Dimensions', 'Width (cm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), TRUE, 'Dimensions', 'Depth (cm)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Features', 'Foldable'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Features', 'Adjustable Height'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Features', 'Storage Compartment'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Condition', 'Used Duration (months)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Condition', 'Scratches/Damage'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Color', 'Primary Color'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Style', 'Design Style'),
    ((SELECT category_id FROM categories WHERE category_slug = 'furniture'), FALSE, 'Assembly', 'Requires Assembly');

-- Sports
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), TRUE, 'General', 'Product Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), TRUE, 'General', 'Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), TRUE, 'General', 'Model'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), TRUE, 'Specifications', 'Weight (kg)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), TRUE, 'Specifications', 'Size'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), TRUE, 'Specifications', 'Color'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), FALSE, 'Specifications', 'Material'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), FALSE, 'Specifications', 'Adjustable'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), FALSE, 'Usage', 'Indoor/Outdoor'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), FALSE, 'Usage', 'Skill Level'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), FALSE, 'Condition', 'Used Duration (months)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'sports'), FALSE, 'Warranty', 'Warranty Period (months)');

-- Kids & Baby
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), TRUE, 'General', 'Product Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), TRUE, 'General', 'Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), TRUE, 'Safety', 'Age Range (months)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), TRUE, 'Safety', 'Certification'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), TRUE, 'Safety', 'Non-Toxic Material'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Details', 'Material'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Details', 'Color'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Dimensions', 'Weight (kg)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Dimensions', 'Foldable'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Usage', 'Battery Required'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Condition', 'Used Duration (months)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'kids_baby'), FALSE, 'Warranty', 'Warranty Period (months)');

-- Fashion
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), TRUE, 'General', 'Product Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), TRUE, 'General', 'Brand'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), TRUE, 'Size & Fit', 'Size'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), TRUE, 'Size & Fit', 'Gender'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), TRUE, 'Material', 'Material'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), FALSE, 'Details', 'Color'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), FALSE, 'Details', 'Pattern'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), FALSE, 'Condition', 'Used Duration (months)'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), FALSE, 'Condition', 'Stains/Damage'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), FALSE, 'Season', 'Season'),
    ((SELECT category_id FROM categories WHERE category_slug = 'fashion'), FALSE, 'Style', 'Style');

-- Books & Stationery
INSERT INTO CategoryProductSpecs (category_id, is_main_spec, spec_type_title, spec_name)
VALUES
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), TRUE, 'Book', 'Title'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), TRUE, 'Book', 'Author'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), TRUE, 'Book', 'ISBN'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), TRUE, 'Book', 'Language'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), TRUE, 'Book', 'Pages'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), TRUE, 'Book', 'Edition'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), FALSE, 'Condition', 'Cover Type'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), FALSE, 'Condition', 'Highlighting/Notes'),
    ((SELECT category_id FROM categories WHERE category_slug = 'books_stationery'), FALSE, 'Condition', 'Used Duration (months)');