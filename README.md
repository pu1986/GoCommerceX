🛍️ GoCommerceX

Full Stack E-Commerce Application built with Next.js, Golang (Gin), and MongoDB

🚀 Modern shopping platform with secure authentication, admin dashboard, image uploads, and full product CRUD management.

🧠 Tech Stack
Layer	Technology
Frontend	Next.js
 (React + App Router)
Backend	Go (Gin Framework)

Database	MongoDB Atlas

Auth	JWT (JSON Web Tokens)
Styling	Tailwind CSS
Image Uploads	Local /uploads folder or S3 (optional)
✨ Features
👤 User Features

🔐 Signup / Login with JWT authentication

🧾 Session Management (login persistence & logout)

🛍️ Browse Products with name, image, price & description

📱 Responsive UI for desktop and mobile

🧑‍💼 Admin Features

🧱 Admin Dashboard with CRUD operations

🖼️ Image Uploads via /upload API

🚫 Protected Routes (AuthMiddleware + AdminMiddleware)

✅ Approve / Manage Products in real-time

⚙️ Project Structure
gocommercex/
│
├── backend/
│   ├── config/          # MongoDB connection setup
│   ├── controllers/     # Business logic (user, product, upload)
│   ├── middleware/      # JWT auth & admin protection
│   ├── models/          # MongoDB schemas
│   ├── routes/          # API routes
│   └── main.go          # Entry point
│
└── frontend/
    ├── src/app/         # Next.js App Router pages
    ├── src/components/  # Navbar, ProductCard, etc.
    ├── public/          # Static assets
    └── package.json

🧩 API Endpoints
🔑 Authentication
Method	Endpoint	Description
POST	/signup	Register a new user
POST	/login	Login & get JWT token
🛍️ Products
Method	Endpoint	Description
GET	/products	Fetch all products
GET	/products/:id	Get product by ID
POST	/admin/products	Create new product (Admin only)
PUT	/admin/products/:id	Update product (Admin only)
DELETE	/admin/products/:id	Delete product (Admin only)
🖼️ Image Upload
Method	Endpoint	Description
POST	/upload	Upload product image
🧭 Local Setup
🧩 Backend
cd backend
go mod tidy
go run main.go
# Server running at http://localhost:8080

🖥️ Frontend
cd frontend
npm install
npm run dev
# Frontend running at http://localhost:3000

🌍 Deployment
Service	Purpose	Link
Frontend	Deploy with Vercel
	
Backend	Use Render
 or Railway
	
Database	MongoDB Atlas Cloud
	
🧑‍💻 Developer

👨‍💻 Pankaj Upadhyay
📧 Email: upadhayay.pankaj1986@gmail.com

💼 GitHub: github.com/pu1986

🚀 Portfolio (Coming soon): https://pankajdev.me (optional placeholder)

🪄 Future Enhancements

🔎 Product search & filtering

🛒 Shopping cart & checkout flow

💳 Razorpay / Stripe payment integration

📦 Admin order management dashboard

📜 License

This project is licensed under the MIT License.
Feel free to fork, modify, and build upon this project.

⭐ Support Open Source
If you liked this project, please give it a ⭐ on GitHub — it helps me grow as a developer 🙌
