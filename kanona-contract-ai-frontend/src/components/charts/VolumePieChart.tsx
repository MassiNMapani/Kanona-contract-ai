"use client";

import { useEffect, useState } from "react";
import { Pie } from "react-chartjs-2";
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend
} from "chart.js";

ChartJS.register(ArcElement, Tooltip, Legend);

interface Contract {
  id: string;
  type: "ppa" | "psa";
  volume: number;
}

export default function VolumePieChart() {
  const [chartData, setChartData] = useState<any>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await fetch("http://localhost:8080/contracts", {
          credentials: "include", // Include token cookie
        });
        const contracts: Contract[] = await res.json();

        // Aggregate volume by type
        let ppaVolume = 0;
        let psaVolume = 0;

        for (const c of contracts) {
          if (c.type === "ppa") ppaVolume += c.volume || 0;
          else if (c.type === "psa") psaVolume += c.volume || 0;
        }

        setChartData({
          labels: ["PPA", "PSA"],
          datasets: [
            {
              label: "Volume by Contract Type",
              data: [ppaVolume, psaVolume],
              backgroundColor: ["#3b82f6", "#10b981"], // blue & green
              borderWidth: 1
            }
          ]
        });
      } catch (error) {
        console.error("Error loading contract volume data:", error);
      }
    };

    fetchData();
  }, []);

  const options = {
    responsive: true,
    plugins: {
      title: { display: true, text: "Contract Volume by Type" }
    }
  };

  if (!chartData) return <div className="p-4 text-gray-500">Loading volume chart...</div>;

  return <Pie data={chartData} options={options} />;
}
