import apiClient from "../utils/api";

// Fetch users
export const loginUser = (postBody) => {
  return apiClient.post("/user/login", postBody);
};

export const signupUser = (postBody) => {
  return apiClient.post("/user/create", postBody);
};

export const getUser = () => {
  return apiClient.get("/user/details");
  // {email:'', {orgstats: {repocount:'', org_name:'', def_pattern_count:0, custom_pattern_count:0}}
};

export const getOrgUsers = () => {
  return apiClient.get("/users/coverage");
}
