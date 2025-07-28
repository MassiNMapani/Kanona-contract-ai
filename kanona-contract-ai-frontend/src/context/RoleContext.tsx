// src/context/RoleContext.tsx
"use client";

import {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
} from "react";

type Role = "admin" | "ceo" | "hod" | "ppa-user" | "psa-user" | "viewer" | null;

interface RoleContextType {
  role: Role;
  setRole: (role: Role) => void;
}

const RoleContext = createContext<RoleContextType | undefined>(undefined);

export const RoleProvider = ({ children }: { children: ReactNode }) => {
  const [role, setRole] = useState<Role>(null);

  useEffect(() => {
    const fetchRole = async () => {
      try {
        const res = await fetch("http://localhost:8080/me", {
          credentials: "include", // ⬅️ ensure cookies are sent
        });

        if (!res.ok) throw new Error("Unauthorized");

        const data = await res.json();
        setRole(data.role as Role); // set role from backend response
      } catch (err) {
        console.error("⚠️ Error fetching user role:", err);
        setRole(null); // fallback to null if not logged in
      }
    };

    fetchRole(); // ⬅️ Call it inside useEffect
  }, []); // ⬅️ Empty dependency array = only run once on mount

  return (
    <RoleContext.Provider value={{ role, setRole }}>
      {children}
    </RoleContext.Provider>
  );
};

export const useRole = () => {
  const context = useContext(RoleContext);
  if (!context)
    throw new Error("useRole must be used within a RoleProvider");
  return context;
};
