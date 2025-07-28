// src/components/Sidebar.jsx
import { useState } from 'react';
import { Menu, X, UploadCloud, File, Trash2 } from 'lucide-react';

export default function Sidebar() {
  const [isOpen, setIsOpen] = useState(true);

  return (
    <aside className={`bg-gray-100 min-h-screen p-4 w-64 transition-all ${isOpen ? 'block' : 'hidden'} md:block`}>
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-lg font-semibold">Menu</h2>
        <button onClick={() => setIsOpen(!isOpen)} className="md:hidden">
          {isOpen ? <X size={20} /> : <Menu size={20} />}
        </button>
      </div>
      <ul className="space-y-4">
        <li><a href="#" className="text-blue-600 font-medium">Dashboard</a></li>
        <li><a href="#" className="text-gray-700 hover:text-blue-600">File Upload</a></li>
        <li><a href="#" className="text-gray-700 hover:text-blue-600">Documents</a></li>
        <li><a href="#" className="text-gray-700 hover:text-blue-600">Settings</a></li>
      </ul>
    </aside>
  );
}
