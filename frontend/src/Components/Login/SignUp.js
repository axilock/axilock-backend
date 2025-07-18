import React, { useState, useEffect } from "react";
import "./Auth.css";
import { signupUser } from "../../services/userService";
import { useNavigate, Link } from "react-router-dom";
import { ReactComponent as LOGO } from '../../Images/logo.svg';
import { MdArrowBack } from "react-icons/md";
import { AiFillEye, AiFillEyeInvisible } from "react-icons/ai";
import { FaSpinner } from "react-icons/fa";
import { Nav } from "react-bootstrap";
import { getUser } from "../../services/userService";
import { validateEmail, validatePassword, getErrorMessageFromResponse } from "../../utils/helpers";



const SignUp = () => {
  const [email, setEmail] = useState("");       // Email state
  const [password, setPassword] = useState(""); // Password state
  const [confirmpassword, setConfirmpassword] = useState(""); // confirmPassword state
  const [showPassword, setShowPassword] = useState(false); // Show/Hide password toggle
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSignup = async () => {
    setLoading(true);
    setError(""); // Reset error message

    if (!validateEmail(email)) {
      setError("Invalid email format.");
      setLoading(false)
      return;
    }
    let passwordErrors = validatePassword(password);
    if (passwordErrors.length > 0) {
      setError(passwordErrors);
      setLoading(false)
      return;
    }

    if (password !== confirmpassword) {
      setError("Passwords donot match.");
      setLoading(false)
      return;
    }
    try {

      const response = await signupUser({ email, password });

      if (response) {
        navigate("/login", { state: { successMsg: "Account created successfully. Please log in." } });
      } else {
        console.error("signup.js:: ", response);
        setError(getErrorMessageFromResponse(response));
      }
    } catch (error) {
      setError(`Login Failed : ${getErrorMessageFromResponse(error)}`);
      console.error("Login.js", error);
    } finally {
      setLoading(false); // Stop loading spinner
    }
  };
  const globalSetting = async () => {
    try {
      const response = await getUser();
      navigate("/insights");
    } catch (e) {
      console.error("Authscreen.js:: ", e);
    }
  };

  useEffect(() => {
    globalSetting();
  }, []);

  return (
    <div>
      <div className="logo-container">
        <Nav.Link href="/" className="logo-link">
          <LOGO style={{ width: '48px' }}></LOGO>
        </Nav.Link>
      </div>

      <h2 className="subtitle">Create Account</h2>
      <div className="form-container" style={{ "paddingTop": "35px", "minHeight": "44vh" }}>
        <div className="login-box-content">
          <label htmlFor="email">EMAIL*</label>
          <input
            type="email"
            id="email"
            placeholder="Enter your email address"
            className="input-box"
            pattern="[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
            required
            onChange={(e) => setEmail(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter') {
                handleSignup();
              }
            }}
          />
          <label htmlFor="password">PASSWORD*</label>
          <div className="password-container">
            <input
              type={showPassword ? "text" : "password"} // Toggle input type
              id="password"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value)

                if (validatePassword(e.target.value).length > 0) { setError(validatePassword(e.target.value)); }
                else setError();
              }} // Update password state
              placeholder="Enter Password"
              required
              className="input-box"
            />
            <button
              type="button"
              className="toggle-password-btn"
              onClick={() => setShowPassword(!showPassword)} // Toggle state
            >
              {showPassword ? <AiFillEyeInvisible /> : <AiFillEye />}
            </button>
          </div>
          {/* repeat */}
          <label htmlFor="password2">REPEAT PASSWORD*</label>
          <div className="password-container">
            <input
              type={showPassword ? "text" : "password"} // Toggle input type
              id="password2"
              value={confirmpassword}
              required
              onChange={(e) => {
                setConfirmpassword(e.target.value)
                if (password !== e.target.value) { setError("Passwords donot match."); }
                else setError();
              }} // Update password state
              placeholder="Repeat Password"
              className="input-box"
              onKeyDown={(e) => {
                if (e.key === 'Enter') {
                  handleSignup();
                }
              }}
            />
            <button
              type="button"
              className="toggle-password-btn"
              onClick={() => setShowPassword(!showPassword)} // Toggle state
            >
              {showPassword ? <AiFillEyeInvisible /> : <AiFillEye />}
            </button>
          </div>
          <button
            className="continue-btn"
            onClick={handleSignup}
            disabled={loading}>
            {loading ? (
              <FaSpinner className="rotating-icon" />
            ) : (
              "Continue"
            )}
          </button>
          <p>
            {" "}
            {error && (
              <div className="error-message">
                {Array.isArray(error)
                  ? error.map((err, index) => <div key={index}>* {err}</div>)
                  : error}
              </div>
            )}{" "}
          </p>
          <p className="policy-text">
            By logging in, you agree to Axilock's{" "}
            <a href="#">Privacy Policy</a> and <a href="#">Terms of Use</a>.
          </p>
        </div>
        <div className="signup-text" style={{ "marginTop": "24px" }}>
          <Link to="/" style={{ "display": "block" }}>
            <MdArrowBack style={{ verticalAlign: 'text-bottom' }}></MdArrowBack> Login a different way
          </Link>
          {/* <Link to="/sign-up" className="signup-link">
            Sign up for a free trial
          </Link> */}
          {"  "}
          {/* <span> or </span> */}

          <div style={{ "margin": "8px auto" }}>
            <Link to="/">
              <span style={{ "fontSize": "16px", "color": "#000" }}>
                {" "}
                Already have Account?
              </span>
            </Link>
            {" "}
            <a href="#contact" className="">
              Contact Us
            </a>
          </div>
        </div>
      </div>

      <div className="footer">
        <div className="footer-container">
          <div className="footer-right">
            <a href="#" className="footer-link">Privacy Policy</a>
            <span className="footer-divider">|</span>
            <a href="#" className="footer-link">Terms of Use</a>
            <span className="footer-divider">|</span>
            <a href="#" className="footer-link">Contact</a>
            <span className="footer-divider">|</span>
            <a href="#" className="footer-link">Help</a>
          </div>
        </div>
        <div className="footer-bottom">
          <p className="footer-text copyright">© 2024 Axilock. All rights reserved.</p>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
