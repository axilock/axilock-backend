import React, { useState, useEffect } from "react";
import "./Auth.css";
import { loginUser } from "../../services/userService";
import { useNavigate, Link, useLocation } from "react-router-dom";
import { ReactComponent as LOGO } from '../../Images/logo.svg';
import { MdArrowBack } from "react-icons/md";
import { AiFillEye, AiFillEyeInvisible } from "react-icons/ai";
import { FaSpinner } from "react-icons/fa";
import { getUser } from "../../services/userService";
import { Nav } from "react-bootstrap";
import { validateEmail, validatePassword, getErrorMessageFromResponse } from "../../utils/helpers";

const Login = () => {
  // const location = useLocation();
  // const [email, setEmail] = useState("");                   // Email state
  // const [password, setPassword] = useState("");             // Password state
  // const [showPassword, setShowPassword] = useState(false); // Show/Hide password toggle
  // const [successMsg, setSuccessMsg] = useState(location.state?.successMsg || "");
  // const [loading, setLoading] = useState(false);
  // const [error, setError] = useState("");
  // const navigate = useNavigate();


  // useEffect(() => {
  //   if (successMsg) {
  //     const timer = setTimeout(() => {
  //       setSuccessMsg("");
  //     }, 5000); // 5 seconds

  //     return () => clearTimeout(timer); // Clean up the timer on component unmount
  //   }
  // }, [successMsg]);

  // const globalSetting = async () => {
  //   try {
  //     const response = await getUser();
  //     navigate("/insights");
  //   } catch (e) {
  //     console.error("Login.js:: ", e);
  //   }
  // };

  // useEffect(() => {
  //   globalSetting();
  // }, []);

  
  // const handleLogin = async () => {
  //   setLoading(true);
  //   setError(""); // Reset error message
  //   setSuccessMsg("");

  //   if (!validateEmail(email)) {
  //     setError("Invalid email format.");
  //     setLoading(false)
  //     return;
  //   }

  //   try {

  //     const response = await loginUser({ email, password }); // Call the API

  //     if (response.accesstoken) {
  //       // Store token in localStorage
  //       localStorage.setItem("auth-token", response["accesstoken"]);
  //       navigate("/insights"); //TODO: Need to redirect to the third page of the home screen
  //     } else {
  //       console.error("Login.js:: ", response);
  //       setError(getErrorMessageFromResponse(response));
  //     }
  //   } catch (error) {
  //     setError(`Login Failed : ${getErrorMessageFromResponse(error)}`);
  //     console.error("Login.js", error);
  //   } finally {
  //     setLoading(false); // Stop loading spinner
  //   }
  // };
  // return (
  //   <div>
  //     <div className="logo-container">
  //       <Nav.Link href="/" className="logo-link">
  //         <LOGO style={{ width: '48px' }}></LOGO>
  //       </Nav.Link>
  //     </div>

  //     <h2 className="subtitle">Sign in to Axilock</h2>
  //     <div className="form-container" style={{ "paddingTop": "35px" }}>
  //       <div className="login-box-content">
  //         <label htmlFor="email">EMAIL*</label>
  //         <input
  //           type="email"
  //           id="email"
  //           placeholder="Enter your email address"
  //           className="input-box"
  //           onChange={(e) => setEmail(e.target.value)}
  //           onKeyDown={(e) => {
  //             if (e.key === 'Enter') {
  //               handleLogin();
  //             }
  //           }}
  //         />
  //         <label htmlFor="password">PASSWORD*</label>
  //         <div className="password-container">
  //           <input
  //             type={showPassword ? "text" : "password"} // Toggle input type
  //             id="password"
  //             value={password}
  //             onChange={(e) => setPassword(e.target.value)} // Update password state
  //             placeholder="Enter Password"
  //             className="input-box"
  //             onKeyDown={(e) => {
  //               if (e.key === 'Enter') {
  //                 handleLogin();
  //               }
  //             }}
  //           />
  //           <button
  //             type="button"
  //             className="toggle-password-btn"
  //             onClick={() => setShowPassword(!showPassword)} // Toggle state
  //           >
  //             {showPassword ? <AiFillEyeInvisible /> : <AiFillEye />}
  //           </button>
  //         </div>
  //         <button
  //           className="continue-btn"
  //           onClick={handleLogin}
  //           disabled={loading}>
  //           {loading ? (
  //             <FaSpinner className="rotating-icon" />
  //           ) : (
  //             "Continue"
  //           )}
  //         </button>
  //         <p>
  //           {" "}
  //           {error && (<div className="error-message">{JSON.stringify(error).slice(1, -1)}</div>)}
  //           {successMsg && <div className="success-message">{successMsg}</div>}
  //           {" "}
  //         </p>
  //         <p className="policy-text">
  //             By logging in, you agree to Axilock's{" "}
  //             <a href="#">Privacy Policy</a> and <a href="#">Terms of Use</a>.
  //           </p>
  //       </div>
  //       <div className="signup-text">
  //         <Link to="/" className="">
  //           <MdArrowBack style={{ verticalAlign: 'text-bottom' }}></MdArrowBack> Login a different way
  //         </Link>
  //         <p style={{ "margin": "10px auto" }}>Don't have an account?</p>
  //         <Link to="/sign-up" className="signup-link">
  //           Sign up for a free trial
  //         </Link>
  //         {" "}
  //         <span> or </span>
  //         <a href="#contact" className="">
  //           Contact Us
  //         </a>
  //       </div>
  //     </div>
  //     <div className="footer">
  //       <div className="footer-container">
  //         <div className="footer-right">
  //           <a href="#" className="footer-link">Privacy Policy</a>
  //           <span className="footer-divider">|</span>
  //           <a href="#" className="footer-link">Terms of Use</a>
  //           <span className="footer-divider">|</span>
  //           <a href="#" className="footer-link">Contact</a>
  //           <span className="footer-divider">|</span>
  //           <a href="#" className="footer-link">Help</a>
  //         </div>
  //       </div>
  //       <div className="footer-bottom">
  //         <p className="footer-text copyright">© 2024 Axilock. All rights reserved.</p>
  //       </div>
  //     </div>
  //   </div>
  // );
};

export default Login;
