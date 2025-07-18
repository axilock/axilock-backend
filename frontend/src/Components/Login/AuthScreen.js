import React from "react";
import "./Auth.css"; // Import the styles
import { ReactComponent as FaGitlab } from "../../Images/logo-gitlab.svg";
import { FaGithub } from "react-icons/fa";
import { Link, useNavigate, useLocation } from "react-router-dom";
import { ReactComponent as LOGO } from '../../Images/logo.svg';
import { useState, useEffect } from "react";
import { Nav } from "react-bootstrap";

import Globalconfig from "../../utils/config";
const helpers = require("../../utils/helpers");

const AuthScreen = () => {
  
  const location = useLocation();
  const navigate = useNavigate(); // Initialize useNavigate
  const [error, setError] = useState(location.state?.errorMsg || ""); // Error state
  const [successMsg, setSuccessMsg]  = useState(location.state?.successMsg || ""); // Success message state
  const handleLogin = async () => {
    navigate("/login");
  };

  const handleLoginGithub = async () => {

    const stateParam = helpers.generateOAuthState();
    sessionStorage.setItem("oauth_state", stateParam);
    const url = `${Globalconfig.github.url}?client_id=${encodeURIComponent(Globalconfig.github.client_id)}&redirect_uri=${encodeURIComponent(Globalconfig.github.redirect_uri)}&scope=${Globalconfig.github.scope}&state=${stateParam}`;
    window.location.href = url;

  }

  useEffect(() => {
    if (localStorage.getItem("auth-token")) {
      navigate("/insights");
    }
  }, []);

  useEffect(() => {
    if (successMsg) {
      const timer = setTimeout(() => {
        setSuccessMsg("");
      }, 3000); // 3 seconds

      return () => clearTimeout(timer); // Clean up the timer on component unmount
    }
  }, [successMsg]);


  return (
    <div className="login-page-container">
      <div className="login-content">
        <div className="logo-container">
          <a href="/" className="logo-link">
            {/* Replace with your actual logo component */}
            <Nav.Link href="/" className="logo-link">
              <LOGO style={{ width: '48px' }}></LOGO>
            </Nav.Link>
          </a>
        </div>

        <div className="login-box">
          <h2 className="login-title">Sign In to Axilock</h2>

          <div className="login-options">
            <button className="login-btn github-btn" onClick={handleLoginGithub}>
              <FaGithub className="auth-btn-logo" style={{ color: "black" }} />
              Login with GitHub
            </button>

            <button className="login-btn gitlab-btn" onClick={() => { }}>
              <FaGitlab className="auth-btn-logo" />
              Login with GitLab
            </button>

            {error && (
              <div className="error-message">
                {Array.isArray(error)
                  ? error.map((err, index) => <div key={index}>• {err}</div>)
                  : error}
              </div>
            )}
            {successMsg && (
              <div className="success-message">
                {Array.isArray(successMsg)
                  ? successMsg.map((msg, index) => <div key={index}>• {msg}</div>)
                  : successMsg}
              </div>
            )}

            <p className="policy-text">
              By logging in, you agree to Axilock's{" "}
              <a href="#" className="policy-link">Privacy Policy</a> and{" "}
              <a href="#" className="policy-link">Terms of Use</a>.
            </p>

            <div className="signup-container">
              <p className="signup-text">Don't have an account yet?</p>
              <a href="#" className="signup-link">
                Contact Us
              </a>
            </div>
          </div>
        </div>
      </div>

      <footer className="login-footer">
        <div className="footer-links">
          <a href="#" className="footer-link">Privacy Policy</a>
          <span className="footer-divider">|</span>
          <a href="#" className="footer-link">Terms of Use</a>
          <span className="footer-divider">|</span>
          <a href="#" className="footer-link">Contact</a>
          <span className="footer-divider">|</span>
          <a href="#" className="footer-link">Help</a>
        </div>
        <p className="copyright">© 2025 Axilock. All rights reserved.</p>
      </footer>
    </div>
  );
};
export default AuthScreen;
