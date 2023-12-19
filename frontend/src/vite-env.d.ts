/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_AUTHORIZE_URL: string;
  readonly VITE_SKIP_AUTH: boolean;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
