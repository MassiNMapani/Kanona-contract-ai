import { useRole } from "../context/RoleContext";
import { useRouter } from "next/router";

const roles = ["admin", "ceo", "hod", "ppa-user", "psa-user", "viewer"];

export default function MockLogin() {
  const { setRole } = useRole();
  const router = useRouter();

  const handleLogin = (role: string) => {
    localStorage.setItem("user", JSON.stringify({ role }));
    setRole(role as any);
    router.push("/dashboard/contracts");
  };

  return (
    <div className="p-10 space-y-4">
      <h2 className="text-xl font-semibold">Login as:</h2>
      {roles.map((role) => (
        <button
          key={role}
          onClick={() => handleLogin(role)}
          className="px-4 py-2 bg-blue-500 text-white rounded"
        >
          {role}
        </button>
      ))}
    </div>
  );
}
