# TourLink — Travel Ticket Booking Backend

**TourLink** is a scalable and modular backend application for managing travel/tour ticket bookings. Built using Golang with a clean architecture and Domain-Driven Design (DDD), it powers the creation, scheduling, and management of tour packages and user bookings with integrated JWT authentication, validation, and payment workflows.

---

## Features

- Tour Package Management (create, update, list)
- User Registration & Authentication (JWT)
- Tour Booking & Order Flow
- Payment Gateway Integration (planned)
- Review & Rating System (upcoming)
- Input Validation using Go Validator
- GORM ORM + PostgreSQL support
- Configurable with Viper
- Modular Domain-Driven Design
- RESTful API using GoFiber

---

## Tech Stack

| Layer         | Technology     |
|---------------|----------------|
| Language      | Golang         |
| Framework     | GoFiber        |
| ORM           | GORM           |
| Database      | PostgreSQL     |
| Auth          | JWT            |
| Validation    | Go Validator   |
| Config        | Viper          |
| Architecture  | Domain-Driven Design (DDD) |

---

## Folder Structure

```bash
tourlink/
├── cmd/                # Main entry point
├── configs/            # Config files (yaml, env)
├── internal/           # Domain modules
│   ├── user/           # User auth & profile
│   ├── tour/           # Tour destinations
│   ├── package/        # Tour packages
│   ├── booking/        # Ticket booking logic
│   ├── payment/        # Payment handling
│   ├── review/         # User reviews
│   └── shared/         # Shared logi
├── pkg/                # Infrastructure packages
│   ├── database/
│   ├── middleware/
│   ├── validator/
│   ├── config/
│   └── utils/
├── migrations/         # Auto migration scripts
├── go.mod
└── go.sum
