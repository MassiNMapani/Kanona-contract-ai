import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080/',
  headers: {
    'Content-Type': 'application/json'
  }
})

export async function getCurrentUser() {
  const res = await fetch("http://localhost:8080/me", {
    method: "GET",
    credentials: "include", // ⬅️ IMPORTANT!
  });

  if (!res.ok) throw new Error("Not authenticated");
  return await res.json();
}

export const fetchContracts = async () => {
  try {
    const res = await fetch("http://localhost:8080/contracts", {
      method: "GET",
      credentials: "include", // ✅ Send cookies (JWT is stored in cookie)
    });

    if (!res.ok) throw new Error("Failed to fetch contracts");

    const data = await res.json();
    return data;
  } catch (err) {
    console.error("⚠️ Error fetching contracts:", err);
    return [];
  }
};

export default api

