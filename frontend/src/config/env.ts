type Environment = "development" | "production";

interface EnvConfig {
  apiProtocol: string;
  apiHost: string;
  apiPort: string;
}

const getEnvironment = (): Environment => {
  return (process.env.NODE_ENV as Environment) || "development";
};

const loadEnvConfig = (): EnvConfig => {
  const env = getEnvironment();

  if (env === "production") {
    return {
      apiProtocol: process.env.REACT_APP_API_PROTOCOL || "https",
      apiHost: process.env.REACT_APP_API_HOST || "your-production-domain.com",
      apiPort: "", // 本番環境ではポートは使用しない
    };
  }

  // development環境
  return {
    apiProtocol: "http",
    apiHost: "localhost",
    apiPort: process.env.SERVER_PORT || "8080",
  };
};

const envConfig = loadEnvConfig();

export const getApiBaseUrl = (): string => {
  const { apiProtocol, apiHost, apiPort } = envConfig;
  const portSuffix = apiPort ? `:${apiPort}` : "";
  return `${apiProtocol}://${apiHost}${portSuffix}/api`;
};
