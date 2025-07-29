// src/components/charts/ExpirationBarChart.tsx
"use client";

import { Bar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
} from "chart.js";

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const data = {
  labels: ["Q3 2024", "Q4 2024", "Q1 2025", "Q2 2025"],
  datasets: [
    {
      label: "Expiring Contracts",
      data: [6, 9, 4, 7],
      backgroundColor: "#facc15" // Tailwind yellow-400
    }
  ]
};

const options = {
  responsive: true,
  plugins: {
    legend: { position: "top" as const },
    title: { display: true, text: "Contract Expirations" }
  }
};

export default function ExpirationBarChart() {
  return <Bar data={data} options={options} />;
}
