import React, { useEffect, useState } from 'react';
import api from './../../services/api';

export default function Dashboard() {
  const [contracts, setContracts] = useState([]);

  useEffect(() => {
    const fetchContracts = async () => {
      const res = await api.get('/contracts');
      setContracts(res.data);
    };
    fetchContracts();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <h1 className="text-2xl font-bold mb-4">📊 Contract Insights Dashboard</h1>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-white p-4 rounded shadow">
          <h2 className="text-sm text-gray-500">Total Contracts</h2>
          <p className="text-2xl font-semibold">{contracts.length}</p>
        </div>
        <div className="bg-white p-4 rounded shadow">
          <h2 className="text-sm text-gray-500">Latest Upload</h2>
          <p className="text-md">{contracts[0]?.filename || 'N/A'}</p>
        </div>
        <div className="bg-white p-4 rounded shadow">
          <h2 className="text-sm text-gray-500">Pending Alerts</h2>
          <p className="text-md">3 (simulated)</p>
        </div>
      </div>

      <div className="mt-6 bg-white p-4 rounded shadow">
        <h3 className="font-semibold mb-2">Recent Contracts</h3>
        <ul className="space-y-1">
          {contracts.slice(0, 5).map((c, i) => (
            <li key={i} className="text-sm text-blue-600 hover:underline">
              {c.filename}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}