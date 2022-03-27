-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 27, 2022 at 02:27 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `codefood-db`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_detail_recipes`
--

CREATE TABLE `tb_detail_recipes` (
  `id` int(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `recipe_category_id` int(11) NOT NULL,
  `image` varchar(50) NOT NULL,
  `n_reaction_like` int(11) NOT NULL,
  `n_reaction_neutral` int(11) NOT NULL,
  `n_reacton_dislike` int(11) NOT NULL,
  `n_serving` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tb_recipe`
--

CREATE TABLE `tb_recipe` (
  `id` int(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `recipe_category_id` int(11) NOT NULL,
  `image` varchar(50) NOT NULL,
  `n_reaction_like` int(11) NOT NULL,
  `n_reaction_neutral` int(11) NOT NULL,
  `n_reaction_dislike` int(11) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tb_recipe_category`
--

CREATE TABLE `tb_recipe_category` (
  `id` int(11) NOT NULL,
  `name` varchar(30) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_detail_recipes`
--
ALTER TABLE `tb_detail_recipes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tb_recipe`
--
ALTER TABLE `tb_recipe`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tb_recipe_category`
--
ALTER TABLE `tb_recipe_category`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_detail_recipes`
--
ALTER TABLE `tb_detail_recipes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `tb_recipe`
--
ALTER TABLE `tb_recipe`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `tb_recipe_category`
--
ALTER TABLE `tb_recipe_category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
