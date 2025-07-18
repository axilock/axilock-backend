import React, { useEffect, useState } from "react";
import { FaSpinner, FaUser, FaCheckCircle, FaClock, FaSearch } from "react-icons/fa";
import { getOrgUsers } from "../services/userService";
import { getErrorMessageFromResponse, formatNumber } from "../utils/helpers";

const Users = ({ darkTheme }) => {

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [users, setUsers] = useState([]);
  const [activeTab, setActiveTab] = useState('onboarded'); // 'onboarded' or 'pending'
  const [searchTerm, setSearchTerm] = useState('');

  const fetchOrgUsers = async () => {
    setLoading(true);
    try{
      const users = await getOrgUsers();
      setUsers(users.data);
    } catch(err) {
      console.error("Home.js:: ", err);
      setError(getErrorMessageFromResponse(err));

    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchOrgUsers();
  }, []);

  // Filter and search users
  const filteredUsers = users
    .filter(user => activeTab === 'onboarded' ? user.onboarded : !user.onboarded)
    .filter(user => user.username.toLowerCase().includes(searchTerm.toLowerCase()));

  // Calculate stats
  const totalUsers = users.length;
  const onboardedUsers = users.filter(user => user.onboarded).length;
  const pendingUsers = totalUsers - onboardedUsers;

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
      <div 
      style={{ "overflow": "hidden"}}x
      className={` 
        ${darkTheme ? 'bg-gray-800 text-white' : 'bg-gray-50 text-gray-800'}
      `}>
        {/* Header Section */}
        <div className={`p-6 border-b ${darkTheme ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'}`}>
          <div className="mb-6">
            <h1 className={`text-3xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
              Organization Users
            </h1>
            <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'} text-sm`}>
              Manage and view all users in your organization
            </p>
          </div>

          {/* Stats Overview */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
            <div className={`p-4 rounded-lg ${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
              <div className="flex items-center">
                <FaUser className={`text-2xl mr-3 ${darkTheme ? 'text-blue-400' : 'text-blue-500'}`} />
                <div>
                  <p className={`text-2xl font-bold ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                    {formatNumber(totalUsers)}
                  </p>
                  <p className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                    Total Users
                  </p>
                </div>
              </div>
            </div>

            <div className={`p-4 rounded-lg ${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
              <div className="flex items-center">
                <FaCheckCircle className={`text-2xl mr-3 ${darkTheme ? 'text-green-400' : 'text-green-500'}`} />
                <div>
                  <p className={`text-2xl font-bold ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                    {formatNumber(onboardedUsers)}
                  </p>
                  <p className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                    Onboarded
                  </p>
                </div>
              </div>
            </div>

            <div className={`p-4 rounded-lg ${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
              <div className="flex items-center">
                <FaClock className={`text-2xl mr-3 ${darkTheme ? 'text-yellow-400' : 'text-yellow-500'}`} />
                <div>
                  <p className={`text-2xl font-bold ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                    {formatNumber(pendingUsers)}
                  </p>
                  <p className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                    Pending
                  </p>
                </div>
              </div>
            </div>
          </div>

          {/* Tab Navigation */}
          <div className="flex space-x-1">
            <button
              onClick={() => setActiveTab('onboarded')}
              className={`px-6 py-2 rounded-lg font-medium transition-all duration-200 ${
                activeTab === 'onboarded'
                  ? (darkTheme 
                      ? 'bg-green-600 text-white' 
                      : 'bg-green-100 text-green-800 border border-green-200')
                  : (darkTheme 
                      ? 'text-gray-300 hover:bg-gray-700' 
                      : 'text-gray-600 hover:bg-gray-100')
              }`}
            >
              <FaCheckCircle className="inline mr-2" />
              Onboarded ({onboardedUsers})
            </button>
            <button
              onClick={() => setActiveTab('pending')}
              className={`px-6 py-2 rounded-lg font-medium transition-all duration-200 ${
                activeTab === 'pending'
                  ? (darkTheme 
                      ? 'bg-yellow-600 text-white' 
                      : 'bg-yellow-100 text-yellow-800 border border-yellow-200')
                  : (darkTheme 
                      ? 'text-gray-300 hover:bg-gray-700' 
                      : 'text-gray-600 hover:bg-gray-100')
              }`}
            >
              <FaClock className="inline mr-2" />
              Pending ({pendingUsers})
            </button>
          </div>
        </div>

        {/* Content Section */}
        <div className="p-6 h-[calc(100vh-460px)] flex flex-col">
          {/* Search Bar */}
          <div className="mb-6 flex-shrink-0">
            <div className="relative max-w-md">
              <FaSearch className={`absolute left-3 top-1/2 transform -translate-y-1/2 ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`} />
              <input
                type="text"
                placeholder="Search users..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className={`w-full pl-10 pr-4 py-2 rounded-lg border focus:outline-none focus:ring-2 ${
                  darkTheme 
                    ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400 focus:ring-blue-500' 
                    : 'bg-white border-gray-300 text-gray-800 placeholder-gray-500 focus:ring-blue-500'
                }`}
              />
            </div>
          </div>

          {/* Users Table */}
          <div className={`rounded-lg border overflow-hidden flex-1 flex flex-col ${darkTheme ? 'border-gray-700' : 'border-gray-200'}`}>
            <div className={`overflow-auto flex-1 ${darkTheme ? 'bg-gray-700' : 'bg-white'}`}>
              <table className="w-full">
                <thead className={`${darkTheme ? 'bg-gray-800' : 'bg-gray-50'}`}>
                  <tr>
                    <th className={`px-6 py-4 text-left text-sm font-semibold ${darkTheme ? 'text-gray-200' : 'text-gray-700'}`}>
                      User
                    </th>
                    <th className={`px-6 py-4 text-left text-sm font-semibold ${darkTheme ? 'text-gray-200' : 'text-gray-700'}`}>
                      Status
                    </th>
                    <th className={`px-6 py-4 text-left text-sm font-semibold ${darkTheme ? 'text-gray-200' : 'text-gray-700'}`}>
                      Username
                    </th>
                    <th className={`px-6 py-4 text-right text-sm font-semibold ${darkTheme ? 'text-gray-200' : 'text-gray-700'}`}>
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200 dark:divide-gray-600">
                  {filteredUsers.length === 0 ? (
                    <tr>
                      <td colSpan="4" className={`px-6 py-12 text-center ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`}>
                        <FaUser className="text-4xl mx-auto mb-4 opacity-50" />
                        <p className="text-lg mb-2">
                          {searchTerm 
                            ? `No ${activeTab} users found matching "${searchTerm}"` 
                            : `No ${activeTab} users found`
                          }
                        </p>
                        <p className="text-sm">
                          {activeTab === 'onboarded' 
                            ? 'Users will appear here once they complete onboarding.' 
                            : 'Users will appear here when they join but haven\'t completed onboarding yet.'
                          }
                        </p>
                      </td>
                    </tr>
                  ) : (
                    filteredUsers.map((user, index) => (
                      <tr key={index} className={`hover:${darkTheme ? 'bg-gray-600' : 'bg-gray-50'} transition-colors duration-150`}>
                        <td className="px-6 py-4">
                          <div className="flex items-center">
                            <div className={`w-10 h-10 rounded-full flex items-center justify-center mr-3
                              ${user.onboarded 
                                ? (darkTheme ? 'bg-green-600' : 'bg-green-100') 
                                : (darkTheme ? 'bg-yellow-600' : 'bg-yellow-100')
                              }`}>
                              <FaUser className={`text-sm
                                ${user.onboarded 
                                  ? (darkTheme ? 'text-white' : 'text-green-600') 
                                  : (darkTheme ? 'text-white' : 'text-yellow-600')
                                }`} 
                              />
                            </div>
                            <div>
                              <p className={`font-medium ${darkTheme ? 'text-white' : 'text-gray-900'}`}>
                                {user.username}
                              </p>
                            </div>
                          </div>
                        </td>
                        <td className="px-6 py-4">
                          <div className="flex items-center">
                            {user.onboarded ? (
                              <>
                                <FaCheckCircle className={`text-sm mr-2 ${darkTheme ? 'text-green-400' : 'text-green-500'}`} />
                                <span className={`text-sm font-medium px-2 py-1 rounded-full
                                  ${darkTheme ? 'bg-green-800 text-green-200' : 'bg-green-100 text-green-800'}`}>
                                  Onboarded
                                </span>
                              </>
                            ) : (
                              <>
                                <FaClock className={`text-sm mr-2 ${darkTheme ? 'text-yellow-400' : 'text-yellow-500'}`} />
                                <span className={`text-sm font-medium px-2 py-1 rounded-full
                                  ${darkTheme ? 'bg-yellow-800 text-yellow-200' : 'bg-yellow-100 text-yellow-800'}`}>
                                  Pending
                                </span>
                              </>
                            )}
                          </div>
                        </td>
                        <td className={`px-6 py-4 text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                          {user.username}
                        </td>
                        <td className="px-6 py-4 text-right">
                          <button className={`text-sm px-3 py-1 rounded hover:underline
                            ${darkTheme ? 'text-blue-400 hover:text-blue-300' : 'text-blue-600 hover:text-blue-800'}`}>
                            View Details
                          </button>
                        </td>
                      </tr>
                    ))
                  )}
                </tbody>
              </table>
            </div>
          </div>

          {/* Table Footer */}
          {/* {filteredUsers.length > 0 && (
            <div className={`mt-4 text-sm flex-shrink-0 ${darkTheme ? 'text-gray-400' : 'text-gray-600'}`}>
              Showing {filteredUsers.length} of {activeTab === 'onboarded' ? onboardedUsers : pendingUsers} {activeTab} users
            </div>
          )} */}
        </div>
      </div>
    </>
  );
};

export default Users;