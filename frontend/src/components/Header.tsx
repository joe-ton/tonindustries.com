// Header.tsx

import { Link } from 'react-router-dom';

export default function Header() {
  return (
    <header className="w-full p-4 bg-gray-800 text-white flex justify-between">
      <h1 className="text-xl font-semibold">TonIndustries</h1>
      <nav className="space-x-4">
        <Link to="/" className="hover:underline">Home</Link>
        <Link to="/articles" className="hover:underline">Articles</Link>
      </nav>
    </header>
  );
}

