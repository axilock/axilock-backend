import React, { useEffect } from "react";
import "./Coverage.css"; // Import the styles
import { useState } from "react";
import Scanicon from "../Images/scan.svg";
import { coverage } from "../services/metricService";
import { FaSpinner } from "react-icons/fa";
import { getErrorMessageFromResponse } from "../utils/helpers"
import { useNavigate } from "react-router-dom";
import { mockCoverageData } from "../utils/mockData"; // Import mock data

const Coverage = ({ darkTheme, config }) => {

  const navigate = useNavigate();

  const [commitsCoverageData, setCommitsCoverageData] = useState(null); // To store API response
  const [loading, setLoading] = useState(true); // Loading state
  const [error, setError] = useState(null); // Error state

  const fetchCoverage = async () => {
    try {
      if (config?.orgstats?.is_github_onboarded) {
        // Only call the real API if GitHub is onboarded
        let response = await coverage();
        response = response["data"];
        
        if( response["total_commits"] !== 0 ) {
          response["scanned_count"] = response["total_commits"] - response["not_scanned_count"]
          response["percent_protected"] = ((response["scanned_count"] / response["total_commits"]) * 100)
          response["percent_not_protected"] = (100 - (response["percent_protected"])).toFixed(2)
        } else {
          response["scanned_count"] = 0.0
          response["percent_protected"] = 100.00
          response["percent_not_protected"] = 0.0
        }
        
        setCommitsCoverageData(response);
      } else {
        // Use mock data if GitHub is not onboarded
        setCommitsCoverageData(mockCoverageData);
      }
    } catch (err) {
      if (err.status === 401) {
        navigate("/login", { state: { errorMsg: "Please log in." } })
      } else {
        console.error("Coverage.js:: ", err);
        setError(getErrorMessageFromResponse(err));
      }
    } finally {
      setLoading(false)
    }
  };
  useEffect(() => {
    fetchCoverage();
  }, [config?.orgstats?.is_github_onboarded]); // Re-fetch when onboarding status changes

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

  console.log((commitsCoverageData?.percent_protected).toFixed(2))
  console.log((commitsCoverageData))

  return (
    <>

      <div className={`min-h-screen 
        ${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}
      `}>
        {/* Banner Section */}
        <div className={`p-6 ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
          <div className="mb-4">
            <h1 className={`text-xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Security Coverage</h1>
            <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'} text-sm`}>Security tool adoption across your organization.</p>
          </div>
          
          {config?.orgstats?.is_github_onboarded === false && (
            <div className="bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700 p-4 mb-4 rounded">
              <p className="font-bold">Integrate with org to unlock all features</p>
              <p className="text-sm">Complete GitHub onboarding to access all data points</p>
            </div>
          )}
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 p-6">
          {/* <div className={`border rounded-lg p-5 ${darkTheme ? 'bg-gray-700 border-gray-600' : 'bg-white border-gray-200'}`}>
            <div className="mb-3">
              <h4 className={`text-lg font-medium ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Repositories</h4>
            </div>
            <div className="mb-4">
              <p className={`${darkTheme ? 'text-gray-200' : 'text-gray-700'}`}>
                <span className="font-bold text-xl">{data?.repositories?.percent_enabled}%</span> of
                repositories have <span className="font-bold">secret</span> enabled
              </p>
            </div>
            <div className="mt-5 mb-4 h-1.5 bg-gray-200">
              <ProgressBar variant="success" now={data?.repositories?.percent_enabled || 0} />
            </div>
            <div className="flex justify-between text-sm">
              <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                <span className="font-medium">{data?.repositories?.enabled}</span> enabled
              </p>
              <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                <span className="font-medium">{data?.repositories?.not_enabled}</span> not enabled
              </p>
            </div>
          </div> */}

          <div className={`border rounded-lg p-5 ${darkTheme ? 'bg-gray-700 border-gray-600' : 'bg-white border-gray-200'}`}>
            <div className="mb-3 flex items-center gap-2">
              <h4 className={`text-lg font-medium ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                Secret Scanning
              </h4>
              <img src={Scanicon} alt="icon" className="h-5 w-5" />
            </div>
            <div className="relative">
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
              <div className={`mb-4 ${config?.orgstats?.is_github_onboarded === false ? `blur-with-shine ${!darkTheme ? 'light-theme-blur' : ''}` : ''}`}>
                <p className={`${darkTheme ? 'text-gray-200' : 'text-gray-700'}`}>
                  <span className="font-bold text-xl">{(commitsCoverageData?.percent_protected).toFixed(2)}%</span> of
                  your commits are protected
                </p>
              </div>

              {/* Custom Progress Bar */}
              <div className={`mt-5 mb-4 ${config?.orgstats?.is_github_onboarded === false ? `blur-with-shine ${!darkTheme ? 'light-theme-blur' : ''}` : ''}`}>
                <div className={`w-full h-1.5 rounded-full ${darkTheme ? 'bg-gray-600' : 'bg-gray-200'}`}>
                  <div
                    className="h-1.5 bg-green-500 rounded-full transition-all duration-300 ease-in-out"
                    style={{ width: `${commitsCoverageData?.percent_protected || 0}%` }}
                  ></div>
                </div>
              </div>

              <div className={`flex justify-between text-sm ${config?.orgstats?.is_github_onboarded === false ? `blur-with-shine ${!darkTheme ? 'light-theme-blur' : ''}` : ''}`}>
                <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                  <span className="font-medium">{commitsCoverageData.scanned_count}</span> Protected Commits
                </p>
                <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                  <span className="font-medium">{commitsCoverageData.not_scanned_count}</span> Needs Attention
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>


    </>
  );
};

export default Coverage;
