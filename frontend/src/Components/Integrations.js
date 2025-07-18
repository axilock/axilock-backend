import React, { useEffect, useState } from "react";
import Header from "../Common/Header";
import { FaSpinner } from "react-icons/fa";
import { fetchIntegrations } from "../services/configService";
import { FaGithub } from "react-icons/fa";
import { getErrorMessageFromResponse } from "../utils/helpers";
import { useNavigate, useLocation } from "react-router-dom";
import Globalconfig from "../utils/config";
const helpers = require("../utils/helpers");


const IntegrationsScreen = ({ darkTheme = false }) => {

  const location = useLocation();

  const [integrations, setIntegrations] = useState(null);
  const [loading, setLoading] = useState(true); 
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleGithubIntegration = () => {
    
    const stateParam = helpers.generateOAuthState();
    sessionStorage.setItem("oauth_state_app", stateParam);
    window.location.href = `${Globalconfig.github.url_app_install}?state=${stateParam}&redirect_uri=${Globalconfig.github.redirect_uri}`;
    return;
  };

  const fetchData = async () => {
    setLoading(true);

    try {
      const [fetchedintegrations] = await Promise.all([fetchIntegrations()]);
      let integrationsdict = Object.create(null);

      for (let integration of fetchedintegrations.integrations) {
        integrationsdict[integration.name] = integration;
      }

      setIntegrations(integrationsdict);
    } catch (err) {
      if (err.status === 401) {
        navigate("/login", { state: { "errorMsg" : "Please log in." } });
      } else {
        console.error("Integrations.js:: ", err);
        setError(getErrorMessageFromResponse(err));
      }
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  if (loading) return (
    <div className="min-h-screen">
      <div className={`flex items-center justify-center h-screen ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
        <FaSpinner className={`animate-spin text-4xl ${darkTheme ? 'text-green-400' : 'text-green-500'}`} />
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
      <div className={`container mx-auto px-4 py-8 ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
        {/* Header section - commented out in original */}
        {/* <Header /> */}

        {/* Banner Section */}
        <div className={`mt-6 p-6 rounded-lg ${darkTheme ? 'bg-gray-700' : 'bg-gray-50'}`}>
          <div className="mb-6">
            <h1 className={`text-2xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
              Integrate Axilock with:
            </h1>
            <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
              Manage your secrets seamlessly across platforms
            </p>
          </div>

          <div className="mt-8">
            {integrations?.github ? (
              <button
                className={`flex items-center gap-2 px-4 py-2 rounded-md transition-colors
                  ${integrations.github.status !== "inactive"
                    ? `${darkTheme ? 'bg-gray-600 text-gray-300' : 'bg-gray-200 text-gray-500'} cursor-not-allowed`
                    : `${darkTheme ? 'bg-green-600 hover:bg-green-700' : 'bg-green-500 hover:bg-green-600'} text-white`
                  }`}
                onClick={handleGithubIntegration}
                disabled={integrations.github.status !== "inactive"}
              >
                <FaGithub className="text-xl" />
                <span>
                  {integrations.github.status !== "inactive"
                    ? "Connected To Github"
                    : "Connect with GitHub"}
                </span>
              </button>
            ) : null}
          </div>
        </div>
      </div>
    </div>
  );
};

export default IntegrationsScreen;