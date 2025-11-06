"use client";
import { useEffect, useState } from "react";

export default function AdminDashboard() {
  const [products, setProducts] = useState([]);
  const [form, setForm] = useState({
    name: "",
    description: "",
    price: "",
    category: "",
    image: "",
  });
  const [file, setFile] = useState(null);
  const [editingProduct, setEditingProduct] = useState(null);

  const fetchProducts = async () => {
    const res = await fetch("http://localhost:8080/products");
    const data = await res.json();
    setProducts(data);
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  // ✅ Upload Image First → Then Add Product
//   const uploadImage = async () => {
//     if (!file) return "";

//     const formData = new FormData();
//     formData.append("image", file);

//     const res = await fetch("http://localhost:8080/upload", {
//       method: "POST",
//       body: formData,
//     });

//     const data = await res.json();
//     return data.url; // backend returns uploaded image URL
//   };

const uploadImage = async () => {
  if (!file) return "";

  const formData = new FormData();
  formData.append("image", file);

  const res = await fetch("http://localhost:8080/upload", {
    method: "POST",
    body: formData,
  });

  const data = await res.json();
  console.log("UPLOAD RESPONSE:", data);
  console.log("UPLOAD URL:", data.url);
  return data.url;
};


  const handleSubmit = async (e) => {
    e.preventDefault();
    const token = localStorage.getItem("token");

    // Upload image first
    let imageUrl = form.image;
    if (file) imageUrl = await uploadImage();

    // const payload = { ...form, image: imageUrl };

    const payload = {
        name: form.name,
        description: form.description,
        price: parseFloat(form.price),
        category: form.category,
        image: imageUrl,
    };


    const method = editingProduct ? "PUT" : "POST";
    const url = editingProduct
      ? `http://localhost:8080/admin/products/${editingProduct}`
      : "http://localhost:8080/admin/products";

    const res = await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(payload),
    });

    if (res.ok) {
      alert(editingProduct ? "✅ Product updated!" : "✅ Product added!");
      setForm({ name: "", description: "", price: "", category: "", image: "" });
      setFile(null);
      setEditingProduct(null);
      fetchProducts();
    } else {
      alert("❌ Operation failed");
    }
  };

  const deleteProduct = async (id) => {
    const token = localStorage.getItem("token");
    if (!confirm("Are you sure?")) return;

    await fetch(`http://localhost:8080/admin/products/${id}`, {
      method: "DELETE",
      headers: { Authorization: `Bearer ${token}` },
    });
    alert("🗑️ Deleted!");
    fetchProducts();
  };

  const startEdit = (p) => {
    setEditingProduct(p._id);
    setForm({
      name: p.name,
      description: p.description,
      price: p.price,
      category: p.category,
      image: p.image,
    });
  };

  return (
    <div className="p-6 max-w-6xl mx-auto">
      <h1 className="text-3xl font-bold mb-6 text-center">🛒 Admin Dashboard</h1>

      {/* Product Form */}
      <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-md mb-8">
        <h2 className="text-xl font-semibold mb-4">
          {editingProduct ? "✏️ Edit Product" : "➕ Add Product"}
        </h2>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <input
            type="text"
            placeholder="Product Name"
            value={form.name}
            onChange={(e) => setForm({ ...form, name: e.target.value })}
            className="border p-2 rounded"
            required
          />
          <input
            type="number"
            placeholder="Price"
            value={form.price}
            onChange={(e) => setForm({ ...form, price: e.target.value })}
            className="border p-2 rounded"
            required
          />
          <input
            type="text"
            placeholder="Category"
            value={form.category}
            onChange={(e) => setForm({ ...form, category: e.target.value })}
            className="border p-2 rounded"
            required
          />

          {/* 🖼️ File Upload */}
          <input
            type="file"
            accept="image/*"
            onChange={(e) => setFile(e.target.files[0])}
            className="border p-2 rounded"
          />
        </div>

        <textarea
          placeholder="Description"
          value={form.description}
          onChange={(e) => setForm({ ...form, description: e.target.value })}
          className="border p-2 rounded w-full mt-4"
          rows={3}
          required
        ></textarea>

        <button
          type="submit"
          className="mt-4 bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          {editingProduct ? "Save Changes" : "Add Product"}
        </button>
      </form>

      {/* Product List */}
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold mb-4">📦 Product List</h2>
        <table className="w-full border">
          <thead>
            <tr className="bg-gray-100">
              <th className="p-2 border">Image</th>
              <th className="p-2 border">Name</th>
              <th className="p-2 border">Price</th>
              <th className="p-2 border">Actions</th>
            </tr>
          </thead>
          <tbody>
            {products.map((p) => (
              <tr key={p._id} className="border-t">
                <td className="p-2 border">
                  {p.image ? (
                    <img
                      src={p.image}
                      alt={p.name}
                      className="w-16 h-16 object-cover rounded"
                    />
                  ) : (
                    "No image"
                  )}
                </td>
                <td className="p-2 border">{p.name}</td>
                <td className="p-2 border">₹{p.price}</td>
                <td className="p-2 border">
                  <button
                    onClick={() => startEdit(p)}
                    className="bg-yellow-500 text-white px-3 py-1 rounded mr-2"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => deleteProduct(p._id)}
                    className="bg-red-500 text-white px-3 py-1 rounded"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
