import apiClient from "../utils/api";

export const processGithubCallbackAuth = (postBody) => {
  // use this for auth
  return apiClient.get("/auth/github/callback", postBody);
};

export const processGithubAppCallbackUser = (postBody) => {
  // use this for user
  return apiClient.post("/github-app/callback", postBody);
}

export const processGithubCallbackAuthClient = (postBody) => {
  // use this for auth
  return apiClient.post("/auth/cli-auth", postBody);
};

export const processGithubCallbackUser = (postBody) => {
  // use this for user
  return apiClient.post("/user/auth/github", postBody);
}

export const processGithubCallback = (postBody) => {
  // use this while installing github app
  return apiClient.get("/app/github/callback", postBody);
};

export const fetchIntegrations = () => {
  return apiClient.get("/integrations/all");
};