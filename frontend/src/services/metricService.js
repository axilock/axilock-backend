import apiClient from "../utils/api";

const cache = new Map(); // KEY : (VALUE, TIMESTAMP)
const one_minutes = 1 * 60 * 1000; // 1 minutes


export const getTopReposWithAlerts = () => {
  return Promise.resolve({"data": []})
}

export const coverage = () => {
  return apiClient.get("/commits/health");
};

export const getTopAlertsByType = (count) => {
  return apiClient.get(`/alerts/secret/type?count=${count}`)
}
  
export const tableData = (state, limit, days) => {
  /**
   * state : prevented/blocked
   * limit : number of alerts to be fetched
   * days : number of days to be fetched
   */
  return apiClient.get(`/alerts/repo?state=${state}&limit=${limit}&days=${days}`);
};

export const getAlertsFromState = (state) => {
  if (state !== "open" && state !== "closed") { // this should be a enum
    throw "Invalid State is being provided to getAlertsFromState"
  }
  return  Promise.resolve({"data": []})
  // return apiClient.get("/alerts/all?state=" + state);
};

export const dataOfgraph = (alert_type) => {
  let path = `alerts/${alert_type}/graph`;

  if (cache.has(path)) {
    const [data, timestamp] = cache.get(path);
    if (Date.now() < timestamp) { 
      // serving from cache
      return Promise.resolve(JSON.parse(data));
    } else {
      cache.delete(path); // remove expired cache
    }
  }
  return apiClient.get(path).then((response) => {
    cache.set(path, [JSON.stringify(response), Date.now() + (one_minutes / 2)]); // 30 sec cache
    return response;
  });
};
  