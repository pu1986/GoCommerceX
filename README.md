# 🛍️ GoCommerceX — Full-Stack E-Commerce Platform

GoCommerceX is a modern full-stack e-commerce application built using **Next.js, Golang (Gin) and MongoDB**.  
It includes **user authentication, admin dashboard, product CRUD, and image uploads**.

---

## 🚀 Tech Stack

| Layer | Technology |
|------|-----------|
Frontend | Next.js (App Router), React, Tailwind CSS  
Backend | Golang (Gin Framework)  
Database | MongoDB (Atlas / Local)  
Auth | JWT (JSON Web Token)  
Uploads | Local file storage `uploads/`  

---

## ✨ Features

### 🧑‍💻 User Features
- User Signup / Login (JWT based)
- Session-based login/logout
- View products with images, description & price
- Fully responsive UI

### 🛠️ Admin Features
- Admin login panel
- Add / Edit / Delete products (CRUD)
- Upload product images
- Protected Admin API (JWT + Admin Middleware)

---

## 📸 Screenshots

> Add images after running your app



📂 Create folder `/screens` and add screenshots later.

---

## 🧠 Folder Structure

```
GoCommerceX/
│
├── backend/
│   ├── config/        # DB connection
│   ├── controllers/   # API logic (user, product, image)
│   ├── middleware/    # Auth + Admin middlewares
│   ├── models/        # DB Schemas
│   ├── routes/        # Route definitions
│   └── main.go        # App entry
│
└── frontend/
    └── src/app/       # Next.js app router pages
```

---

## 🧪 API Endpoints

### 🔐 Auth
| Method | Endpoint | Description |
|---|---|---|
POST | `/signup` | Register new user  
POST | `/login` | Login and get JWT  

### 🛒 Products
| Method | Endpoint | Access |
|---|---|---|
GET | `/products` | Public  
GET | `/products/:id` | Public  
POST | `/admin/products` | Admin only  
PUT | `/admin/products/:id` | Admin only  
DELETE | `/admin/products/:id` | Admin only  

### 🖼️ Upload
| Method | Endpoint | Description |
|---|---|---|
POST | `/upload` | Upload image file  

---

## 🛠️ Setup Instructions

### ✅ Backend (Go)
```bash
cd backend
go mod tidy
go run main.go
```
Runs on 👉 **http://localhost:8080**

### ✅ Frontend (Next.js)
```bash
cd frontend
npm install
npm run dev
```
Runs on 👉 **http://localhost:3000**

---

## 🌐 Deployment Recommendations

| Service | Purpose |
|--------|--------|
Vercel | Frontend  
Render / Railway | Go backend  
MongoDB Atlas | Cloud database  

---

## 👨‍💻 Developer

**Pankaj Upadhyay**  
📧 Email: `upadhayay.pankaj1986@gmail.com`  
💼 GitHub: https://github.com/pu1986  

---

## ⭐ Support

Agar ye project pasand aaye toh **GitHub par Star ⭐ zarur dena!**  
Open-source developer ko support milta hai ❤️

---

### 📜 License
MIT License — free to use for learning and commercial projects.
