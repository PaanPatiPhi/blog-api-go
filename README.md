# blog-api-go

Rewrite ของ blog-api จากเดิมที่เขียนด้วย Node.js/Express.js มาเป็น Go เพื่อศึกษา backend architecture, API design และ performance ของ Go ecosystem

ใช้ Fiber เป็น web framework และ PostgreSQL เป็นฐานข้อมูล  
โปรเจคนี้ใช้เป็น backend สำหรับ personal-blog-react

---

## Current Status

โปรเจคอยู่ระหว่าง migrate จาก Express.js มาเป็น Go/Fiber  
ปัจจุบันมี API หลักสำหรับ Posts และ Categories แล้ว และกำลังพัฒนา authentication / authorization flow เพิ่มเติม

---

## Tech Stack

- Language: Go
- Framework: Fiber v2
- Database: PostgreSQL
- Database Library: sqlx / GORM
- Authentication: Supabase JWT (ECDSA / JWKS)
- Validation: go-playground/validator

---

## Project Structure

blog-api-go/
├── database/         # database connection setup
├── handlers/         # request handlers
├── middlewares/      # JWT auth / role middlewares
├── models/           # request/response structs
├── repositories/     # database queries
├── routes/           # route registration
└── main.go

---

## Request Flow

Client Request
    ↓
Middleware
    ↓
Handler
    ↓
Repository
    ↓
PostgreSQL

---

## Features

- REST API สำหรับ personal blog
- Pagination / filtering / search
- Role-based authorization middleware
- JWT authentication ด้วย Supabase
- PostgreSQL integration
- แยก project structure ตาม responsibility

---

## API Endpoints

### Posts

GET /posts/published - ดึงโพสต์ที่ published แล้ว
GET /posts/ - ดึงโพสต์ทั้งหมด (Admin)
GET /posts/:id - ดึงโพสต์ตาม ID
POST /posts/ - สร้างโพสต์ใหม่ (Admin)
PUT /posts/:id - แก้ไขโพสต์ (Admin)
DELETE /posts/:id - ลบโพสต์ (Admin)

### Categories

GET /categories/ - ดึง categories ทั้งหมด

### Users

GET /users - ดึง users ทั้งหมด

---

## Query Parameters

GET /posts/published

- page = หน้าที่ต้องการ
- limit = จำนวนโพสต์ต่อหน้า
- category = กรองตาม category
- search = ค้นหาจากชื่อโพสต์

---

## Environment Variables

DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

SUPABASE_URL=

---

## Getting Started

Prerequisites
- Go 1.21+
- PostgreSQL
- Supabase Project

Run Project

go mod tidy
go run main.go

Server:
http://localhost:4002

---

## Goals of This Project

- ฝึกการพัฒนา backend ด้วย Go
- ศึกษา project structure และ separation of concerns
- เรียนรู้ authentication / middleware flow
- ศึกษาการทำ REST API และ database integration
- เปรียบเทียบการพัฒนาระหว่าง Express.js และ Go/Fiber
