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

export default api

