"use client";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function Navbar() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    const userData = localStorage.getItem("user");

    if (token && userData) {
      setIsLoggedIn(true);
      try {
        setUser(JSON.parse(userData)); // ✅ convert string to object
      } catch (err) {
        console.error("Error parsing user data:", err);
      }
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    setIsLoggedIn(false);
    setUser(null);
    alert("Logged out successfully!");
    window.location.href = "/";
  };

  return (
    <nav className="bg-white shadow-md px-6 py-4 flex justify-between items-center">
      <Link href="/" className="text-xl font-bold text-blue-600">
        GoCommerce 🛍️
      </Link>

      <div className="flex items-center space-x-6">
        {isLoggedIn && user ? (
          <>
            <span className="text-gray-700 font-medium">
              👋 Welcome,{" "}
              <span className="text-blue-600 font-semibold">{user.name}</span>
            </span>
            <button
              onClick={handleLogout}
              className="text-gray-600 hover:text-red-600 font-semibold"
            >
              Logout
            </button>
          </>
        ) : (
          <>
            <Link href="/login" className="text-gray-600 hover:text-blue-600">
              Login
            </Link>
            <Link href="/signup" className="text-gray-600 hover:text-blue-600">
              Signup
            </Link>
          </>
        )}
      </div>
    </nav>
  );
}
