-- phpMyAdmin SQL Dump
-- version 4.9.4deb1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 23, 2020 at 02:46 AM
-- Server version: 8.0.20-0ubuntu0.19.10.1
-- PHP Version: 7.3.11-0ubuntu0.19.10.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tobalobs`
--

-- --------------------------------------------------------

--
-- Table structure for table `tambak`
--

CREATE TABLE `tambak` (
  `tambak_id` int NOT NULL,
  `user_id` int NOT NULL,
  `nama_tambak` varchar(500) NOT NULL,
  `panjang` float NOT NULL,
  `lebar` float NOT NULL,
  `jenis_budidaya` varchar(100) NOT NULL,
  `tanggal_mulai_budidaya` date NOT NULL,
  `usia_lobster` int NOT NULL,
  `jumlah_lobster` int NOT NULL,
  `jumlah_lobster_jantan` int DEFAULT '0',
  `jumlah_lobster_betina` int DEFAULT '0',
  `status` varchar(100) NOT NULL,
  `pakan_pagi` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '07:00',
  `pakan_sore` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '18:00',
  `ganti_air` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '3'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tambak`
--

INSERT INTO `tambak` (`tambak_id`, `user_id`, `nama_tambak`, `panjang`, `lebar`, `jenis_budidaya`, `tanggal_mulai_budidaya`, `usia_lobster`, `jumlah_lobster`, `jumlah_lobster_jantan`, `jumlah_lobster_betina`, `status`, `pakan_pagi`, `pakan_sore`, `ganti_air`) VALUES
(368, 182, 'Tambak pembesaran', 2, 2, 'pembesaran', '2020-06-29', 4, 40, 10, 30, 'aktif', '01:25', '14:38', '3'),
(496, 163, 'Tambak Pembenihan', 1, 2, 'pembenihan', '2020-07-14', 4, 20, 5, 15, 'aktif', '08:00', '17:00', '3'),
(532, 163, 'haha', 9, 5, 'pembesaran', '2020-07-19', 2, 450, 113, 337, 'aktif', '08:00', '17:00', '3'),
(551, 184, 'Tobalobs', 1.5, 1, 'pembesaran', '2020-07-20', 6, 15, 4, 11, 'aktif', '07:00', '18:00', '4'),
(552, 182, 'Test', 1, 1, 'pembesaran', '2020-07-22', 1, 10, 3, 7, 'aktif', '08:00', '17:00', '3'),
(553, 182, 'Test', 1, 1, 'pembesaran', '2020-07-22', 1, 10, 3, 7, 'aktif', '08:00', '17:00', '3'),
(565, 186, 'Tobalobs 2020', 1.5, 1, 'pembesaran', '2020-07-22', 6, 15, 4, 11, 'aktif', '08:00', '17:00', '3');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tambak`
--
ALTER TABLE `tambak`
  ADD PRIMARY KEY (`tambak_id`),
  ADD KEY `tambak_user` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tambak`
--
ALTER TABLE `tambak`
  MODIFY `tambak_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=566;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tambak`
--
ALTER TABLE `tambak`
  ADD CONSTRAINT `tambak_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
