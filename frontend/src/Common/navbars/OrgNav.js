import { Link, useLocation } from "react-router-dom";
import overviewicon from "../../Images/overview.svg";
import { FaLink } from "react-icons/fa6";

const OrgNav = ({ show, darkTheme }) => {
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
            {/* <li className="mb-1">
              <Link
                to="/org-setting"
                className={`flex items-center px-3 py-2 rounded-md ${
                  pathwithhash === "/org-setting" 
                    ? darkTheme
                      ? "bg-gray-700 text-green-400"
                      : "bg-gray-100 text-blue-600"
                    : darkTheme 
                      ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                      : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <img src={overviewicon} alt="overview icon" className="w-4 h-4 mr-2" />
                <span>Org Structure</span>
              </Link>
            </li> */}
            
            <li className="mb-1">
              <Link
                to="/org-setting/integrations"
                className={`flex items-center px-3 py-2 rounded-md ${
                  pathwithhash === "/org-setting/integrations" 
                    ? darkTheme
                      ? "bg-gray-700 text-green-400"
                      : "bg-gray-100 text-blue-600"
                    : darkTheme 
                      ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                      : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <FaLink className="w-4 h-4 mr-2" />
                <span>Integrations</span>
              </Link>
            </li>
          </ul>
        </nav>
        
        {/* <div className="mb-4">
          <h2 className={`text-lg font-medium mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Security</h2>
          
          <ul>
            <li className="mb-1">
              <Link
                to="/org-setting/custom-secrets"
                className={`flex items-center px-3 py-2 rounded-md ${
                  pathwithhash === "/org-setting/custom-secrets" 
                    ? darkTheme
                      ? "bg-gray-700 text-green-400"
                      : "bg-gray-100 text-blue-600"
                    : darkTheme 
                      ? "text-gray-300 hover:bg-gray-700 hover:text-white" 
                      : "text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                }`}
              >
                <img src={overviewicon} alt="overview icon" className="w-4 h-4 mr-2" />
                <span>Secret Scanning</span>
              </Link>
            </li>
          </ul>
        </div> */}
      </div>
    </div>
  );
};

export default OrgNav;