-- Produk 1: iPhone 15 Pro Max
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'iPhone 15 Pro Max 256GB', 'iphone-15-pro-max-256gb', 'Brand new, titanium finish. Never used, still sealed with full warranty.', 15999000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), '/assets/apple_iphone_15_pro_max_natural_titanium_1_1_2.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), '/assets/apple_iphone_15_pro_max_natural_titanium_2_1_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), 1, 'Apple'),
    ((SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), 2, 'iPhone 15 Pro Max'),
    ((SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), 10, '256'),
    ((SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), 21, 'Titanium Black');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'iPhone 15 Pro Max 256GB' ORDER BY product_id DESC LIMIT 1), -6.2088, 106.8456);

-- Produk 2: MacBook Air M2
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'MacBook Air M2 13-inch', 'macbook-air-m2-13-inch', 'Used for 3 months, like new condition. Comes with charger and sleeve.', 12500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'MacBook Air M2 13-inch' ORDER BY product_id DESC LIMIT 1), '/assets/macbook-air-m2-13-inch.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'MacBook Air M2 13-inch' ORDER BY product_id DESC LIMIT 1), '/assets/macbook-air-m2-13-inch_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'MacBook Air M2 13-inch' ORDER BY product_id DESC LIMIT 1), 1, 'Apple'),
    ((SELECT product_id FROM products WHERE title = 'MacBook Air M2 13-inch' ORDER BY product_id DESC LIMIT 1), 2, 'MacBook Air M2'),
    ((SELECT product_id FROM products WHERE title = 'MacBook Air M2 13-inch' ORDER BY product_id DESC LIMIT 1), 10, '512');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'MacBook Air M2 13-inch' ORDER BY product_id DESC LIMIT 1), -6.9175, 107.6191);

-- Produk 3: Sofa Minimalis
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 5, 'Sofa Minimalis 2 Seater', 'sofa-minimalis-2-seater', 'Sofa kain halus, warna abu-abu. Masih sangat bagus, pindahan rumah.', 1800000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Sofa Minimalis 2 Seater' ORDER BY product_id DESC LIMIT 1), '/assets/sofa-minimalis-2-seater.png', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Sofa Minimalis 2 Seater' ORDER BY product_id DESC LIMIT 1), 28, 'Kain'),
    ((SELECT product_id FROM products WHERE title = 'Sofa Minimalis 2 Seater' ORDER BY product_id DESC LIMIT 1), 29, 'Abu-abu');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Sofa Minimalis 2 Seater' ORDER BY product_id DESC LIMIT 1), -7.7956, 110.3695);

-- Produk 4: Sepatu Nike Air Jordan
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 2, 'Sepatu Nike Air Jordan 1', 'sepatu-nike-air-jordan-1', 'Original, size 42. Dipakai 2x, kondisi 95%.', 2500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Sepatu Nike Air Jordan 1' ORDER BY product_id DESC LIMIT 1), '/assets/sepatu-nike-air-jordan-1.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Sepatu Nike Air Jordan 1' ORDER BY product_id DESC LIMIT 1), '/assets/sepatu-nike-air-jordan-1_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Sepatu Nike Air Jordan 1' ORDER BY product_id DESC LIMIT 1), 31, 'Nike'),
    ((SELECT product_id FROM products WHERE title = 'Sepatu Nike Air Jordan 1' ORDER BY product_id DESC LIMIT 1), 32, '42');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Sepatu Nike Air Jordan 1' ORDER BY product_id DESC LIMIT 1), -6.1754, 106.8272);

-- Produk 5: Kamera Canon EOS R50
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'Kamera Canon EOS R50', 'kamera-canon-eos-r50', 'Mirrorless terbaru dari Canon. Body only, kondisi mulus.', 9500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Kamera Canon EOS R50' ORDER BY product_id DESC LIMIT 1), '/assets/kamera-canon-eos-r50.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Kamera Canon EOS R50' ORDER BY product_id DESC LIMIT 1), '/assets/kamera-canon-eos-r50_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Kamera Canon EOS R50' ORDER BY product_id DESC LIMIT 1), 1, 'Canon'),
    ((SELECT product_id FROM products WHERE title = 'Kamera Canon EOS R50' ORDER BY product_id DESC LIMIT 1), 2, 'EOS R50');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Kamera Canon EOS R50' ORDER BY product_id DESC LIMIT 1), -6.2658, 106.8102);

-- Produk 6: Meja Belajar Kayu Jati
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 5, 'Meja Belajar Kayu Jati', 'meja-belajar-kayu-jati', 'Meja solid wood, ukuran 120x60 cm. Tidak ada cacat.', 2200000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Meja Belajar Kayu Jati' ORDER BY product_id DESC LIMIT 1), '/assets/meja-belajar-kayu-jati.webp', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Meja Belajar Kayu Jati' ORDER BY product_id DESC LIMIT 1), 34, 'Jati'),
    ((SELECT product_id FROM products WHERE title = 'Meja Belajar Kayu Jati' ORDER BY product_id DESC LIMIT 1), 35, '120x60 cm');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Meja Belajar Kayu Jati' ORDER BY product_id DESC LIMIT 1), -7.2575, 112.7521);

-- Produk 7: Tas Gucci Marmont Mini
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 2, 'Tas Gucci Marmont Mini', 'tas-gucci-marmont-mini', 'Authentic Gucci, beli di Paris. Sudah ada dust bag.', 18000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Tas Gucci Marmont Mini' ORDER BY product_id DESC LIMIT 1), '/assets/tas-gucci-marmont-mini.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Tas Gucci Marmont Mini' ORDER BY product_id DESC LIMIT 1), '/assets/tas-gucci-marmont-mini_1.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Tas Gucci Marmont Mini' ORDER BY product_id DESC LIMIT 1), 36, 'Gucci'),
    ((SELECT product_id FROM products WHERE title = 'Tas Gucci Marmont Mini' ORDER BY product_id DESC LIMIT 1), 29, 'Black');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Tas Gucci Marmont Mini' ORDER BY product_id DESC LIMIT 1), -6.2146, 106.8227);

-- Produk 8: Samsung Galaxy S23 Ultra
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'Samsung Galaxy S23 Ultra', 'samsung-galaxy-s23-ultra', '512GB, warna hijau. Masih garansi resmi.', 13500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Samsung Galaxy S23 Ultra' ORDER BY product_id DESC LIMIT 1), '/assets/samsung-galaxy-s23-ultra.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Samsung Galaxy S23 Ultra' ORDER BY product_id DESC LIMIT 1), '/assets/samsung-galaxy-s23-ultra_2.png', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Samsung Galaxy S23 Ultra' ORDER BY product_id DESC LIMIT 1), 1, 'Samsung'),
    ((SELECT product_id FROM products WHERE title = 'Samsung Galaxy S23 Ultra' ORDER BY product_id DESC LIMIT 1), 2, 'Galaxy S23 Ultra'),
    ((SELECT product_id FROM products WHERE title = 'Samsung Galaxy S23 Ultra' ORDER BY product_id DESC LIMIT 1), 10, '512');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Samsung Galaxy S23 Ultra' ORDER BY product_id DESC LIMIT 1), -6.9040, 107.6154);

-- Produk 9: Lemari Pakaian 3 Pintu
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 5, 'Lemari Pakaian 3 Pintu', 'lemari-pakaian-3-pintu', 'Lemari kayu dengan cermin. Ukuran besar, muat banyak baju.', 3500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Lemari Pakaian 3 Pintu' ORDER BY product_id DESC LIMIT 1), '/assets/lemari-pakaian-3-pintu.jpeg', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Lemari Pakaian 3 Pintu' ORDER BY product_id DESC LIMIT 1), 34, 'Kayu Lapis'),
    ((SELECT product_id FROM products WHERE title = 'Lemari Pakaian 3 Pintu' ORDER BY product_id DESC LIMIT 1), 38, '3 Pintu');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Lemari Pakaian 3 Pintu' ORDER BY product_id DESC LIMIT 1), -7.7776, 110.3748);

-- Produk 10: Jam Tangan Rolex Submariner
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 2, 'Jam Tangan Rolex Submariner', 'jam-tangan-rolex-submariner', 'Original, full set box & papers. Harga nego!', 180000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Jam Tangan Rolex Submariner' ORDER BY product_id DESC LIMIT 1), '/assets/jam-tangan-rolex-submariner.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Jam Tangan Rolex Submariner' ORDER BY product_id DESC LIMIT 1), '/assets/jam-tangan-rolex-submariner_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Jam Tangan Rolex Submariner' ORDER BY product_id DESC LIMIT 1), 39, 'Rolex'),
    ((SELECT product_id FROM products WHERE title = 'Jam Tangan Rolex Submariner' ORDER BY product_id DESC LIMIT 1), 29, 'Black');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Jam Tangan Rolex Submariner' ORDER BY product_id DESC LIMIT 1), -6.1884, 106.8312);

-- Produk 11: iPad Pro 12.9-inch M2
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'iPad Pro 12.9-inch M2', 'ipad-pro-12-9-inch-m2', '256GB, Wi-Fi only. Masih segel!', 11000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'iPad Pro 12.9-inch M2' ORDER BY product_id DESC LIMIT 1), '/assets/ipad-pro-12-9-inch-m2.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'iPad Pro 12.9-inch M2' ORDER BY product_id DESC LIMIT 1), '/assets/ipad-pro-12-9-inch-m2_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'iPad Pro 12.9-inch M2' ORDER BY product_id DESC LIMIT 1), 1, 'Apple'),
    ((SELECT product_id FROM products WHERE title = 'iPad Pro 12.9-inch M2' ORDER BY product_id DESC LIMIT 1), 2, 'iPad Pro 12.9'),
    ((SELECT product_id FROM products WHERE title = 'iPad Pro 12.9-inch M2' ORDER BY product_id DESC LIMIT 1), 10, '256');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'iPad Pro 12.9-inch M2' ORDER BY product_id DESC LIMIT 1), -6.2658, 106.8102);

-- Produk 12: Kursi Ergonomis Kantor
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 5, 'Kursi Ergonomis Kantor', 'kursi-ergonomis-kantor', 'Kursi gaming premium, bisa recline. Nyaman dipakai 8 jam.', 2800000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Kursi Ergonomis Kantor' ORDER BY product_id DESC LIMIT 1), '/assets/kursi-ergonomis-kantor.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Kursi Ergonomis Kantor' ORDER BY product_id DESC LIMIT 1), '/assets/kursi-ergonomis-kantor_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Kursi Ergonomis Kantor' ORDER BY product_id DESC LIMIT 1), 28, 'Kulit Sintetis'),
    ((SELECT product_id FROM products WHERE title = 'Kursi Ergonomis Kantor' ORDER BY product_id DESC LIMIT 1), 29, 'Merah Hitam');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Kursi Ergonomis Kantor' ORDER BY product_id DESC LIMIT 1), -6.9175, 107.6191);

-- Produk 13: Dompet Louis Vuitton
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 2, 'Dompet Kulit Pria Louis Vuitton', 'dompet-kulit-pria-louis-vuitton', 'Authentic, kondisi 99%. Sudah ada tas kecil.', 6500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Dompet Kulit Pria Louis Vuitton' ORDER BY product_id DESC LIMIT 1), '/assets/dompet-kulit-pria-louis-vuitton.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Dompet Kulit Pria Louis Vuitton' ORDER BY product_id DESC LIMIT 1), '/assets/dompet-kulit-pria-louis-vuitton_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Dompet Kulit Pria Louis Vuitton' ORDER BY product_id DESC LIMIT 1), 36, 'Louis Vuitton'),
    ((SELECT product_id FROM products WHERE title = 'Dompet Kulit Pria Louis Vuitton' ORDER BY product_id DESC LIMIT 1), 29, 'Brown');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Dompet Kulit Pria Louis Vuitton' ORDER BY product_id DESC LIMIT 1), -6.2088, 106.8456);

-- Produk 14: Laptop Dell XPS 13
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'Laptop Dell XPS 13', 'laptop-dell-xps-13', 'Core i7, 16GB RAM, 512GB SSD. Masih sangat cepat.', 9000000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), '/assets/laptop-dell-xps-13.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), '/assets/laptop-dell-xps-13_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), 1, 'Dell'),
    ((SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), 2, 'XPS 13'),
    ((SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), 8, '16'),
    ((SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), 10, '512');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Laptop Dell XPS 13' ORDER BY product_id DESC LIMIT 1), -7.2575, 112.7521);

-- Produk 15: Rak Buku Minimalis
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 5, 'Rak Buku Minimalis', 'rak-buku-minimalis', 'Rak kayu solid, 5 susun. Cocok untuk kamar atau ruang tamu.', 1500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Rak Buku Minimalis' ORDER BY product_id DESC LIMIT 1), '/assets/rak-buku-minimalis.jpeg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Rak Buku Minimalis' ORDER BY product_id DESC LIMIT 1), '/assets/rak-buku-minimalis_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Rak Buku Minimalis' ORDER BY product_id DESC LIMIT 1), 34, 'Kayu Solid'),
    ((SELECT product_id FROM products WHERE title = 'Rak Buku Minimalis' ORDER BY product_id DESC LIMIT 1), 41, '5 Susun');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Rak Buku Minimalis' ORDER BY product_id DESC LIMIT 1), -7.7956, 110.3695);

-- Produk 16: Kacamata Ray-Ban Original
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 2, 'Kacamata Ray-Ban Original', 'kacamata-ray-ban-original', 'Model Aviator, lensa polarized. Masih ada case-nya.', 1200000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Kacamata Ray-Ban Original' ORDER BY product_id DESC LIMIT 1), '/assets/kacamata-ray-ban-original.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Kacamata Ray-Ban Original' ORDER BY product_id DESC LIMIT 1), '/assets/kacamata-ray-ban-original_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Kacamata Ray-Ban Original' ORDER BY product_id DESC LIMIT 1), 42, 'Ray-Ban'),
    ((SELECT product_id FROM products WHERE title = 'Kacamata Ray-Ban Original' ORDER BY product_id DESC LIMIT 1), 29, 'Gold/Smoke');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Kacamata Ray-Ban Original' ORDER BY product_id DESC LIMIT 1), -6.1754, 106.8272);

-- Produk 17: Headphone Sony WH-1000XM5
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'Headphone Sony WH-1000XM5', 'headphone-sony-wh-1000xm5', 'Noise cancelling terbaik. Masih garansi, kondisi mulus.', 3800000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Headphone Sony WH-1000XM5' ORDER BY product_id DESC LIMIT 1), '/assets/headphone-sony-wh-1000xm5.jpg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Headphone Sony WH-1000XM5' ORDER BY product_id DESC LIMIT 1), '/assets/headphone-sony-wh-1000xm5_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Headphone Sony WH-1000XM5' ORDER BY product_id DESC LIMIT 1), 1, 'Sony'),
    ((SELECT product_id FROM products WHERE title = 'Headphone Sony WH-1000XM5' ORDER BY product_id DESC LIMIT 1), 2, 'WH-1000XM5');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Headphone Sony WH-1000XM5' ORDER BY product_id DESC LIMIT 1), -6.2658, 106.8102);

-- Produk 18: Meja Makan Kayu Jati
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 5, 'Meja Makan Kayu Jati', 'meja-makan-kayu-jati', 'Untuk 6 orang, sudah termasuk kursi. Kualitas premium.', 8500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Meja Makan Kayu Jati' ORDER BY product_id DESC LIMIT 1), '/assets/meja-makan-kayu-jati.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Meja Makan Kayu Jati' ORDER BY product_id DESC LIMIT 1), '/assets/meja-makan-kayu-jati_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Meja Makan Kayu Jati' ORDER BY product_id DESC LIMIT 1), 34, 'Jati'),
    ((SELECT product_id FROM products WHERE title = 'Meja Makan Kayu Jati' ORDER BY product_id DESC LIMIT 1), 44, '6 Orang');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Meja Makan Kayu Jati' ORDER BY product_id DESC LIMIT 1), -6.9040, 107.6154);

-- Produk 19: Sepatu Adidas Ultraboost
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 2, 'Sepatu Adidas Ultraboost', 'sepatu-adidas-ultraboost', 'Original, size 41. Dipakai lari 5x, masih seperti baru.', 1500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Sepatu Adidas Ultraboost' ORDER BY product_id DESC LIMIT 1), '/assets/sepatu-adidas-ultraboost.jpeg', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Sepatu Adidas Ultraboost' ORDER BY product_id DESC LIMIT 1), '/assets/sepatu-adidas-ultraboost_2.webp8', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Sepatu Adidas Ultraboost' ORDER BY product_id DESC LIMIT 1), 31, 'Adidas'),
    ((SELECT product_id FROM products WHERE title = 'Sepatu Adidas Ultraboost' ORDER BY product_id DESC LIMIT 1), 32, '41');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Sepatu Adidas Ultraboost' ORDER BY product_id DESC LIMIT 1), -7.7776, 110.3748);

-- Produk 20: Drone DJI Mini 3 Pro
INSERT INTO products (listing_user_id, category_id, title, product_slug, description, price, condition, status)
VALUES (2, 1, 'Drone DJI Mini 3 Pro', 'drone-dji-mini-3-pro', '4K camera, full set. Sudah terbang 10 jam total.', 7500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Drone DJI Mini 3 Pro' ORDER BY product_id DESC LIMIT 1), '/assets/drone-dji-mini-3-pro.webp', true, 1),
    ((SELECT product_id FROM products WHERE title = 'Drone DJI Mini 3 Pro' ORDER BY product_id DESC LIMIT 1), '/assets/drone-dji-mini-3-pro_2.png', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
VALUES
    ((SELECT product_id FROM products WHERE title = 'Drone DJI Mini 3 Pro' ORDER BY product_id DESC LIMIT 1), 1, 'DJI'),
    ((SELECT product_id FROM products WHERE title = 'Drone DJI Mini 3 Pro' ORDER BY product_id DESC LIMIT 1), 2, 'Mini 3 Pro'),
    ((SELECT product_id FROM products WHERE title = 'Drone DJI Mini 3 Pro' ORDER BY product_id DESC LIMIT 1), 45, '4K');
INSERT INTO location (user_id, product_id, latitude, longitude)
VALUES (2, (SELECT product_id FROM products WHERE title = 'Drone DJI Mini 3 Pro' ORDER BY product_id DESC LIMIT 1), -6.1884, 106.8312);