import { useRole } from "../../context/RoleContext";

export default function RoleTestPage() {
  const { role } = useRole();

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold">👤 Current Role: {role || "None"}</h1>
    </div>
  );
}
