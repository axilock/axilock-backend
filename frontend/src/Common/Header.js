import { Link, useLocation, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { IoSearch } from "react-icons/io5";
import { FaPlus, FaCaretDown } from "react-icons/fa";
import { ReactComponent as LOGOSVG } from '../Images/logo.svg';
import { FaUserCircle } from "react-icons/fa";
import { FaMoon, FaSun } from "react-icons/fa";
import { MdCancel } from "react-icons/md";
import { PiSignOut } from "react-icons/pi";
import { RiHome4Line } from "react-icons/ri";
import { VscRepo } from "react-icons/vsc";
import { FaUser } from "react-icons/fa";
import { IoSettingsOutline } from "react-icons/io5";

const utils = require("../utils/helpers");


const Header = ({ toggleTheme, darkTheme, toggleSidebar, currentTopPage, config }) => {

  const navigate = useNavigate();
  const location = useLocation();

  const [profileToggle, setProfileToggle] = useState(false);
  const [addDropdownOpen, setAddDropdownOpen] = useState(false); // ADD Repo button
  const [searchText, setSearchText] = useState('');

  const handleLogout = () => {
    // TODO : make api call to remove the token from backend
    localStorage.setItem("auth-token", "");
    navigate("/login", { state: { successMsg: "Logged out successfully" } })
  }

  // Set icon colors based on theme and active state
  const getIconColor = (isActive) => {
    if (isActive) {
      return darkTheme ? "text-green-400" : "text-blue-600";
    }
    return darkTheme ? "text-gray-300" : "text-gray-800";
  };

  return (
    <>
      {/* Semi-transparent Overlay when pressed on profile btn  */}
      {(profileToggle || addDropdownOpen) && (
        <div
          className="fixed top-0 left-0 w-full h-full bg-black bg-opacity-70 z-50"
          onClick={() => {
            setProfileToggle(false);
            setAddDropdownOpen(false);
          }}
        ></div>
      )}

      {/* Top Navigation Bar */}
      <div className={`flex items-center justify-between px-4 py-2 ${darkTheme ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'}`}>
        <div className="flex items-center">
          {/* Toggle Button */}
          <button
            onClick={toggleSidebar}
            className={`flex items-center justify-center w-8 h-8 p-0 mr-4 mt-[0.375rem] ${darkTheme
              ? 'text-white bg-gray-700 border-gray-600 hover:bg-gray-600'
              : 'text-gray-600 bg-transparent border-gray-300 hover:bg-gray-100'
              } border rounded`}
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill={darkTheme ? "white" : "black"}>
              <rect x="3" y="6" width="18" height="1.5" rx="1" />
              <rect x="3" y="12" width="18" height="1.5" rx="1" />
              <rect x="3" y="18" width="18" height="1.5" rx="1" />
            </svg>
          </button>

          {/* Logo and Organization Name */}
          <div className="relative flex items-center">
            <div className="w-10 h-8">
              <LOGOSVG
                viewBox="20 6 90 100"
                className={`h-full ${darkTheme ? 'logo-dark' : 'logo-light'}`}
                style={{
                  filter: darkTheme ? 'invert(100%) brightness(200%)' : 'none'
                }}
              />
            </div>
            <span className={`ml-2 mt-1 text-base font-semibold ${darkTheme ? 'text-white' : 'text-gray-700'}`}>
              {utils.capitalize(config?.orgstats?.org_name)}
            </span>
          </div>
        </div>

        {/* Search Bar */}
        <div className="flex-1 flex justify-end">
          <div className={`flex items-center max-w-xs px-3 py-2 mr-7 mt-2 mx-4 rounded-md ${darkTheme ? 'bg-gray-700' : 'bg-gray-100'
            }`}>
            <IoSearch className={`${darkTheme ? 'text-gray-300' : 'text-gray-500'} mr-2`} />
            <input
              type="text"
              placeholder="Type to Search"
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              className={`w-full bg-transparent border-none outline-none text-sm ${darkTheme ? 'text-white placeholder-gray-400' : 'text-gray-800 placeholder-gray-500'
                }`}
            />
          </div>
        </div>

        {/* Action Buttons */}
        <div className="flex items-center mt-1">
          {/* Theme Toggle Button */}
          <button
            onClick={toggleTheme}
            className={`flex items-center justify-center w-8 h-8 mr-3 rounded-full ${darkTheme
              ? 'bg-gray-700 text-yellow-300 hover:bg-gray-600'
              : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            aria-label="Toggle theme"
          >
            {darkTheme ? <FaSun size={14} /> : <FaMoon size={14} />}
          </button>

          {/* Add Button with Dropdown */}
          <div className="relative mr-4">
            {/* <div className="flex">
              <button
                onClick={(e) => {
                  e.stopPropagation();
                  setAddDropdownOpen(!addDropdownOpen);
                }}
                className={`flex items-center justify-center h-8 px-2 rounded-l-md ${darkTheme
                  ? 'bg-green-500 hover:bg-green-600 text-white border-0'
                  : 'bg-blue-500 hover:bg-blue-600 text-white border-0'
                  }`}
              >
                <FaPlus className="text-white" />
              </button>
              <button
                onClick={(e) => {
                  e.stopPropagation();
                  setAddDropdownOpen(!addDropdownOpen);
                }}
                className={`flex items-center justify-center h-8 px-1 rounded-r-md ${darkTheme
                  ? 'bg-green-500 hover:bg-green-600 text-white border-0'
                  : 'bg-blue-500 hover:bg-blue-600 text-white border-0'
                  }`}
              >
                <FaCaretDown className="text-white" />
              </button>
            </div> */}

            {/* Dropdown Menu */}
            {/* {addDropdownOpen && (
              <div className={`absolute top-10 right-0 w-40 rounded-md shadow-lg z-50 ${darkTheme ? 'bg-gray-800' : 'bg-white'
                }`}>
                <div
                  className={`py-2 px-4 text-sm cursor-pointer border-b ${darkTheme
                    ? 'text-white hover:bg-gray-700 border-gray-700'
                    : 'text-gray-700 hover:bg-gray-100 border-gray-200'
                    }`}
                  onClick={() => setAddDropdownOpen(false)}
                >
                  Add repo
                </div>
                <div
                  className={`py-2 px-4 text-sm cursor-pointer ${darkTheme
                    ? 'text-white hover:bg-gray-700'
                    : 'text-gray-700 hover:bg-gray-100'
                    }`}
                  onClick={() => setAddDropdownOpen(false)}
                >
                  Import repo
                </div>
              </div>
            )} */}
          </div>

          {/* User Profile */}
          <div className="relative">
            <FaUserCircle
              className={`w-8 h-8 cursor-pointer ${darkTheme
                ? 'text-gray-300 hover:text-white'
                : 'text-gray-800 hover:text-gray-700'
                }`}
              onClick={() => setProfileToggle(true)}
            />

            {/* Profile Dropdown (You can add this later) */}
          </div>
        </div>
      </div>

      {/* Top 2 Navigation */}
      <div className={`border-b ${darkTheme ? 'border-gray-700 bg-gray-800' : 'border-gray-200 bg-white'}`}>
        <div className={darkTheme ? 'bg-gray-800' : 'bg-white'}>
          <div className="flex px-4">
            <Link
              to="/insights"
              className={`flex items-center px-4 py-3 text-sm font-medium border-b-2 ${currentTopPage === "Home"
                ? darkTheme
                  ? 'border-green-400 text-green-400'
                  : 'border-blue-500 text-blue-600'
                : darkTheme
                  ? 'border-transparent text-gray-300 hover:text-white hover:border-gray-600'
                  : 'border-transparent text-gray-800 hover:text-gray-700 hover:border-gray-300'
                }`}
            >
              <RiHome4Line className={`w-4 h-4 mr-2 ${getIconColor(currentTopPage === "Home")}`} />
              Home
            </Link>

            {/* <Link
              to="/repositories"
              className={`flex items-center px-4 py-3 text-sm font-medium border-b-2 ${currentTopPage === "Repositories"
                ? darkTheme
                  ? 'border-green-400 text-green-400'
                  : 'border-blue-500 text-blue-600'
                : darkTheme
                  ? 'border-transparent text-gray-300 hover:text-white hover:border-gray-600'
                  : 'border-transparent text-gray-800 hover:text-gray-700 hover:border-gray-300'
                }`}
            >
              <VscRepo className={`w-4 h-4 mr-2 ${getIconColor(currentTopPage === "Repositories")}`} />
              Repositories
              <span className={`ml-1 px-2 py-0.5 text-xs font-medium rounded-full ${darkTheme
                ? 'bg-gray-700 text-gray-300'
                : 'bg-gray-100 text-gray-600'
                }`}>
                {utils.formatNumber(config?.orgstats?.repocount || 0)}
              </span>
            </Link> */}

            {/* <Link
              to="/users"
              className={`flex items-center px-4 py-3 text-sm font-medium border-b-2 ${currentTopPage === "Users"
                ? darkTheme
                  ? 'border-green-400 text-green-400'
                  : 'border-blue-500 text-blue-600'
                : darkTheme
                  ? 'border-transparent text-gray-300 hover:text-white hover:border-gray-600'
                  : 'border-transparent text-gray-800 hover:text-gray-700 hover:border-gray-300'
                }`}
            >
              <FaUser className={`w-4 h-4 mr-2 ${getIconColor(currentTopPage === "Users")}`} />
              Users
            </Link> */}

            <Link
              to="/org-setting/integrations"
              className={`flex items-center px-4 py-3 text-sm font-medium border-b-2 ${currentTopPage === "OrgSetting"
                ? darkTheme
                  ? 'border-green-400 text-green-400'
                  : 'border-blue-500 text-blue-600'
                : darkTheme
                  ? 'border-transparent text-gray-300 hover:text-white hover:border-gray-600'
                  : 'border-transparent text-gray-800 hover:text-gray-700 hover:border-gray-300'
                }`}
            >
              <IoSettingsOutline className={`w-4 h-4 mr-2 ${getIconColor(currentTopPage === "OrgSetting")}`} />
              Settings
            </Link>
          </div>
        </div>
      </div>

      {/* Profile Dropdown */}
      <div
        className={`fixed top-0 right-0 w-72 h-full shadow-lg transition-transform duration-300 ease-in-out transform ${profileToggle ? 'translate-x-0' : 'translate-x-full'} z-50 ${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}`}
      >
        <div className="p-4 border-b flex items-center justify-between">
          <div className="flex items-center">
            <FaUserCircle className="w-8 h-8 mr-2" />
            <span className="font-medium">{utils.capitalize(config?.username || "")}</span>
          </div>
          <button
            type="button"
            className="text-xl focus:outline-none"
            aria-label="Close"
            onClick={() => setProfileToggle(false)}
          >
            <MdCancel />
          </button>
        </div>
        <div className="p-4">
          <div
            className={`flex items-center py-2 px-3 rounded cursor-pointer ${darkTheme ? 'hover:bg-gray-700' : 'hover:bg-gray-100'}`}
            onClick={handleLogout}
          >
            <PiSignOut className="mr-2" />
            <span>Logout</span>
          </div>
        </div>
      </div>

    </>
  );
};

export default Header;