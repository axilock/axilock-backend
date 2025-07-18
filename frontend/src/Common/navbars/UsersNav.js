import { Link, useLocation } from "react-router-dom";
import overviewicon from "../../Images/overview.svg";
import { FaUsers } from "react-icons/fa";


const UsersNav = ({ show, darkTheme }) => {
  const location = useLocation();
  let pathwithhash = location.pathname + window.location.hash;

  return (
    <div className={`offcanvas-navbar-left ${show ? "slide-in" : "slide-out"} ${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}`}>
      <div className="p-4">
        <h2 className={`text-xl font-semibold ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Settings</h2>
      </div>
      
      <div className="px-4">
        <nav className="mb-6">
          <ul>
            <li className="mb-1">
              <Link
                to="/users"
                className={`flex items-center px-3 py-2 rounded-md ${
                  pathwithhash === "/users" 
                    ? darkTheme
                      ? "bg-gray-700 text-green-400"
                      : "bg-gray-100 text-blue-600"
                    : darkTheme 
                      ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                      : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <FaUsers className="w-5 h-5 mr-2" />
                <span>Users</span>
              </Link>
            </li>
            
          </ul>
        </nav>
      </div>
    </div>
  );
};

export default UsersNav;