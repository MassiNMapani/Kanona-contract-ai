import RoleGuard from "../../../components/RoleGuard";

export default function ContractsDashboard() {
  return (
    <RoleGuard allowedRoles={["admin", "ceo", "hod", "ppa-user", "psa-user", "viewer"]}>
      <div className="p-6">
        <h1 className="text-2xl font-bold mb-4">📜 Contracts Overview</h1>
        {/* Contract Table will go here */}
      </div>
    </RoleGuard>
  );
}
