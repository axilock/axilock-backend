import React, { useEffect, useState } from "react";
import Header from "../Common/Header";
import { FaSpinner } from "react-icons/fa";

const EmptyPage = ({ darkTheme }) => {

  const [loading, setLoading] = useState(true); // Loading state
  const [error, setError] = useState(null); // Error state

  setTimeout(() => {
    setLoading(false);
  }, 2000)


  if (loading) return (
    <div className="min-h-screen">
      <div className={`flex items-center justify-center h-screen ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
        <FaSpinner className={`animate-spin text-4xl ${darkTheme ? 'text-green-400' : 'text-green-500'}`} />
      </div>
    </div>
  );

  if (error) return (
    <div className="main-body">
      <div className="main">
        <div className="home-container">
        </div>
      </div>
      <div className="main-banner">
        <div className="banner">
          <div className="main-title">
            <p>{error}</p>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <>
      <div className={`min-h-screen 
        ${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}
      `}>
        {/* Banner Section */}
        <div className={`p-6 ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
          <div className="mb-4">
            <h1 className={`text-xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Empty Page</h1>
            <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'} text-sm`}>Some contents.</p>
          </div>
        </div>
      </div>
    </>
  );
};

export default EmptyPage;