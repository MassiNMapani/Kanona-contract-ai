"use client";

import { useEffect, useState } from "react";
import { fetchContracts } from "../../services/api";
import { useRole } from "../context/RoleContext";
import { format } from "date-fns";
import { Bar } from "react-chartjs-2";
import "chart.js/auto";

interface Contract {
  id: string;
  name: string;
  type: "ppa" | "psa";
  startDate: string;
  endDate: string;
  tariff: number;
  volume: number;
  renegotiationDate: string;
}

export default function ContractTable() {
  const { role } = useRole();
  const [contracts, setContracts] = useState<Contract[]>([]);
  const [filter, setFilter] = useState<"ppa" | "psa" | "all">("all");

  useEffect(() => {
    const loadContracts = async () => {
      const data = await fetchContracts(); // JWT-auth fetch
      setContracts(data);
    };
    loadContracts();
  }, []);

  const canEdit = ["admin", "hod"].includes(role!);
  const canDelete = ["admin", "hod"].includes(role!);
  const canUpload = ["admin", "ppa-user", "psa-user", "hod"].includes(role!);
  const canView = role !== "viewer" || role === "viewer"; // viewer access assumed to be filtered backend-side

  const filteredContracts =
    filter === "all" ? contracts : contracts.filter((c) => c.type === filter);

  const chartData = {
    labels: filteredContracts.map((c) => c.name),
    datasets: [
      {
        label: "Contract Volumes (MWh)",
        data: filteredContracts.map((c) => c.volume),
        backgroundColor: "rgba(54, 162, 235, 0.6)",
      },
    ],
  };

  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Contracts</h1>
        <div className="space-x-2">
          <button onClick={() => setFilter("all")} className="btn">
            All
          </button>
          <button onClick={() => setFilter("ppa")} className="btn">
            PPA
          </button>
          <button onClick={() => setFilter("psa")} className="btn">
            PSA
          </button>
        </div>
      </div>

      {canUpload && (
        <button className="bg-blue-600 text-white px-4 py-2 rounded mb-4">
          Upload New Contract
        </button>
      )}

      {canView && (
        <div className="overflow-x-auto bg-white rounded shadow">
          <table className="min-w-full table-auto">
            <thead>
              <tr className="bg-gray-100 text-left">
                <th className="p-2">Name</th>
                <th className="p-2">Type</th>
                <th className="p-2">Start</th>
                <th className="p-2">End</th>
                <th className="p-2">Tariff</th>
                <th className="p-2">Volume</th>
                <th className="p-2">Renegotiation</th>
                {(canEdit || canDelete) && <th className="p-2">Actions</th>}
              </tr>
            </thead>
            <tbody>
              {filteredContracts.map((contract) => (
                <tr key={contract.id} className="border-t hover:bg-gray-50">
                  <td className="p-2">{contract.name}</td>
                  <td className="p-2 uppercase">{contract.type}</td>
                  <td className="p-2">
                    {format(new Date(contract.startDate), "yyyy-MM-dd")}
                  </td>
                  <td className="p-2">
                    {format(new Date(contract.endDate), "yyyy-MM-dd")}
                  </td>
                  <td className="p-2">${contract.tariff.toFixed(2)}</td>
                  <td className="p-2">{contract.volume}</td>
                  <td className="p-2">{contract.renegotiationDate}</td>
                  {(canEdit || canDelete) && (
                    <td className="p-2 space-x-2">
                      {canEdit && (
                        <button className="text-blue-600 hover:underline">
                          Edit
                        </button>
                      )}
                      {canDelete && (
                        <button className="text-red-600 hover:underline">
                          Delete
                        </button>
                      )}
                    </td>
                  )}
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      <div className="mt-8 bg-white p-4 rounded shadow">
        <h2 className="text-lg font-semibold mb-2">📊 Volume Overview</h2>
        <Bar data={chartData} />
      </div>
    </div>
  );
}
