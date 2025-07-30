"use client";

import { useEffect, useState } from "react";
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
import { fetchContracts } from "../../../services/api";
import { getQuarter } from "date-fns";

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

interface Contract {
  endDate?: string;
}

export default function ExpirationBarChart() {
  const [chartData, setChartData] = useState<any>({
    labels: [],
    datasets: []
  });

  useEffect(() => {
    const loadChartData = async () => {
      const contracts: Contract[] = await fetchContracts();

      // Map of quarter-year -> count
      const expirationMap: Record<string, number> = {};

      contracts.forEach((contract) => {
        if (!contract.endDate) return;

        const date = new Date(contract.endDate);
        const quarter = getQuarter(date);
        const year = date.getFullYear();
        const label = `Q${quarter} ${year}`;

        expirationMap[label] = (expirationMap[label] || 0) + 1;
      });

      // Sort quarters chronologically
      const sortedLabels = Object.keys(expirationMap).sort((a, b) => {
        const [qA, yA] = a.split(" ");
        const [qB, yB] = b.split(" ");
        return (
          parseInt(yA) - parseInt(yB) || parseInt(qA[1]) - parseInt(qB[1])
        );
      });

      const data = {
        labels: sortedLabels,
        datasets: [
          {
            label: "Expiring Contracts",
            data: sortedLabels.map((label) => expirationMap[label]),
            backgroundColor: "#facc15" // Tailwind yellow-400
          }
        ]
      };

      setChartData(data);
    };

    loadChartData();
  }, []);

  const options = {
    responsive: true,
    plugins: {
      legend: { position: "top" as const },
      title: { display: true, text: "Contract Expirations" }
    }
  };

  return <Bar data={chartData} options={options} />;
}
