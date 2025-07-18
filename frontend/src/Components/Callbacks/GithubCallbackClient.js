import React, { useEffect, useState } from "react";
import Header from "../../Common/Header"
import { FaSpinner } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { processGithubCallbackAuthClient } from "../../services/configService";
import { getErrorMessageFromResponse } from "../../utils/helpers"
import Globalconfig from "../../utils/config";

const GithubCallbackClient = () => {

  const [loading, setLoading] = useState(true); // Loading state
  const [error, setError] = useState(null); // Error state
  const [clientToken, setClientToken] = useState(""); // Client token state
  const [success, setSuccess] = useState(false); // Success state
  const navigate = useNavigate();

  const handleGithubAuth = async () => {
    setLoading(true); // Set loading to true before fetching data

    const params = new URLSearchParams(window.location.search);
    const paramsObj = Object.create(null);

    for (const [key, value] of params.entries()) {
      paramsObj[key] = value;
    }

    if(paramsObj.state !== sessionStorage.getItem("clientToken")) {
      setError("Invalid state parameter");
      setLoading(false);
      return;
    }
    paramsObj["clitoken"] = sessionStorage.getItem("clientToken");
    paramsObj["provider"] =  Globalconfig.provider.github;
    

    try {
      let response = await processGithubCallbackAuthClient(paramsObj);
      // Set success to true after successful authentication
      setSuccess(true);
      return;
    
    } catch(err) {
      if (err.status === 401) {
        navigate("/login", { state: { errorMsg: "Please log in." } })
      } else {
        console.error("GithubCallbackClient.js:: ",err);
        setError(getErrorMessageFromResponse(err));
      }
    } finally {
      setLoading(false);
    }
  }


  useEffect(() => {
    let token = sessionStorage.getItem("clientToken")
    setClientToken(token);
    
    if(!token) {
      setError("Invalid Client token");
      setLoading(false);
      return;
    }
    
    handleGithubAuth();
  }, []) // Added clientToken as a dependency

  if (loading) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white">
        <div className="animate-spin text-6xl text-green-500 mb-8">
          {/* Spinner icon */}
          <svg className="w-16 h-16" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
        </div>
        <h1 className="text-3xl font-bold">Authenticating with GitHub...</h1>
        <p className="mt-4 text-gray-300">Just a moment, please.</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white">
        <div className="rounded-lg bg-red-500 bg-opacity-20 p-8 max-w-md">
          <h1 className="text-3xl font-bold text-red-400 mb-4">Authentication Failed</h1>
          <p className="text-lg">{error}</p>
          <button 
            onClick={() => navigate("/login")}
            className="mt-6 bg-white text-gray-900 py-3 px-6 rounded-full font-medium hover:bg-gray-200 transition duration-300"
          >
            Return to Login
          </button>
        </div>
      </div>
    );
  }

  if (success) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-b from-gray-900 to-green-900 text-white">
        <div className="flex flex-col items-center max-w-md text-center">
          {/* Check circle icon */}
          <div className="text-green-400 text-6xl mb-6">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-16 w-16" viewBox="0 0 20 20" fill="currentColor">
              <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
            </svg>
          </div>
          <h1 className="text-4xl font-bold mb-4">Authentication Successful!</h1>
          <p className="text-xl mb-10">You've successfully connected with GitHub</p>
          <div className="w-16 h-1 bg-green-400 rounded mb-10"></div>
          <p className="text-gray-300">You can close this window and return to the application.</p>
          <button 
            onClick={() => navigate("/")}
            className="mt-8 bg-green-500 text-white py-3 px-8 rounded-full font-medium hover:bg-green-600 transition duration-300"
          >
            Continue to Dashboard
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white">
      <h1 className="text-3xl font-bold">GitHub Callback</h1>
      <p className="mt-4 text-gray-300">Processing your authentication...</p>
    </div>
  );
};

export default GithubCallbackClient;