import React from "react";
import "./Auth.css"; // Import the styles
import { ReactComponent as FaGitlab } from "../../Images/logo-gitlab.svg";
import { FaGithub } from "react-icons/fa";
import { Link, useNavigate } from "react-router-dom";
import { ReactComponent as LOGO } from '../../Images/logo.svg';
import Globalconfig from "../../utils/config";
import { Nav } from "react-bootstrap";
import { useState, useEffect } from "react";

const AuthScreenClient = () => {
  const navigate = useNavigate(); // Initialize useNavigate

  const [config, setConfig] = useState(); // To store API response
  const [error, setError] = useState(""); // Error state
  const [clientToken, setClientToken] = useState(""); // Client token state

  const handleLoginGithub = async () => {

    sessionStorage.setItem("clientToken", clientToken);
    const url = `${Globalconfig.github.url}?client_id=${encodeURIComponent(Globalconfig.github.client_id)}&redirect_uri=${encodeURIComponent(Globalconfig.github.redirect_uri_client)}&scope=${Globalconfig.github.scope}&state=${clientToken}`;
    window.location.href = url;

  }

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const clitoken = params.get("clitoken");
    if (!clitoken) {
      setError("Invalid Client token");
    }
    setClientToken(clitoken);
  }, []);

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
              <FaGithub className="auth-btn-logo" style={{color: "black"}} />
              Login with GitHub
            </button>

            <button className="login-btn gitlab-btn" onClick={() => { }}>
              <FaGitlab className="auth-btn-logo"/>
              Login with GitLab
            </button>

            {error && (
              <div className="error-message">
                {Array.isArray(error)
                  ? error.map((err, index) => <div key={index}>• {err}</div>)
                  : error}
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
export default AuthScreenClient;
