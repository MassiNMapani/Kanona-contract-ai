// src/app/dashboard/contracts/index.tsx
"use client";

import RoleGuard from "../../../components/RoleGuard";
import ContractTable from "../../../components/ContractTable";
import ExpirationBarChart from "../../../components/charts/ExpirationBarChart";
import TariffTrendChart from "../../../components/charts/TariffTrendChart";
import VolumePieChart from "../../../components/charts/VolumePieChart";

export default function ContractsDashboard() {
  return (
    <RoleGuard allowedRoles={["admin", "ceo", "hod", "ppa-user", "psa-user", "viewer"]}>
      <div className="p-6 space-y-10">
        {/* Heading */}
        <div>
          <h1 className="text-3xl font-bold mb-2">📜 Contracts Dashboard</h1>
          <p className="text-gray-600">Track performance, volumes, expirations & more</p>
        </div>

        {/* Chart Section */}
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          <ExpirationBarChart />
          <TariffTrendChart />
          <VolumePieChart />
        </div>

        {/* Contract Table Section */}
        <div>
          <h2 className="text-xl font-semibold mb-4">📄 All Contracts</h2>
          <ContractTable />
        </div>
      </div>
    </RoleGuard>
  );
}
