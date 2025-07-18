// Mock data for when GitHub is not onboarded

// Coverage data
export const mockCoverageData = {
  total_commits: 0,
  not_scanned_count: 0,
  scanned_count: 0,
  percent_protected: 100,
  percent_not_protected: 0
};

// Custom rules data
export const mockAllAlertsList = [
  { alert_name: "AWS Secret Key", source: "GitHub", source_repo: "frontend-app", date: "2025-06-01", path: "/src/config.js" },
  { alert_name: "Database Password", source: "GitLab", source_repo: "backend-api", date: "2025-06-02", path: "/db/config.py" },
  { alert_name: "API Token", source: "GitHub", source_repo: "mobile-client", date: "2025-06-05", path: "/app/services/api.js" },
];

export const mockTopAlertsByType = [
  { type: "AWS Keys", count: 12 },
  { type: "Database Credentials", count: 8 },
  { type: "API Tokens", count: 5 },
];

export const mockTopReposWithSecrets = [
  { repo: "frontend-app", count: 23 },
  { repo: "backend-api", count: 18 },
  { repo: "data-pipeline", count: 12 },
  { repo: "mobile-client", count: 9 }
];
