// src/components/charts/VolumePieChart.tsx
"use client";

import { Pie } from "react-chartjs-2";
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend
} from "chart.js";

ChartJS.register(ArcElement, Tooltip, Legend);

const data = {
  labels: ["Water", "Electricity", "Telecom"],
  datasets: [
    {
      label: "Volume by Sector",
      data: [12, 19, 9],
      backgroundColor: [
        "#3b82f6", // blue
        "#10b981", // green
        "#f472b6"  // pink
      ],
      borderWidth: 1
    }
  ]
};

const options = {
  responsive: true,
  plugins: {
    title: { display: true, text: "Contract Volume by Sector" }
  }
};

export default function VolumePieChart() {
  return <Pie data={data} options={options} />;
}
