/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_AUTHORIZE_URL: string;
  readonly VITE_SKIP_AUTH: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
