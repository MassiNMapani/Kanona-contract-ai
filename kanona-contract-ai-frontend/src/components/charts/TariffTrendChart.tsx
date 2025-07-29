// src/components/charts/TariffTrendChart.tsx
"use client";

import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from "chart.js";

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const data = {
  labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun"],
  datasets: [
    {
      label: "Tariff ($)",
      data: [100, 120, 115, 130, 125, 140],
      borderColor: "#4f46e5", // Tailwind indigo-600
      backgroundColor: "#c7d2fe"
    }
  ]
};

const options = {
  responsive: true,
  plugins: {
    legend: { position: "top" as const },
    title: { display: true, text: "Tariff Trends" }
  }
};

export default function TariffTrendChart() {
  return <Line data={data} options={options} />;
}
