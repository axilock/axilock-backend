import React, { useEffect, useState } from "react";
import Header from "../../Common/Header"
import { FaSpinner } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { processGithubCallbackUser } from "../../services/configService";
import { getErrorMessageFromResponse } from "../../utils/helpers"

const GithubCallback = () => {

  const [loading, setLoading] = useState(true); // Loading state
  const [error, setError] = useState(null); // Error state
  const navigate = useNavigate();

  const handleGithubAuth = async () => {
    setLoading(true); // Set loading to true before fetching data

    const params = new URLSearchParams(window.location.search);
    const paramsObj = Object.create(null);

    for (const [key, value] of params.entries()) {
      paramsObj[key] = value;
    }

    if (paramsObj.error === "access_denied") {
      setLoading(false);
      navigate("/login", { state: { errorMsg: "Access denied by user" } })
      return;
    }

    if (paramsObj.state !== sessionStorage.getItem("oauth_state")) {
      setLoading(false);
      navigate("/login", { state: { errorMsg: "Invalid state parameter" } })
      return;
    }

    try {
      let response = await processGithubCallbackUser(paramsObj);
      if (response.accesstoken) {
        // Store token in localStorage
        localStorage.setItem("auth-token", response["accesstoken"]);
        navigate("/insights", { state : {successToast: "Logged In Successfully"}});
      } else {
        console.error("Githubcallback.js:: ", response);
        setError(getErrorMessageFromResponse(response));
        navigate("/login", { state: { errorMsg: getErrorMessageFromResponse(response) } })
      }
      return

    } catch (err) {
      if (err.status === 401) {
        navigate("/login", { state: { errorMsg: "Invalid Login details" } })
      } else {
        console.error("GithubCallback.js:: ", err);
        setError(getErrorMessageFromResponse(err));
        navigate("/login", { state: { errorMsg: getErrorMessageFromResponse(err) } })
      }
    } finally {
      setLoading(false);
    }
  }


  useEffect(() => {
    handleGithubAuth();
  }, [])

  if (loading) return (
    <div className="main-body">
      <div className="main">
        <div className="home-container">
          
        </div>
      </div>
      <div className="main-banner">
        <div className="banner">
          <div className="main-title">
            <FaSpinner className="rotating-icon" />
          </div>
        </div>
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
            <p>Redirecting to login page</p>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <>
      <div className="main">
        <div className="home-container">
          {/* <Header /> */}
        </div>
      </div>
      <div className="main-banner">
        <div className="banner">
          <div className="main-title">
            <div className="banner-title">Github Callback</div>
          </div>
        </div>
      </div>
    </>
  );
};

export default GithubCallback;