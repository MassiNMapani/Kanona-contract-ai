"use client";

import { useState } from "react";

export default function FileUpload() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setSelectedFile(file);
      console.log("Selected file:", file);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    setUploading(true);

    try {
      const formData = new FormData();
      formData.append("file", selectedFile);

      // Replace this URL with your Go backend API
      const response = await fetch("http://localhost:8080/upload", {
        method: "POST",
        body: formData,
        credentials: "include",
      });

      if (!response.ok) throw new Error("Upload failed");

      const result = await response.json();
      console.log("Success:", result);
    } catch (error) {
      console.error("Error uploading file:", error);
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-4 border rounded shadow bg-white">
      <h2 className="text-xl font-semibold mb-4">Upload a File</h2>

      <input
        type="file"
        accept="*"
        onChange={handleFileChange}
        className="mb-4"
      />

      {selectedFile && (
        <div className="mb-4">
          <p className="text-sm text-gray-600">
            Selected: <strong>{selectedFile.name}</strong>
          </p>
        </div>
      )}

      <button
        onClick={handleUpload}
        disabled={!selectedFile || uploading}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50"
      >
        {uploading ? "Uploading..." : "Submit"}
      </button>
    </div>
  );
}
