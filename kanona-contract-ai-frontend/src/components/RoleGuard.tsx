import { useRole } from "../context/RoleContext";
import { useRouter } from "next/router";
import { useEffect } from "react";

interface Props {
  children: React.ReactNode;
  allowedRoles: string[];
}

const RoleGuard = ({ children, allowedRoles }: Props) => {
  const { role } = useRole();
  const router = useRouter();

  useEffect(() => {
    if (role && !allowedRoles.includes(role)) {
      router.push("/unauthorized");
    }
  }, [role, allowedRoles, router]);

  // Optional loading screen
  if (!role) {
    return <p className="text-center mt-4">🔐 Checking access...</p>;
  }

  return <>{children}</>;
};

export default RoleGuard;
