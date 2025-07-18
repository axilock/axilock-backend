import { useState, useEffect } from "react";
import { IoSearch } from "react-icons/io5";
import { FaSpinner } from "react-icons/fa";
import { RiAlertLine } from "react-icons/ri";
import { FiEdit2 } from "react-icons/fi";
import "./Customrules.css";
import "../Components/Coverage.css"; // Import the Coverage.css for the blur-with-shine effect
import { mockAllAlertsList, mockTopAlertsByType, mockTopReposWithSecrets } from "../utils/mockData"; // Import mock data
import { formatNumber, getErrorMessageFromResponse } from "../utils/helpers";
import { getTopAlertsByType, getAlertsFromState, getTopReposWithAlerts } from "../services/metricService";


const Customrules = ({ darkTheme , config }) => {
  const [alertsList, setAlertsList] = useState([]);
  const [topReposWithAlerts, setTopReposWithAlerts] = useState([]);
  const [topAlertsByType, setTopAlertsByType] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [alertType, setAlertType] = useState("default");
  const [searchQuery, setSearchQuery] = useState("");

  const fetchData = async (type = alertType) => {
    setLoading(true);
    try {
      if (config?.orgstats?.is_github_onboarded === true) {
        // Only call the real APIs if GitHub is onboarded
        const [topAlertsByTypeResponse, alertsResponse, topReposWithAlertsResponse] = await Promise.all([
          getTopAlertsByType(type), 
          getAlertsFromState("open", type),
          getTopReposWithAlerts()
        ]);

        setTopAlertsByType(topAlertsByTypeResponse.data);
        setAlertsList(alertsResponse.data);
        setTopReposWithAlerts(topReposWithAlertsResponse.data)
      } else if(config?.orgstats?.is_github_onboarded === false) {
        // Use mock data if GitHub is not onboarded
        
        setTopAlertsByType(mockTopAlertsByType);
        setAlertsList(mockAllAlertsList);
        setTopReposWithAlerts(mockTopReposWithSecrets)
      } 

      console.log(topReposWithAlerts, topAlertsByType, config)
    } catch (err) {
      console.error("CustomRules.js:: ", err);
      setError(getErrorMessageFromResponse(err));
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [config?.orgstats?.is_github_onboarded]); // Re-fetch when onboarding status changes

  const handleToggleChange = (type) => {
    setAlertType(type);
    fetchData(type);
  };

  const filteredAlerts = alertsList.filter(alert => 
    searchQuery ? alert.alert_name.toLowerCase().includes(searchQuery.toLowerCase()) : true
  );

  if (loading) return (
    <div className="min-h-screen">
      <div className={`flex items-center justify-center h-screen ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
        <FaSpinner className={`animate-spin text-4xl ${darkTheme ? 'text-green-400' : 'text-green-800'}`} />
      </div>
    </div>
  );

  if (error) return (
    <div className="min-h-screen">
      <div className={`flex items-center justify-center h-screen ${darkTheme ? 'bg-gray-900' : 'bg-gray-50'}`}>
        <p className={`text-lg ${darkTheme ? 'text-red-400' : 'text-red-500'}`}>{error}</p>
      </div>
    </div>
  );

  return (
    <div className={`min-h-screen ${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}`}>
      <div className="mb-4">
        {/* Main Banner */}
        <div className={` p-6 rounded-lg ${darkTheme ? 'bg-gray-700' : 'bg-white'}`}>
          {/* Title Section with Professional Toggle */}
          {config?.orgstats?.is_github_onboarded === false && (
              <div className="bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700 p-4 mb-4 rounded w-full md:w-auto">
                <p className="font-bold">Integrate with org to unlock all features</p>
                <p className="text-sm">Complete GitHub onboarding to access all data points</p>
              </div>
            )}
          <div className="flex flex-col md:flex-row justify-between items-center mb-6">
            <h1 className={`text-xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
              Secret scanning alerts
            </h1>   
            <div className="mt-4 md:mt-0 flex items-center">
              <div className="mr-2 text-sm font-medium">
                <span className={alertType === "default" 
                  ? `${darkTheme ? 'text-blue-400' : 'text-blue-600'}`
                  : `${darkTheme ? 'text-gray-400' : 'text-gray-500'}`
                }>
                  Default Patterns
                </span>
                <span className={`ml-1 ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`}>
                  ({formatNumber(config?.orgstats?.def_pattern_count)})
                </span>
              </div>
              
              {/* Toggle Switch */}
              <div 
                className={`relative inline-block w-12 h-6 transition duration-200 ease-in-out rounded-full cursor-pointer ${
                  darkTheme ? 'bg-gray-600' : 'bg-gray-200'
                }`}
                onClick={() => handleToggleChange(alertType === "default" ? "custom" : "default")}
              >
                <div
                  className={`absolute left-0 top-0 w-6 h-6 transition-transform duration-200 ease-in-out transform ${
                    alertType === "custom" ? "translate-x-6" : "translate-x-0"
                  } bg-white rounded-full shadow-md`}
                ></div>
              </div>
              
              <div className="ml-2 text-sm font-medium">
                <span className={alertType === "custom" 
                  ? `${darkTheme ? 'text-blue-400' : 'text-blue-600'}`
                  : `${darkTheme ? 'text-gray-400' : 'text-gray-500'}`
                }>
                  Custom Patterns
                </span>
                <span className={`ml-1 ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`}>
                  ({formatNumber(config?.orgstats?.custom_pattern_count)})
                </span>
              </div>
            </div>
          </div>

          {/* Top Alerts and Repos Section */}
          <div className="mb-8 grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Top Alert Types Table */}
            <div>
              <h2 className={`text-l font-semibold mb-4 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                Top Alert Types
              </h2>
              <div className={`rounded-lg overflow-hidden ${darkTheme ? 'bg-gray-600' : 'bg-white border border-gray-200'} relative`}>
                {config?.orgstats?.is_github_onboarded === false && (
                  <div 
                    style={{
                      position: 'absolute',
                      top: 0,
                      left: 0,
                      right: 0,
                      bottom: 0,
                      backgroundColor: 'rgba(255, 255, 255, 0.05)',
                      zIndex: 10,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      borderRadius: '4px'
                    }}
                  >
                    <div 
                      style={{
                        background: darkTheme ? 'rgba(50, 50, 50, 0.9)' : 'rgba(70, 70, 70, 0.9)',
                        color: 'white',
                        padding: '8px 16px',
                        borderRadius: '4px',
                        fontSize: '14px',
                        fontWeight: 500
                      }}
                    >
                      Complete GitHub onboarding to view
                    </div>
                  </div>
                )}
                <table className="w-full">
                  <thead className={`${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
                    <tr>
                      <th className={`px-4 py-3 text-left text-sm font-medium ${darkTheme ? 'text-gray-200' : 'text-gray-600'}`}>Type</th>
                      <th className={`px-4 py-3 text-right text-sm font-medium ${darkTheme ? 'text-gray-200' : 'text-gray-600'}`}>Count</th>
                    </tr>
                  </thead>
                  <tbody className={`divide-y divide-gray-200 ${config?.orgstats?.is_github_onboarded === false ? `blur-with-shine ${!darkTheme ? 'light-theme-blur' : ''}` : ''}`}>

                    {topAlertsByType.map((alert, index) => (
                      <tr key={index}>
                        <td className={`px-4 py-3 text-sm ${darkTheme ? 'text-white' : 'text-gray-800'}`}>{alert.type}</td>
                        <td className={`px-4 py-3 text-sm text-right font-medium ${darkTheme ? 'text-green-400' : 'text-green-600'}`}>{alert.count}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>

            {/* Top Repos with Secrets Table */}
            <div>
              <h2 className={`text-l font-semibold mb-4 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                Top Repos with Secrets
              </h2>
              <div className={`rounded-lg overflow-hidden ${darkTheme ? 'bg-gray-600' : 'bg-white border border-gray-200'} relative`}>
                {config?.orgstats?.is_github_onboarded === false && (
                  <div 
                    style={{
                      position: 'absolute',
                      top: 0,
                      left: 0,
                      right: 0,
                      bottom: 0,
                      backgroundColor: 'rgba(255, 255, 255, 0.05)',
                      zIndex: 10,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      borderRadius: '4px'
                    }}
                  >
                    <div 
                      style={{
                        background: darkTheme ? 'rgba(50, 50, 50, 0.9)' : 'rgba(70, 70, 70, 0.9)',
                        color: 'white',
                        padding: '8px 16px',
                        borderRadius: '4px',
                        fontSize: '14px',
                        fontWeight: 500
                      }}
                    >
                      Complete GitHub onboarding to view
                    </div>
                  </div>
                )}
                <table className="w-full">
                  <thead className={`${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
                    <tr>
                      <th className={`px-4 py-3 text-left text-sm font-medium ${darkTheme ? 'text-gray-200' : 'text-gray-600'}`}>Repository</th>
                      <th className={`px-4 py-3 text-right text-sm font-medium ${darkTheme ? 'text-gray-200' : 'text-gray-600'}`}>Count</th>
                    </tr>
                  </thead>
                  <tbody className={`divide-y divide-gray-200 ${config?.orgstats?.is_github_onboarded === false ? `blur-with-shine ${!darkTheme ? 'light-theme-blur' : ''}` : ''}`}>
                    {/* Top repos with secrets */}
                    {topReposWithAlerts.map((repo, index) => (
                      <tr key={index}>
                        <td className={`px-4 py-3 text-sm ${darkTheme ? 'text-white' : 'text-gray-800'}`}>{repo.repo}</td>
                        <td className={`px-4 py-3 text-sm text-right font-medium ${darkTheme ? 'text-green-400' : 'text-green-600'}`}>{repo.count}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          {/* Search Section - Simplified */}
          <div className="mb-6">
            <div className={`relative flex items-center w-full max-w-md ${darkTheme ? 'bg-gray-600' : 'bg-white'} rounded-lg overflow-hidden`}>
              <button type="button" className={`absolute left-3 ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`}>
                <IoSearch />
              </button>
              <input
                type="text"
                className={`w-full py-2 pl-10 pr-4 outline-none ${darkTheme ? 'bg-gray-600 text-white' : 'bg-white text-gray-800'} border ${darkTheme ? 'border-gray-600' : 'border-gray-300'} rounded-lg`}
                placeholder="Search alerts..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
              />
            </div>
          </div>

          {/* Recent Alerts Section with Professional Table */}
          <div className="mb-4">
            <h2 className={`text-l font-semibold mb-4 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
              {alertType === "default" ? "Default Pattern Alerts" : "Custom Pattern Alerts"}
            </h2>
          </div>

          {/* Professional Alert Table */}
          <div className={`overflow-hidden rounded-lg ${darkTheme ? 'bg-gray-600' : 'bg-white border border-gray-200'} relative`}>
            {config?.orgstats?.is_github_onboarded === false && (
              <div 
                style={{
                  position: 'absolute',
                  top: 0,
                  left: 0,
                  right: 0,
                  bottom: 0,
                  backgroundColor: 'rgba(255, 255, 255, 0.05)',
                  zIndex: 10,
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  borderRadius: '4px'
                }}
              >
                <div 
                  style={{
                    background: darkTheme ? 'rgba(50, 50, 50, 0.9)' : 'rgba(70, 70, 70, 0.9)',
                    color: 'white',
                    padding: '8px 16px',
                    borderRadius: '4px',
                    fontSize: '14px',
                    fontWeight: 500
                  }}
                >
                  Complete GitHub onboarding to view
                </div>
              </div>
            )}
            <table className="min-w-full divide-y divide-gray-200">
              <thead className={`${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
                <tr>
                  <th scope="col" className={`px-6 py-3 text-left text-xs font-medium uppercase tracking-wider ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                    Alert Type
                  </th>
                  <th scope="col" className={`px-6 py-3 text-left text-xs font-medium uppercase tracking-wider ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                    Source
                  </th>
                  <th scope="col" className={`px-6 py-3 text-left text-xs font-medium uppercase tracking-wider ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                    Repository
                  </th>
                  <th scope="col" className={`px-6 py-3 text-left text-xs font-medium uppercase tracking-wider ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                    Date
                  </th>
                  <th scope="col" className={`px-6 py-3 text-left text-xs font-medium uppercase tracking-wider ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                    Path
                  </th>
                  <th scope="col" className={`px-6 py-3 text-right text-xs font-medium uppercase tracking-wider ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                    Action
                  </th>
                </tr>
              </thead>
              <tbody className={`divide-y divide-gray-200 ${darkTheme ? 'bg-gray-600' : 'bg-white'} ${config?.orgstats?.is_github_onboarded === false ? `blur-with-shine ${!darkTheme ? 'light-theme-blur' : ''}` : ''}`}>
                {filteredAlerts.length > 0 ? (
                  filteredAlerts.map((item, index) => (
                    <tr key={index} className={darkTheme ? 'hover:bg-gray-500' : 'hover:bg-gray-50'}>
                      <td className={`px-6 py-4 whitespace-nowrap text-sm font-medium ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                        <div className="flex items-center">
                          <RiAlertLine className={`w-4 h-4 mr-2 ${darkTheme ? 'text-yellow-400' : 'text-orange-500'}`} />
                          {item.alert_name}
                        </div>
                      </td>
                      <td className={`px-6 py-4 whitespace-nowrap text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                        {item.source}
                      </td>
                      <td className={`px-6 py-4 whitespace-nowrap text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                        {item.source_repo}
                      </td>
                      <td className={`px-6 py-4 whitespace-nowrap text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                        {item.date}
                      </td>
                      <td className={`px-6 py-4 text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'} max-w-xs truncate`}>
                        {item.path}
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm">
                        <button className={`p-1 rounded transition-colors ${darkTheme ? 'hover:bg-gray-500 text-gray-300 hover:text-white' : 'hover:bg-gray-100 text-gray-600 hover:text-gray-800'}`}>
                          <FiEdit2 className="w-4 h-4" />
                        </button>
                      </td>
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan="6" className={`px-6 py-8 text-center text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-500'}`}>
                      No alerts found. Try adjusting your search.
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Customrules;
