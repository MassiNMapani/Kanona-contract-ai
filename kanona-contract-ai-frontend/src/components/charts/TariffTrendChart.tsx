"use client";

import { useEffect, useState } from "react";
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
import { format } from "date-fns";

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

interface Contract {
  id: string;
  name: string;
  tariff: number;
  startDate: string;
}

export default function TariffTrendChart() {
  const [chartData, setChartData] = useState<any>(null);

  useEffect(() => {
    const fetchContracts = async () => {
      try {
        const response = await fetch("http://localhost:8080/contracts", {
          credentials: "include", // Include JWT cookie
        });

        const contracts: Contract[] = await response.json();

        // Group tariffs by month
        const monthMap: { [key: string]: number[] } = {};
        contracts.forEach((c) => {
          if (c.tariff && c.startDate) {
            const month = format(new Date(c.startDate), "yyyy-MM");
            if (!monthMap[month]) monthMap[month] = [];
            monthMap[month].push(c.tariff);
          }
        });

        // Format data for chart
        const labels = Object.keys(monthMap).sort();
        const data = labels.map((month) => {
          const tariffs = monthMap[month];
          const avg = tariffs.reduce((sum, val) => sum + val, 0) / tariffs.length;
          return Number(avg.toFixed(2));
        });

        setChartData({
          labels,
          datasets: [
            {
              label: "Tariff ($)",
              data,
              borderColor: "#4f46e5",
              backgroundColor: "#c7d2fe"
            }
          ]
        });
      } catch (error) {
        console.error("Failed to load tariff trend data:", error);
      }
    };

    fetchContracts();
  }, []);

  const options = {
    responsive: true,
    plugins: {
      legend: { position: "top" as const },
      title: { display: true, text: "Tariff Trends" }
    }
  };

  if (!chartData) {
    return <div className="p-4 text-gray-600">Loading tariff trends...</div>;
  }

  return <Line data={chartData} options={options} />;
}
