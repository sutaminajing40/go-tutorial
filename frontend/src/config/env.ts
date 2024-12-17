interface EnvConfig {
  apiProtocol: string;
  apiHost: string;
  apiPort: string;
}

const loadEnvConfig = (): EnvConfig => {
  // システムの環境変数から読み込む（.env.localまたはサーバー設定の環境変数）
  return {
    apiProtocol: process.env.REACT_APP_API_PROTOCOL || "http",
    apiHost: process.env.REACT_APP_API_HOST || "localhost",
    apiPort: process.env.REACT_APP_API_PORT || "8080",
  };
};

const envConfig = loadEnvConfig();

export const getApiBaseUrl = (): string => {
  const { apiProtocol, apiHost, apiPort } = envConfig;
  const portSuffix = apiPort ? `:${apiPort}` : "";
  return `${apiProtocol}://${apiHost}${portSuffix}/api`;
};
