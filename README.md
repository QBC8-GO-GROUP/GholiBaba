# Goli Baba Integrated Travel Management System

This project is a microservices-based travel management platform, developed as part of the Quera Golang Bootcamp (Autumn 1403). The platform is designed to help users efficiently plan and manage travel bookings for buses, trains, flights, hotels, and more.

## Overview

Goli Baba aims to create an intelligent system for travel planning that provides users with the best options for tours and prices, ensuring minimal hassle and maximum satisfaction.

## Features

### User Features

- **User Registration and Login**: Secure registration with email and password.
- **Travel Search and Booking**: Find and book tickets for tours, hotels, and transport.
- **Wallet Integration**: Manage payments through an integrated wallet system.
- **Booking Cancellation**: Easily cancel bookings and get refunds.
- **Notifications**: Receive updates on bookings and relevant travel information.

### Admin Features

- **User and Company Management**: View and manage users and service providers.
- **Blocking Functionality**: Block users or companies for policy violations.
- **Access Control**: Define and manage access roles for users.
- **End-of-Trip Settlements**: Manage and finalize payments after trips.

### Service Provider Features

- **Dynamic Travel Offerings**: Add new types of trips (bus, train, flight, etc.).
- **Pricing and Availability Management**: Adjust rates and availability dynamically.
- **Revenue Sharing**: Automate revenue distribution after trip completion.

### Backend Features

- **Microservices Architecture**: Modular design for scalability and maintainability.
- **gRPC Communication**: Efficient service-to-service communication.
- **Queueing System**: Improved performance using queueing mechanisms.
- **Database**: PostgreSQL for data storage.
- **Caching**: Enhanced performance with caching.
- **Security**: Strong focus on data security using modern encryption techniques.

## Entity Relationship Diagram (ERD)

Below is the ERD representing the main entities in the system:

```plaintext
Travel
+-------------+
| id          |
| user_id     |
| travel_id   |
| status      |
| created_at  |
| updated_at  |
| deleted_at  |
+-------------+

Vehicle
+-------------+
| id          |
| company_id  |
| type        |
| source      |
| destination |
| start_time  |
| end_time    |
| price       |
| seats       |
| available   |
| approved    |
| vehicle_id  |
| created_at  |
| updated_at  |
| deleted_at  |
+-------------+

Wallet
+-------------+
| id          |
| user_id     |
| money       |
| created_at  |
| updated_at  |
| deleted_at  |
+-------------+

Cards
+-------------+
| id          |
| numbers     |
| wallet_id   |
| created_at  |
| updated_at  |
| deleted_at  |
+-------------+

History
+-------------+
| id          |
| code        |
| is_approved |
| price       |
| source      |
| destination |
| title       |
| description |
| created_at  |
| updated_at  |
| deleted_at  |
+-------------+
```
