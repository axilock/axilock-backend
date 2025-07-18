import { Link, useLocation } from "react-router-dom";
import { FaChevronDown } from "react-icons/fa";
import { RiDashboardLine } from "react-icons/ri";
import { BiShield } from "react-icons/bi";
import { IoShieldCheckmarkOutline } from "react-icons/io5";
import { BsKey } from "react-icons/bs";
import { formatNumber } from "../../utils/helpers";
import { useState } from "react";

const HomeNav = ({  onClickItem, config, darkTheme }) => {
  const location = useLocation();
  const [isOpen, setIsOpen] = useState(true);

  const toggleDropdown = () => setIsOpen(!isOpen);
  let pathwithhash = location.pathname + window.location.hash;

  // Helper function to determine icon color based on active state and theme
  const getIconColor = (isActive) => {
    if (isActive) {
      return darkTheme ? "text-green-400" : "text-blue-500";
    }
    return darkTheme ? "text-gray-300" : "text-gray-700";
  };

  return (
    <div className={`${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}`}>
      <div className="p-4">
        <h2 className={`text-xl font-semibold ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Security</h2>
      </div>
      
      <div className="px-4">
        <nav className="mb-6">
          <ul className="space-y-1">
            <li>
              <Link
                to="/insights"
                className={`flex items-center px-3 py-2 rounded-md ${
                  pathwithhash === "/insights" 
                    ? darkTheme
                      ? "bg-gray-700 text-green-400"
                      : "bg-gray-100 text-blue-500"
                    : darkTheme 
                      ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                      : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <RiDashboardLine 
                  className={`w-4 h-4 mr-2 ${getIconColor(pathwithhash === "/insights")}`} 
                />
                <span>Overview</span>
              </Link>
            </li>
            
            {/* <li>
              <Link
                to="#" 
                className={`flex items-center px-3 py-2 rounded-md ${
                  darkTheme 
                    ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                    : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <BiShield 
                  className={`w-4 h-4 mr-2 ${getIconColor(false)}`} 
                />
                <span>Risk</span>
              </Link>
            </li> */}
            
            <li>
              <Link
                to="/insights/coverage"
                className={`flex items-center px-3 py-2 rounded-md ${
                  pathwithhash === "/insights/coverage" 
                    ? darkTheme
                      ? "bg-gray-700 text-green-400"
                      : "bg-gray-100 text-blue-500"
                    : darkTheme 
                      ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                      : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <IoShieldCheckmarkOutline 
                  className={`w-4 h-4 mr-2 ${getIconColor(pathwithhash === "/insights/coverage")}`} 
                />
                <span>Coverage</span>
              </Link>
            </li>
          </ul>
        </nav>
        
        {/* Secrets Scanning Dropdown */}
        <div className="relative mb-4">
          <div
            className={`flex items-center justify-between px-3 py-2 rounded-md cursor-pointer ${
              darkTheme 
                ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                : "text-gray-700 hover:bg-gray-100"
            }`}
            onClick={toggleDropdown}
          >
            <div className="flex items-center">
              <BsKey 
                className={`w-4 h-4 mr-2 ${darkTheme ? "text-gray-300" : "text-gray-700"}`} 
              />
              <span className="text-sm font-medium">Alerts</span>
            </div>
            <FaChevronDown 
              className={`transition-transform duration-300 ${
                isOpen ? 'transform rotate-0' : 'transform rotate-90'
              } ${darkTheme ? "text-gray-300" : "text-gray-700"}`} 
            />
          </div>
          
          {isOpen && (
            <ul className={`mt-1 space-y-1 ${
              darkTheme ? 'bg-gray-800' : 'bg-white'
            }`}>
              <li>
                <Link
                  to="/insights/alerts/secret-scanning#default"
                  className={`flex items-center justify-between px-3 py-2 pl-7 rounded-md ${
                    pathwithhash === "/insights/alerts/secret-scanning#default" 
                      ? darkTheme
                        ? "bg-gray-700 text-green-400"
                        : "bg-gray-100 text-blue-500"
                      : darkTheme 
                        ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                        : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                  }`}
                  onClick={() => {
                    if (onClickItem) { onClickItem("default"); }
                  }}
                >
                  <span>Secret Alerts</span>
                  <span className={`flex items-center justify-center w-6 h-6 text-xs font-medium rounded-full ${
                    darkTheme ? 'bg-gray-600 text-white' : 'bg-gray-200 text-gray-700'
                  }`}>
                    {formatNumber(config?.orgstats?.def_pattern_count || 0)}
                  </span>
                </Link>
              </li>
            </ul>
          )}
        </div>
      </div>
    </div>
  );
};

export default HomeNav;