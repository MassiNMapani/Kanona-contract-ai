// src/components/Layout.jsx
import Navbar from './Navbar';
import Sidebar from './Sidebar';

export default function Layout({ children }) {
  return (
    <div className="flex h-screen overflow-hidden">
      <Sidebar />
      <div className="flex flex-col flex-1">
        <Navbar />
        <main className="p-6 bg-gray-50 overflow-y-auto flex-1">
          {children}
        </main>
      </div>
    </div>
  );
}
