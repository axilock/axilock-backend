import React, { useEffect, useState } from "react";
import Header from "../Common/Header";
import Edit from "../Images/edit.svg";
import deleteicon from "../Images/delete.svg";
import { IoCheckmark } from "react-icons/io5";
import { FaSpinner } from "react-icons/fa";

const Secretscan = ({ darkTheme }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const toggleDropdown = () => setIsOpen(!isOpen);

  useEffect(() => {
    // Simulate loading
    const timer = setTimeout(() => {
      setLoading(false);
    }, 2000);

    return () => clearTimeout(timer);
  }, []);

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
            <p className={`${darkTheme ? 'text-red-400' : 'text-red-600'}`}>{error}</p>
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
          <div className="mb-6">
            <h1 className={`text-2xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
              Secret Scanning Rules
            </h1>
          </div>

          {/* Secret Scan Table Container */}
          <div className={`rounded-lg border ${darkTheme ? 'border-gray-700 bg-gray-900' : 'border-gray-200 bg-white'} overflow-hidden`}>
            
            {/* Header Section */}
            <div className={`flex justify-between items-center p-4 border-b ${darkTheme ? 'border-gray-700 bg-gray-800' : 'border-gray-200 bg-gray-50'}`}>
              <div>
                <p className={`font-medium ${darkTheme ? 'text-white' : 'text-gray-800'}`}>0 Patterns</p>
              </div>
              
              {/* Filter Dropdown */}
              {/*   */}
            </div>

            {/* Rules List */}
            <div className="divide-y divide-gray-200 dark:divide-gray-700">
              
              {/* Rule Item 1 */}
              {/* <div className={`flex items-center justify-between p-4 hover:bg-opacity-50 transition-colors duration-150
                ${darkTheme ? 'hover:bg-gray-700' : 'hover:bg-gray-50'}
              `}>
                <div className="flex items-center space-x-4">
                  <input 
                    type="checkbox" 
                    className={`h-4 w-4 rounded border focus:ring-2 focus:ring-offset-2
                      ${darkTheme 
                        ? 'bg-gray-700 border-gray-600 text-green-500 focus:ring-green-400 focus:ring-offset-gray-800' 
                        : 'bg-white border-gray-300 text-green-600 focus:ring-green-500 focus:ring-offset-white'
                      }`}
                  />
                  <div>
                    <h5 className={`font-medium ${darkTheme ? 'text-white' : 'text-gray-900'}`}>
                      Secret-regex-name
                    </h5>
                    <p className={`text-sm ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`}>
                      Unpublished
                    </p>
                  </div>
                </div>
                
                <div className="flex items-center space-x-2">
                  <button 
                    type="button"
                    className={`p-2 rounded-md transition-colors duration-150 hover:bg-opacity-80
                      ${darkTheme 
                        ? 'bg-gray-700 hover:bg-gray-600' 
                        : 'bg-gray-100 hover:bg-gray-200'
                      }`}
                  >
                    <img src={Edit} alt="Edit" className="h-4 w-4" />
                  </button>
                  <button 
                    type="button"
                    className={`p-2 rounded-md transition-colors duration-150 hover:bg-opacity-80
                      ${darkTheme 
                        ? 'bg-red-900 hover:bg-red-800' 
                        : 'bg-red-100 hover:bg-red-200'
                      }`}
                  >
                    <img src={deleteicon} alt="Delete" className="h-4 w-4" />
                  </button>
                </div>
              </div> */}

            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Secretscan;