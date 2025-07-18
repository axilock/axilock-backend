import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import AuthScreen from "./Components/Login/AuthScreen.js";
import AuthScreenClient from "./Components/Login/AuthScreenClient.js";
import Home from "./Components/Home.js";
import Coverage from "./Components/Coverage.js";
import Repositories from "./Components/Repositories.js";
import Customrules from "./Components/Customrules.js";
import Secretscan from "./Components/Secretscanningrules.js";

import EmptyPage from "./Components/EmptyPage.js";
import Users from "./Components/Users.js";
import IntegrationsScreen from "./Components/Integrations.js"
import GithubCallback from "./Components/Callbacks/GithubCallback"
import GithubAppCallback from "./Components/Callbacks/GithubAppCallback"
import GithubCallbackClient from "./Components/Callbacks/GithubCallbackClient"
import CommonMain from "./Common/CommonMain.js";


function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AuthScreen currentTopPage="Home"/>} />
        <Route path="/login" element={<AuthScreen currentTopPage="AuthScreen" />} />
        <Route path="/client/login" element={<AuthScreenClient currentTopPage="AuthScreen" />} />
        <Route path="/insights" element={<CommonMain component={Home} currentTopPage="Home" />} />
        <Route path="/insights/coverage" element={<CommonMain component={Coverage} currentTopPage="Home" />} />
        <Route path="/insights/alerts/secret-scanning" element={<CommonMain component={Customrules} currentTopPage="Home" />} />
        <Route path="/users" element={<CommonMain component={Users} currentTopPage="Users" />} />
        <Route path="/repositories" element={<CommonMain component={Repositories} currentTopPage="Repositories" />} />
        <Route path="/org-setting" element={<CommonMain component={EmptyPage} currentTopPage="OrgSetting" />} />
        <Route path="/org-setting/custom-secrets" element={<CommonMain component={Secretscan} currentTopPage="OrgSetting" />} />
        <Route path="/org-setting/integrations" element={<CommonMain component={IntegrationsScreen} currentTopPage="OrgSetting" />} />
        <Route path="/auth/github/callback" element={<GithubCallback />} />
        <Route path="/auth/github-app/callback" element={<GithubAppCallback />} />
        <Route path="/client/github/callback" element={<GithubCallbackClient />} />
        {/* <Route path="/test" element={<CommonMain component={EmptyPage} />} /> */}
        {/* <Route path="/login" element={<Login />} /> */}
        {/* <Route path="/sign-up" element={<SignUp />} /> */}
        {/* <Route path="/secret-scan-edit" element={<Secretscanedit />} />        */}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
