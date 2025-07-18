import React from 'react';
import HomeNav from './navbars/HomeNav';
import Header from './Header';
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';
import OrgNav from './navbars/OrgNav';
import UsersNav from './navbars/UsersNav';
import { getUser } from "../services/userService";

const CommonMain = ({ component: Component, currentTopPage }) => {

  const location = useLocation();
  const navigate = useNavigate();
  
  const [config, setConfig] = useState(); 
  const [darkTheme, setDarkTheme] = useState(() => {
    // Initialize theme from localStorage
    const savedTheme = localStorage.getItem('darkTheme');
    if (savedTheme === null) {
      localStorage.setItem('darkTheme', 'false');
      return false;
    }
    return savedTheme === 'true' ? true : false;
  });

  const toggleTheme = () => {
    setDarkTheme(!darkTheme);
    localStorage.setItem('darkTheme', !darkTheme);
  }

  const globalSetting = async () => {
    try {
      const response = await getUser();
      setConfig(response);
    } catch (e) {
      if (e.status == 401) {
        navigate("/login", { state: { errorMsg: "Please log in" } })
      } else {
        console.error("Header.js:: ", e);
        // navigate("/login", { state: { successMsg: getErrorMessageFromResponse(e) } }) 
        // TODO : if eror in config api, where do we show error ? we dont have a screen
      }
    }
  };


  useEffect(() => {
    globalSetting();
  }, []);
  
  const errorToast = location.state?.errorToast || null
  const successToast = location.state?.successToast || null
  
  // Add state to track sidebar visibility
  const [sidebarCollapsed, setSidebarCollapsed] = useState(false);
  
  // Toggle sidebar function
  const toggleSidebar = () => {
    setSidebarCollapsed(!sidebarCollapsed);
  };
  
  let leftNavComponent = <HomeNav darkTheme={darkTheme} config={config} />;
  if (currentTopPage === "Home") {
    leftNavComponent = <HomeNav darkTheme={darkTheme} config={config} />;  
  } else if (currentTopPage === "OrgSetting") {
    leftNavComponent = <OrgNav darkTheme={darkTheme} config={config} />;
  } else if(currentTopPage === "Users") {
    leftNavComponent = <UsersNav darkTheme={darkTheme} config={config} />;
  }
  
  useEffect(() => {
    if (errorToast) {
      toast.error(errorToast);
    }
    if(successToast) {
      toast.success(successToast);
    }
  }, [errorToast, successToast]);
  
  return (
    <div>
      <ToastContainer
        position="top-center"
        autoClose={4000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme={darkTheme ? "dark" : "light"}
      />
      <div className={`flex flex-col h-screen ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
        <Header 
          toggleTheme={toggleTheme} 
          darkTheme={darkTheme} 
          toggleSidebar={toggleSidebar} 
          currentTopPage={currentTopPage}
          config={config}
        />
        <div className="flex flex-1 overflow-hidden">
          {/* Add transition and transform classes */}
          <div 
            className={`transition-transform duration-300 ease-in-out absolute ${sidebarCollapsed ? '-translate-x-72' : 'translate-x-0'} w-72 flex-shrink-0 z-10`}
          >
            {leftNavComponent}
          </div>
          <main 
            className={`transition-all duration-300 ease-in-out flex-1 px-6 py-6 overflow-y-auto ${sidebarCollapsed ? 'ml-0' : 'ml-72'}`}
          >
            <div className="w-3/4">
              <Component darkTheme={darkTheme} config={config} />
            </div>
          </main>
        </div>
      </div>
    </div>
  )
}

export default CommonMain;
