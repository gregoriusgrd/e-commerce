INSERT INTO categories (name, parent_id, created_at, updated_at) VALUES

-- ============================================
-- LEVEL 1: ROOT CATEGORIES (10 kategori)
-- ============================================
('Ruang Tamu',      NULL, NOW(), NOW()),  -- id: 1
('Ruang Makan',     NULL, NOW(), NOW()),  -- id: 2
('Kamar Tidur',     NULL, NOW(), NOW()),  -- id: 3
('Kamar Mandi',     NULL, NOW(), NOW()),  -- id: 4
('Ruang Kerja',     NULL, NOW(), NOW()),  -- id: 5
('Dapur',           NULL, NOW(), NOW()),  -- id: 6
('Outdoor & Taman', NULL, NOW(), NOW()),  -- id: 7
('Anak-Anak',       NULL, NOW(), NOW()),  -- id: 8
('Penyimpanan',     NULL, NOW(), NOW()),  -- id: 9
('Dekorasi',        NULL, NOW(), NOW()),  -- id: 10

-- ============================================
-- LEVEL 2: SUB KATEGORI (40 kategori)
-- ============================================

-- Ruang Tamu (parent: 1)
('Sofa & Kursi',        1, NOW(), NOW()),  -- id: 11
('Meja Tamu',           1, NOW(), NOW()),  -- id: 12
('Lemari & Bufet',      1, NOW(), NOW()),  -- id: 13
('Rak TV & Media',      1, NOW(), NOW()),  -- id: 14

-- Ruang Makan (parent: 2)
('Set Meja Makan',      2, NOW(), NOW()),  -- id: 15
('Kursi Makan',         2, NOW(), NOW()),  -- id: 16
('Meja Bar & Barstool', 2, NOW(), NOW()),  -- id: 17
('Lemari Dapur',        2, NOW(), NOW()),  -- id: 18

-- Kamar Tidur (parent: 3)
('Tempat Tidur',        3, NOW(), NOW()),  -- id: 19
('Lemari Pakaian',      3, NOW(), NOW()),  -- id: 20
('Nakas & Meja Rias',   3, NOW(), NOW()),  -- id: 21
('Meja Belajar Kamar',  3, NOW(), NOW()),  -- id: 22

-- Kamar Mandi (parent: 4)
('Kabinet Cermin',      4, NOW(), NOW()),  -- id: 23
('Rak Kamar Mandi',     4, NOW(), NOW()),  -- id: 24
('Bangku Kamar Mandi',  4, NOW(), NOW()),  -- id: 25

-- Ruang Kerja (parent: 5)
('Meja Kerja',          5, NOW(), NOW()),  -- id: 26
('Kursi Kantor',        5, NOW(), NOW()),  -- id: 27
('Lemari Arsip',        5, NOW(), NOW()),  -- id: 28
('Rak Buku',            5, NOW(), NOW()),  -- id: 29

-- Dapur (parent: 6)
('Meja Dapur',          6, NOW(), NOW()),  -- id: 30
('Kabinet Dapur',       6, NOW(), NOW()),  -- id: 31
('Rak Dapur',           6, NOW(), NOW()),  -- id: 32

-- Outdoor & Taman (parent: 7)
('Kursi & Sofa Taman',  7, NOW(), NOW()),  -- id: 33
('Meja Taman',          7, NOW(), NOW()),  -- id: 34
('Ayunan & Hammock',    7, NOW(), NOW()),  -- id: 35
('Pergola & Gazebo',    7, NOW(), NOW()),  -- id: 36

-- Anak-Anak (parent: 8)
('Tempat Tidur Anak',   8, NOW(), NOW()),  -- id: 37
('Meja Belajar Anak',   8, NOW(), NOW()),  -- id: 38
('Lemari Anak',         8, NOW(), NOW()),  -- id: 39
('Rak Mainan',          8, NOW(), NOW()),  -- id: 40

-- Penyimpanan (parent: 9)
('Lemari Serbaguna',    9, NOW(), NOW()),  -- id: 41
('Rak Serbaguna',       9, NOW(), NOW()),  -- id: 42
('Laci & Komodo',       9, NOW(), NOW()),  -- id: 43

-- Dekorasi (parent: 10)
('Cermin',              10, NOW(), NOW()),  -- id: 44
('Jam Dinding',         10, NOW(), NOW()),  -- id: 45
('Vas & Pot',           10, NOW(), NOW()),  -- id: 46
('Karpet & Permadani',  10, NOW(), NOW()),  -- id: 47
('Lampu & Aksesoris',   10, NOW(), NOW()),  -- id: 48
('Bantal & Selimut',    10, NOW(), NOW()),  -- id: 49
('Figura & Hiasan',     10, NOW(), NOW()),  -- id: 50

-- ============================================
-- LEVEL 3: SUB-SUB KATEGORI (50 kategori)
-- ============================================

-- Sofa & Kursi (parent: 11)
('Sofa L-Shape',         11, NOW(), NOW()),  -- id: 51
('Sofa 2 Seater',        11, NOW(), NOW()),  -- id: 52
('Sofa 3 Seater',        11, NOW(), NOW()),  -- id: 53
('Sofa Bed',             11, NOW(), NOW()),  -- id: 54
('Kursi Santai',         11, NOW(), NOW()),  -- id: 55
('Kursi Rotan',          11, NOW(), NOW()),  -- id: 56

-- Meja Tamu (parent: 12)
('Coffee Table',         12, NOW(), NOW()),  -- id: 57
('Side Table',           12, NOW(), NOW()),  -- id: 58
('Ottoman Table',        12, NOW(), NOW()),  -- id: 59

-- Rak TV & Media (parent: 14)
('Rak TV Minimalis',     14, NOW(), NOW()),  -- id: 60
('Rak TV Gantung',       14, NOW(), NOW()),  -- id: 61
('Meja TV Kaki Kayu',    14, NOW(), NOW()),  -- id: 62

-- Set Meja Makan (parent: 15)
('Meja Makan 4 Kursi',   15, NOW(), NOW()),  -- id: 63
('Meja Makan 6 Kursi',   15, NOW(), NOW()),  -- id: 64
('Meja Makan 8 Kursi',   15, NOW(), NOW()),  -- id: 65
('Meja Makan Lipat',     15, NOW(), NOW()),  -- id: 66

-- Kursi Makan (parent: 16)
('Kursi Kayu Solid',     16, NOW(), NOW()),  -- id: 67
('Kursi Besi Industri',  16, NOW(), NOW()),  -- id: 68
('Kursi Rotan Makan',    16, NOW(), NOW()),  -- id: 69
('Kursi Scandinavian',   16, NOW(), NOW()),  -- id: 70

-- Tempat Tidur (parent: 19)
('Tempat Tidur Single',  19, NOW(), NOW()),  -- id: 71
('Tempat Tidur Double',  19, NOW(), NOW()),  -- id: 72
('Tempat Tidur Queen',   19, NOW(), NOW()),  -- id: 73
('Tempat Tidur King',    19, NOW(), NOW()),  -- id: 74
('Divan Minimalis',      19, NOW(), NOW()),  -- id: 75

-- Lemari Pakaian (parent: 20)
('Lemari 2 Pintu',       20, NOW(), NOW()),  -- id: 76
('Lemari 3 Pintu',       20, NOW(), NOW()),  -- id: 77
('Lemari Geser',         20, NOW(), NOW()),  -- id: 78
('Lemari Walk-in',       20, NOW(), NOW()),  -- id: 79

-- Meja Kerja (parent: 26)
('Meja Kerja Minimalis', 26, NOW(), NOW()),  -- id: 80
('Meja Kerja L-Shape',   26, NOW(), NOW()),  -- id: 81
('Meja Kerja Standing',  26, NOW(), NOW()),  -- id: 82
('Meja Kerja Gaming',    26, NOW(), NOW()),  -- id: 83

-- Kursi Kantor (parent: 27)
('Kursi Direktur',       27, NOW(), NOW()),  -- id: 84
('Kursi Mesh',           27, NOW(), NOW()),  -- id: 85
('Kursi Gaming',         27, NOW(), NOW()),  -- id: 86
('Kursi Ergonomis',      27, NOW(), NOW()),  -- id: 87

-- Kursi & Sofa Taman (parent: 33)
('Kursi Plastik Outdoor',33, NOW(), NOW()),  -- id: 88
('Sofa Rotan Outdoor',   33, NOW(), NOW()),  -- id: 89
('Bean Bag Outdoor',     33, NOW(), NOW()),  -- id: 90

-- Meja Taman (parent: 34)
('Meja Lipat Outdoor',   34, NOW(), NOW()),  -- id: 91
('Meja Besi Taman',      34, NOW(), NOW()),  -- id: 92

-- Tempat Tidur Anak (parent: 37)
('Tempat Tidur Tingkat', 37, NOW(), NOW()),  -- id: 93
('Tempat Tidur Berumah', 37, NOW(), NOW()),  -- id: 94
('Tempat Tidur Bayi',    37, NOW(), NOW()),  -- id: 95

-- Meja Belajar Anak (parent: 38)
('Meja Belajar Lipat',   38, NOW(), NOW()),  -- id: 96
('Set Meja Kursi Anak',  38, NOW(), NOW()),  -- id: 97

-- Rak Mainan (parent: 40)
('Rak Mainan Kayu',      40, NOW(), NOW()),  -- id: 98
('Rak Buku Anak',        40, NOW(), NOW()),  -- id: 99

-- Lemari Serbaguna (parent: 41)
('Lemari Sepatu',        41, NOW(), NOW());  -- id: 100