// src/context/RoleContext.tsx
import { createContext, useContext, useState, useEffect, ReactNode } from "react";

type Role = "admin" | "ceo" | "hod" | "ppa-user" | "psa-user" | "viewer" | null;

interface RoleContextType {
  role: Role;
  setRole: (role: Role) => void;
}

const RoleContext = createContext<RoleContextType | undefined>(undefined);

export const RoleProvider = ({ children }: { children: ReactNode }) => {
  const [role, setRole] = useState<Role>(null);

  useEffect(() => {
    const user = JSON.parse(localStorage.getItem("user") || "{}");
    setRole(user?.role || null);
  }, []);

  return (
    <RoleContext.Provider value={{ role, setRole }}>
      {children}
    </RoleContext.Provider>
  );
};

export const useRole = () => {
  const context = useContext(RoleContext);
  if (!context) throw new Error("useRole must be used within RoleProvider");
  return context;
};
