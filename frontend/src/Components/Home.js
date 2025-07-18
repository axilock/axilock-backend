import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { FaArrowUpLong, FaSpinner } from "react-icons/fa6";
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { tableData, dataOfgraph } from "../services/metricService";
import { getErrorMessageFromResponse, formatNumber } from "../utils/helpers";
import Header from "../Common/Header";
import HomeNav from "../Common/navbars/HomeNav";

const Home = ({ darkTheme }) => {
  const navigate = useNavigate();

  const [curr_alert_type, setCurrAlertType] = useState("protected");
  const [data, setData] = useState([]);
  const [graphData, setgraphData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const dateToday = (new Date()).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });

  // Format date for the X-axis
  const formatDate = (dateStr) => {
    if (!dateStr) return "";
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-US', { day: 'numeric', month: 'short' });
  };

  // Format date with year for tooltip
  const formatFullDate = (dateStr) => {
    if (!dateStr) return "";
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-US', { day: 'numeric', month: 'short', year: 'numeric' });
  };

  // Custom tooltip component
  const CustomTooltip = ({ active, payload, label }) => {
    if (active && payload && payload.length) {
      // Calculate total count
      const totalCount = payload.reduce((sum, entry) => sum + (entry.value || 0), 0);

      // Get severity data sorted in order (Critical, High, Medium, Low)
      const severityOrder = ['value_critical', 'value_high', 'value_medium', 'value_low'];
      const sortedPayload = [...payload].sort((a, b) =>
        severityOrder.indexOf(a.dataKey) - severityOrder.indexOf(b.dataKey)
      );

      return (
        <div className={`p-3 rounded shadow-md min-w-[150px] text-xs ${darkTheme ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'} border`}>
          <p className={`font-bold mb-1.5 pb-1 border-b ${darkTheme ? 'border-gray-700 text-gray-200' : 'border-gray-200 text-gray-700'}`}>
            {formatFullDate(label)}
          </p>
          <table className="w-full">
            <tbody>
              {sortedPayload.map((entry, index) => {
                let severityName = '';
                let borderStyle = '';

                switch (entry.dataKey) {
                  case 'value_critical':
                    severityName = 'CRITICAL';
                    borderStyle = 'border-t-2 border-dotted border-pink-500';
                    break;
                  case 'value_high':
                    severityName = 'HIGH';
                    borderStyle = 'border-t-2 border-dashed border-orange-500';
                    break;
                  case 'value_medium':
                    severityName = 'MEDIUM';
                    borderStyle = 'border-t-2 border-solid border-yellow-500';
                    break;
                  case 'value_low':
                    severityName = 'LOW';
                    borderStyle = 'border-t-2 border-dotted border-gray-500';
                    break;
                  default:
                    severityName = entry.dataKey;
                }

                return (
                  <tr key={`row-${index}`} className={index % 2 === 0 ? (darkTheme ? 'bg-gray-700' : 'bg-gray-50') : ''}>
                    <td className="py-1 w-6 align-middle">
                      <div className={`w-5 inline-block ${borderStyle}`}></div>
                    </td>
                    <td className={`py-1 font-medium align-middle ${darkTheme ? 'text-gray-300' : 'text-gray-700'}`}>
                      {severityName}
                    </td>
                    <td className={`py-1 font-bold text-right align-middle ${darkTheme ? 'text-white' : 'text-gray-900'}`}>
                      {entry.value}
                    </td>
                  </tr>
                );
              })}
              <tr className={darkTheme ? 'bg-gray-700' : 'bg-gray-200'}>
                <td colSpan="2" className={`py-1.5 font-medium pt-1.5 border-t ${darkTheme ? 'border-gray-600 text-gray-200' : 'border-gray-300 text-gray-700'}`}>
                  TOTAL
                </td>
                <td className={`py-1.5 font-bold text-right pt-1.5 border-t ${darkTheme ? 'border-gray-600 text-white' : 'border-gray-300 text-gray-900'}`}>
                  {totalCount}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      );
    }
    return null;
  };

  const fetchData = async () => {
    setLoading(true);
    try {
      const [fetchTableData, graphData] = await Promise.all([
        tableData(curr_alert_type, 10, 30),
        dataOfgraph(curr_alert_type),
      ]);

      setData(fetchTableData.data);
      let graph_formatted_data = {};

      if (graphData.data?.buckets?.length > 0) {
        graphData.data.buckets.forEach((val) => {
          if(graph_formatted_data[val.bucket_end] === undefined) {
            graph_formatted_data[val.bucket_end] = {bucket_end: val.bucket_end}
          }
          graph_formatted_data[val.bucket_end][`value_${val.severity}`] = val.cumulative_count
        })
        graphData.data.buckets = Object.values(graph_formatted_data)
        const sortedGraphData = {
          ...graphData.data,
          buckets: graphData.data.buckets.sort((a, b) => {
            return new Date(a.bucket_end) - new Date(b.bucket_end);
          })
        };
        console.log(sortedGraphData)
        setgraphData(sortedGraphData);
      }
    } catch (err) {
      if (err.status === 401) {
        navigate("/login", { state: { errorMsg: "Please log in." } });
      } else {
        console.error("Home.js:: ", err);
        const errorMessage = getErrorMessageFromResponse(err);
        setError(errorMessage);

        toast.error(errorMessage, {
          position: "top-center",
          autoClose: 4000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
        });
      }
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [curr_alert_type]); // Refetch when alert type changes

  useEffect(() => {
    if (error) {
      setData([]);
      setgraphData({});
    }
  }, [error]);


  if (loading) return (
    <div className="min-h-screen">
      <div className={`flex items-center justify-center h-screen ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
        <FaSpinner className={`animate-spin text-4xl ${darkTheme ? 'text-green-400' : 'text-green-500'}`} />
      </div>
    </div>
  );

  return (
    <>
      <div className={`min-h-screen 
        ${darkTheme ? 'bg-gray-800 text-white' : 'bg-white text-gray-800'}
      `}>
        {/* Banner Section */}
        <div className={`p-6 ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
          <div className="mb-4">
            <h1 className={`text-xl font-bold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>Overview</h1>
            <p className={`${darkTheme ? 'text-gray-300' : 'text-gray-600'} text-sm`}>Trends and insights across your organization.</p>
          </div>

          {/* Navigation Tabs */}
          <div className="mb-4">
            <ul className={`flex flex-wrap border-b ${darkTheme ? 'border-gray-700' : 'border-gray-200'}`}>
              <li className="mr-2">
                <button
                  onClick={() => setCurrAlertType("detected")}
                  className={`inline-block py-2 px-4 font-medium text-sm border-b-2 ${curr_alert_type === "detected"
                    ? darkTheme ? 'text-green-400 border-green-400' : 'text-blue-500 border-blue-500'
                    : darkTheme ? 'text-gray-400 border-transparent hover:text-gray-300' : 'text-gray-500 border-transparent hover:text-gray-700'
                    }`}
                >Detected</button>
              </li>
              <li>
                <button
                  onClick={() => setCurrAlertType("protected")}
                  className={`inline-block py-2 px-4 font-medium text-sm border-b-2 ${curr_alert_type === "protected"
                    ? darkTheme ? 'text-green-400 border-green-400' : 'text-blue-500 border-blue-500'
                    : darkTheme ? 'text-gray-400 border-transparent hover:text-gray-300' : 'text-gray-500 border-transparent hover:text-gray-700'
                    }`}
                >Protected</button>
              </li>
            </ul>
          </div>

          {/* Graph Section with Added Shadow/Border */}
          <div style={{ "minHeight": "200px" }} className={`mb-6 p-4 rounded-lg ${darkTheme ? 'bg-gray-800 border border-gray-700 shadow-sm' : 'bg-white border border-gray-200 shadow-sm'}`}>
            {/* Graph Title and Stats */}
            <div className="mb-4">
              <h2 className={`text-lg font-semibold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
                {curr_alert_type === "detected" ? "Open alerts over time" : "Protected Secrets over time"}
              </h2>

              <div className="flex items-center gap-4 mb-2">
                <p className={`text-2xl font-bold ${darkTheme ? 'text-white' : 'text-gray-900'}`}>
                  {graphData?.total}
                </p>
                {graphData?.trend !== undefined && (
                  <p className={`flex items-center text-sm font-medium ${graphData?.trend >= 0
                    ? 'text-green-500'
                    : 'text-red-500'
                    }`}>
                    <FaArrowUpLong className={`mr-1 ${graphData?.trend < 0 ? 'transform rotate-180' : ''
                      }`} />
                    {Math.abs(graphData?.trend.toFixed(0))}%
                  </p>
                )}
                <p className={`text-sm ${darkTheme ? 'text-gray-400' : 'text-gray-500'}`}>
                  as of {dateToday}
                </p>
              </div>
            </div>

            {/* Graph Legend */}
            <div className="flex flex-wrap gap-4 mb-4">
              <div className="flex items-center">
                <div className="w-5 border-t-2 border-dotted border-pink-500 inline-block mr-2"></div>
                <span className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-700'}`}>Critical</span>
              </div>
              <div className="flex items-center">
                <div className="w-5 border-t-2 border-dashed border-orange-500 inline-block mr-2"></div>
                <span className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-700'}`}>High</span>
              </div>
              <div className="flex items-center">
                <div className="w-5 border-t-2 border-solid border-yellow-500 inline-block mr-2"></div>
                <span className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-700'}`}>Medium</span>
              </div>
              <div className="flex items-center">
                <div className="w-5 border-t-2 border-dotted border-gray-500 inline-block mr-2"></div>
                <span className={`text-sm ${darkTheme ? 'text-gray-300' : 'text-gray-700'}`}>Low</span>
              </div>
            </div>

            {/* Graph */}
            {graphData?.buckets?.length > 0 && (
              <div className={`w-full h-80 mt-4 ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
                <ResponsiveContainer width="100%" height="100%">
                  <AreaChart data={graphData?.buckets ?? []}>
                    <CartesianGrid strokeDasharray="3 3" stroke={darkTheme ? '#374151' : '#e5e7eb'} />
                    <XAxis
                      dataKey="bucket_end"
                      tick={{ fontSize: 12 }}
                      tickFormatter={formatDate}
                      stroke={darkTheme ? '#9ca3af' : '#6b7280'}
                    />
                    <YAxis
                      tick={{ fontSize: 12 }}
                      stroke={darkTheme ? '#9ca3af' : '#6b7280'}
                    />
                    <Tooltip content={<CustomTooltip />} />

                    <defs>
                      <linearGradient id="colorValueLow" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="0%" stopColor={darkTheme ? "#6b7280" : "#9ca3af"} stopOpacity={0.25} />
                        <stop offset="100%" stopColor={darkTheme ? "#6b7280" : "#9ca3af"} stopOpacity={0.05} />
                      </linearGradient>
                      <linearGradient id="colorValueMedium" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="0%" stopColor={darkTheme ? "#eab308" : "#facc15"} stopOpacity={0.25} />
                        <stop offset="100%" stopColor={darkTheme ? "#eab308" : "#facc15"} stopOpacity={0.05} />
                      </linearGradient>
                      <linearGradient id="colorValueHigh" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="0%" stopColor={darkTheme ? "#f97316" : "#fb923c"} stopOpacity={0.25} />
                        <stop offset="100%" stopColor={darkTheme ? "#f97316" : "#fb923c"} stopOpacity={0.05} />
                      </linearGradient>
                      <linearGradient id="colorValueCritical" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="0%" stopColor={darkTheme ? "#ec4899" : "#f472b6"} stopOpacity={0.25} />
                        <stop offset="100%" stopColor={darkTheme ? "#ec4899" : "#f472b6"} stopOpacity={0.05} />
                      </linearGradient>
                    </defs>

                    <Area
                      type="monotone"
                      dataKey="value_low"
                      stackId="1"
                      stroke={darkTheme ? "#6b7280" : "#9ca3af"}
                      fill="url(#colorValueLow)"
                      strokeWidth={1.8}
                      name="Low"
                      strokeDasharray="2 2 5 2"
                    />
                    <Area
                      type="monotone"
                      dataKey="value_medium"
                      stackId="1"
                      stroke={darkTheme ? "#eab308" : "#facc15"}
                      fill="url(#colorValueMedium)"
                      strokeWidth={1.8}
                      name="Medium"
                    />
                    <Area
                      type="monotone"
                      dataKey="value_high"
                      stackId="1"
                      stroke={darkTheme ? "#f97316" : "#fb923c"}
                      fill="url(#colorValueHigh)"
                      strokeWidth={1.8}
                      name="High"
                      strokeDasharray="3 4"
                    />
                    <Area
                      type="monotone"
                      dataKey="value_critical"
                      stackId="1"
                      stroke={darkTheme ? "#ec4899" : "#f472b6"}
                      fill="url(#colorValueCritical)"
                      strokeWidth={1.5}
                      name="Critical"
                      strokeDasharray="1 2"
                    />
                  </AreaChart>
                </ResponsiveContainer>
              </div>
            )}
          </div>
        </div>

        {/* Impact Analysis Section */}
        <div className={`p-6 ${darkTheme ? 'bg-gray-800' : 'bg-white'}`}>
          <h2 className={`text-lg font-semibold mb-2 ${darkTheme ? 'text-white' : 'text-gray-800'}`}>
            Impact analysis
          </h2>
          <p className={`mb-4 ${darkTheme ? 'text-gray-300' : 'text-gray-600'}`}>
            Top 10 repositories that pose the biggest impact on your application security.
          </p>

          {/* Table */}
          <div className={`overflow-x-auto rounded-lg ${darkTheme ? 'border border-gray-700 shadow-sm' : 'border border-gray-200 shadow-sm'}`}>
            <table className={`min-w-full ${darkTheme ? 'border-gray-700' : 'border-gray-200'} text-sm`}>
              <thead className={darkTheme ? 'bg-gray-700' : 'bg-gray-50'}>
                <tr>
                  <th className={`px-4 py-2 text-left font-medium ${darkTheme ? 'text-gray-200 border-gray-600' : 'text-gray-700 border-gray-200'} border-b`}>
                    Repositories
                  </th>
                  {curr_alert_type === "detected" ?
                    <th className={`px-4 py-2 text-left font-medium ${darkTheme ? 'text-gray-200 border-gray-600' : 'text-gray-700 border-gray-200'} border-b`}>
                      Open alerts
                    </th>
                    :
                    ""
                  }
                  <th className={`px-4 py-2 text-left font-medium ${darkTheme ? 'text-gray-200 border-gray-600' : 'text-gray-700 border-gray-200'} border-b`}>
                    Critical
                  </th>
                  <th className={`px-4 py-2 text-left font-medium ${darkTheme ? 'text-gray-200 border-gray-600' : 'text-gray-700 border-gray-200'} border-b`}>
                    High
                  </th>
                  <th className={`px-4 py-2 text-left font-medium ${darkTheme ? 'text-gray-200 border-gray-600' : 'text-gray-700 border-gray-200'} border-b`}>
                    Medium
                  </th>
                  <th className={`px-4 py-2 text-left font-medium ${darkTheme ? 'text-gray-200 border-gray-600' : 'text-gray-700 border-gray-200'} border-b`}>
                    Low
                  </th>
                </tr>
              </thead>
              <tbody>
                {data && data.map((item, index) => (
                  <tr key={index} className={index % 2 === 0 ? (darkTheme ? 'bg-gray-700' : 'bg-gray-50') : ''}>
                    <td className={`px-4 py-2 ${darkTheme ? 'border-gray-600' : 'border-gray-200'} border-b`}>
                      {item.repo_name}
                    </td>
                    {curr_alert_type === "detected" ?
                      <td className={`px-4 py-2 ${darkTheme ? 'border-gray-600' : 'border-gray-200'} border-b`}>{item.open_alerts || 0}</td>
                      :
                      ""
                    }
                    <td className={`px-4 py-2 ${darkTheme ? 'border-gray-600' : 'border-gray-200'} border-b`}>
                      {item.critical || item.secret_count || 0}
                    </td>
                    <td className={`px-4 py-2 ${darkTheme ? 'border-gray-600' : 'border-gray-200'} border-b`}>
                      {item.high || 0}
                    </td>
                    <td className={`px-4 py-2 ${darkTheme ? 'border-gray-600' : 'border-gray-200'} border-b`}>
                      {item.medium || 0}
                    </td>
                    <td className={`px-4 py-2 ${darkTheme ? 'border-gray-600' : 'border-gray-200'} border-b`}>
                      {item.low || 0}
                    </td>
                  </tr>
                ))}
                {(!data || data.length === 0) && (
                  <tr>
                    <td
                      colSpan="6"
                      className={`px-4 py-4 text-center ${darkTheme ? 'text-gray-400 border-gray-600' : 'text-gray-500 border-gray-200'} border-b`}
                    >
                      No data available
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </>
  );
}
export default Home;