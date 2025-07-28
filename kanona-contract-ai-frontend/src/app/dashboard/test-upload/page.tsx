import FileUpload from "../../../components/FileUpload";


export default function TestUploadPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-6">
      <div className="w-full max-w-md">
        <h1 className="text-xl font-semibold mb-4">Upload a File</h1>
        <FileUpload />
      </div>
    </div>
  );
}
