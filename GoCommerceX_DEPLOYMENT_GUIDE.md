# 🚀 GoCommerceX Deployment Guide

This guide helps you deploy your **GoCommerceX** (Next.js + Golang + MongoDB) project to production.

---

## 🧩 Project Overview

| Layer | Technology |
|------|-------------|
Frontend | Next.js (React + Tailwind CSS)  
Backend | Go (Gin Framework)  
Database | MongoDB Atlas  
Auth | JWT-based Authentication  
Uploads | Local or Cloud Storage (like AWS S3)  

---

## 🌐 Deployment Overview

| Service | Role |
|----------|------|
**Vercel** | Deploys the Next.js frontend  
**Render / Railway / Fly.io** | Deploys the Golang backend  
**MongoDB Atlas** | Cloud database service  

---

## ⚙️ STEP 1 — Setup MongoDB Atlas

1. Visit 👉 [https://www.mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
2. Create a free cluster.
3. Add your IP (0.0.0.0/0 for testing).  
4. Create a **Database User** with username and password.
5. Copy your connection string, for example:

```
mongodb+srv://<username>:<password>@cluster0.mongodb.net/gocommerce
```

6. Replace it in your Go backend file (`backend/config/db.go`).

---

## 🧠 STEP 2 — Deploy Backend (Render.com)

### Create a new Render service

1. Visit [https://render.com](https://render.com)
2. Click **New → Web Service**
3. Connect your GitHub repository (`GoCommerceX`)
4. Select the `/backend` folder as the root.

### Configure settings:

| Setting | Value |
|----------|--------|
Environment | Go  
Build Command | `go build -o main .`  
Start Command | `./main`  
Environment Variables | Add your MongoDB connection string, e.g.  
`MONGODB_URI=mongodb+srv://...`  

Once deployed, Render will give you a backend URL, for example:

```
https://gocommercex-backend.onrender.com
```

---

## 🖥️ STEP 3 — Deploy Frontend (Vercel)

1. Go to [https://vercel.com](https://vercel.com)
2. Click **Add New Project → Import GitHub Repository**
3. Choose your GoCommerceX repo.
4. Set the root directory to `/frontend`
5. Click **Deploy**

### Update API URLs

In `frontend/src` files where API calls are made (like `/login`, `/products`),  
replace:

```js
http://localhost:8080
```
with your backend URL, e.g.
```js
https://gocommercex-backend.onrender.com
```

Then redeploy.

---

## ☁️ Optional — Uploads to S3 (Production Ready)

If you want image uploads to persist (not local):
- Create an AWS S3 bucket
- Use the Go AWS SDK in your `/upload` controller instead of saving to `/uploads`
- Store the returned S3 URL in MongoDB

---

## 🧩 STEP 4 — Final Testing

✅ Check Backend health:  
`https://gocommercex-backend.onrender.com/` → should show `GoCommerce Backend Running 🚀`  

✅ Check Frontend:  
`https://gocommercex.vercel.app/` → should load your site  

✅ Try Login, Product Add, and CRUD features.

---

## 🧑‍💻 Developer Info

**👨‍💻 Pankaj Upadhyay**  
📧 Email: `upadhayay.pankaj1986@gmail.com`  
💼 GitHub: [github.com/pu1986](https://github.com/pu1986)  

---

### 🏁 Congratulations! 🎉  
Your full-stack e-commerce platform **GoCommerceX** is live!  

⭐ Don’t forget to star the repo and share it with recruiters.
