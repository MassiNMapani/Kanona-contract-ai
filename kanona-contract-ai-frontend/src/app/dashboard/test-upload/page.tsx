// src/app/dashboard/test-upload/page.tsx
"use client";

import { useRole } from "@/context/RoleContext";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import FileUpload from "@/components/FileUpload";

export default function TestUploadPage() {
  const { role } = useRole();
  const router = useRouter();

  useEffect(() => {
    // Redirect if user does not have permission
    if (role && !["admin", "ppa-user", "psa-user"].includes(role)) {
      router.push("/unauthorized");
    }
  }, [role]);

  // If role is not yet loaded, you may show a loader
  if (!role) {
    return <p className="text-center mt-10">Checking access...</p>;
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-6">
      <div className="w-full max-w-md">
        <h1 className="text-xl font-semibold mb-4">Upload a File</h1>
        <FileUpload />
      </div>
    </div>
  );
}
